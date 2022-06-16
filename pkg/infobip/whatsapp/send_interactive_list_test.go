package whatsapp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/infobip-community/infobip-api-go-sdk/v2/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInteractiveListValidReq(t *testing.T) {
	apiKey := "secret"
	msg := models.WAInteractiveListMsg{
		MsgCommon: models.GenerateTestMsgCommon(),
		Content: models.InteractiveListContent{
			Body: models.InteractiveListBody{Text: "Some text"},
			Action: models.InteractiveListAction{
				Title:    "Choose one",
				Sections: []models.InteractiveListSection{{Rows: []models.SectionRow{{ID: "1", Title: "row title"}}}},
			},
		},
	}
	rawJSONResp := []byte(`{
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
	}`)
	var expectedResp models.SendWAMsgResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendInteractiveListPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.WAInteractiveListMsg
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

	msgResponse, respDetails, err := whatsApp.SendInteractiveList(context.Background(), msg)

	require.NoError(t, err)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResponse)
	assert.Equal(t, expectedResp, msgResponse)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidInteractiveListMsg(t *testing.T) {
	msg := models.WAInteractiveListMsg{
		MsgCommon: models.GenerateTestMsgCommon(),
		Content: models.InteractiveListContent{
			Body: models.InteractiveListBody{Text: "Some text"},
			Action: models.InteractiveListAction{
				Title:    "Choose one",
				Sections: []models.InteractiveListSection{{Rows: []models.SectionRow{{ID: "1"}}}},
			},
		},
	}
	whatsApp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    "https://something.api.infobip.com",
		APIKey:     "secret",
	}}

	msgResponse, respDetails, err := whatsApp.SendInteractiveList(context.Background(), msg)

	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.SendWAMsgResponse{}, msgResponse)
	assert.Equal(t, models.ResponseDetails{}, respDetails)
}

func TestInteractiveList4xxErrors(t *testing.T) {
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
							"content.header.text": [
								"size must be between 1 and 60",
								"must not be blank"
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
	msg := models.WAInteractiveListMsg{
		MsgCommon: models.GenerateTestMsgCommon(),
		Content: models.InteractiveListContent{
			Body: models.InteractiveListBody{Text: "Some text"},
			Action: models.InteractiveListAction{
				Title:    "Choose one",
				Sections: []models.InteractiveListSection{{Rows: []models.SectionRow{{ID: "1", Title: "row title"}}}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.statusCode), func(t *testing.T) {
			var expectedResp models.ErrorDetails
			err := json.Unmarshal(tc.rawJSONResp, &expectedResp)
			require.NoError(t, err)
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

			messageResponse, respDetails, err := whatsApp.SendInteractiveList(context.Background(), msg)
			serv.Close()

			require.NoError(t, err)
			assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
			assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
			assert.Equal(t, expectedResp, respDetails.ErrorResponse)
			assert.Equal(t, tc.statusCode, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.SendWAMsgResponse{}, messageResponse)
		})
	}
}
