package whatsapp

import (
	"context"
	"encoding/json"
	"fmt"
	"infobip-go-client/internal"
	"infobip-go-client/pkg/infobip/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateMsgValidReq(t *testing.T) {
	apiKey := "secret"
	msg := models.TemplateMsgs{
		Messages: []models.TemplateMsg{
			{
				MsgCommon: models.MsgCommon{From: "16175551213", To: "16175551212"},
				Content: models.TemplateMsgContent{
					TemplateName: "template_name",
					TemplateData: models.TemplateData{
						Body: models.TemplateBody{Placeholders: []string{}},
					},
					Language: "en_GB",
				},
			},
		},
	}
	rawJSONResp := []byte(`{
		"messages": [
			{
				"to": "441134960001",
				"messageCount": 1,
				"messageId": "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				"status": {
					"groupId": 1,
					"groupName": "PENDING",
					"id": 7,
					"name": "PENDING_ENROUTE",
					"description": "Message sent to next instance"
				}
			}
		],
		"bulkId": "2034072219640523073"
	}`)
	var expectedResp models.BulkMsgResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.Nil(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendTemplateMessagesPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.TemplateMsgs
		servErr = json.Unmarshal(parsedBody, &receivedMsg)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedMsg, msg)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	whatsApp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := whatsApp.SendTemplateMsgs(context.Background(), msg)

	require.Nil(t, err)
	assert.NotEqual(t, models.BulkMsgResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidTemplateMsg(t *testing.T) {
	msg := models.TemplateMsgs{
		Messages: []models.TemplateMsg{
			{
				MsgCommon: models.MsgCommon{From: "16175551213", To: "16175551212"},
				Content: models.TemplateMsgContent{
					TemplateName: "INVALID",
					TemplateData: models.TemplateData{
						Body: models.TemplateBody{Placeholders: []string{}},
					},
					Language: "en_GB",
				},
			},
		},
	}
	whatsApp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    "https://something.api.infobip.com",
		APIKey:     "secret",
	}}

	msgResp, respDetails, err := whatsApp.SendTemplateMsgs(context.Background(), msg)

	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.BulkMsgResponse{}, msgResp)
	assert.Equal(t, models.ResponseDetails{}, respDetails)
}

func TestTemplateMsg4xxErrors(t *testing.T) {
	tests := []struct {
		rawJSONResp []byte
		statusCode  int
	}{
		{
			rawJSONResp: []byte(`{
				"requestError": {
					"serviceException": {
						"messageId": "BAD_REQUEST",
						"text": "Bad request",
						"validationErrors": {
							"content.templateData.header.placeholder": [
								"cannot have new-line/tab characters or more than 4 consecutive spaces"	
							]
						}
					}
				}
			}`),
			statusCode: http.StatusBadRequest,
		},
		{
			rawJSONResp: []byte(`{
				"requestError": {
					"serviceException": {
						"messageId": "UNAUTHORIZED",
						"text": "Invalid login details"
					}
				}
			}`),
			statusCode: http.StatusUnauthorized,
		},
		{
			rawJSONResp: []byte(`{
				"requestError": {
					"serviceException": {
						"messageId": "TOO_MANY_REQUESTS",
						"text": "Too many requests"
					}
				}
			}`),
			statusCode: http.StatusTooManyRequests,
		},
	}
	msg := models.TemplateMsgs{
		Messages: []models.TemplateMsg{
			{
				MsgCommon: models.MsgCommon{From: "16175551213", To: "16175551212"},
				Content: models.TemplateMsgContent{
					TemplateName: "template_name",
					TemplateData: models.TemplateData{
						Body: models.TemplateBody{Placeholders: []string{}},
					},
					Language: "en_GB",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.statusCode), func(t *testing.T) {
			var expectedResp models.ErrorDetails
			err := json.Unmarshal(tc.rawJSONResp, &expectedResp)
			require.Nil(t, err)
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				_, servErr := w.Write(tc.rawJSONResp)
				assert.Nil(t, servErr)
			}))
			whatsApp := Channel{ReqHandler: internal.HTTPHandler{
				HTTPClient: http.Client{},
				BaseURL:    serv.URL,
				APIKey:     "secret",
			}}

			msgResp, respDetails, err := whatsApp.SendTemplateMsgs(context.Background(), msg)
			serv.Close()

			require.Nil(t, err)
			assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
			assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
			assert.Equal(t, expectedResp, respDetails.ErrorResponse)
			assert.Equal(t, tc.statusCode, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.BulkMsgResponse{}, msgResp)
		})
	}
}
