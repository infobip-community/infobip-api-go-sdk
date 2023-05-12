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
	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTemplateValidReq(t *testing.T) {
	sender := "16175551213"
	apiKey := "secret"
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "MARKETING",
		Structure: models.TemplateStructure{
			Body: &models.TemplateStructureBody{
				Text: "body {{1}} content",
			},
			Type: "TEXT",
		},
	}
	rawJSONResp := []byte(`
		{
		  "id": "111",
		  "businessAccountId": 222,
		  "name": "media_template",
		  "language": "en",
		  "status": "APPROVED",
		  "category": "MARKETING",
		  "structure": {
			"header": {
			  "format": "IMAGE"
			},
			"body": {
			  "text": "example {{1}} body"
			},
			"footer": {
			  "text": "exampleFooter"
			},
			"type": "MEDIA"
		  }
		}
	`)
	var expectedResp models.CreateWATemplateResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(templatesPath, sender)))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.TemplateCreate
		servErr = json.Unmarshal(parsedBody, &receivedMsg)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedMsg, template)

		w.WriteHeader(http.StatusCreated)
		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	whatsApp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	messageResponse, respDetails, err := whatsApp.CreateTemplate(context.Background(), sender, template)

	require.NoError(t, err)
	assert.NotEqual(t, models.CreateWATemplateResponse{}, messageResponse)
	assert.Equal(t, expectedResp, messageResponse)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidCreateTemplate(t *testing.T) {
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "invalid",
		Structure: models.TemplateStructure{
			Body: &models.TemplateStructureBody{
				Text: "body {{1}} content",
			},
			Type: "TEXT",
		},
	}
	whatsApp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    "https://something.api.infobip.com",
		APIKey:     "secret",
	}}

	messageResponse, respDetails, err := whatsApp.CreateTemplate(context.Background(), "16175551213", template)

	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.CreateWATemplateResponse{}, messageResponse)
	assert.Equal(t, models.ResponseDetails{}, respDetails)
}

func TestCreateTemplate4xxErrors(t *testing.T) {
	sender := "16175551213"
	tests := []struct {
		rawJSONResp []byte
		statusCode  int
	}{
		{
			rawJSONResp: []byte(`{
				"requestError": {
					"serviceException": {
						"messageId": "BAD_REQUEST",
						"text": "Bad request"
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
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "MARKETING",
		Structure: models.TemplateStructure{
			Body: &models.TemplateStructureBody{
				Text: "body {{1}} content",
			},
			Type: "TEXT",
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

			msgResp, respDetails, err := whatsApp.CreateTemplate(context.Background(), sender, template)
			serv.Close()

			require.NoError(t, err)
			assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
			assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
			assert.Equal(t, expectedResp, respDetails.ErrorResponse)
			assert.Equal(t, tc.statusCode, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.CreateWATemplateResponse{}, msgResp)
		})
	}
}
