package infobip

import (
	"context"
	"encoding/json"
	"fmt"
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

func TestCreateTemplateValidReq(t *testing.T) {
	sender := "16175551213"
	apiKey := "secret"
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "ACCOUNT_UPDATE",
		Structure: models.TemplateStructure{
			Body: "body {{1}} content",
			Type: "TEXT",
		},
	}
	rawJSONResp := []byte(`{
		"id": "111",
		"businessAccountId": 222,
		"name": "exampleName",
		"language": "en",
		"status": "APPROVED",
		"category": "ACCOUNT_UPDATE",
		"structure": {
			"header": {"format": "IMAGE"},
			"body": "example {{1}} body",
			"footer": "exampleFooter",
			"type": "MEDIA"
		}
	}`)
	var expectedResp models.TemplateResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.Nil(t, err)

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

	host := serv.URL
	whatsApp := whatsAppChannel{reqHandler: httpHandler{
		httpClient: http.Client{},
		baseURL:    host,
		apiKey:     apiKey,
	}}
	messageResponse, respDetails, err := whatsApp.CreateTemplate(context.Background(), sender, template)

	require.Nil(t, err)
	assert.NotEqual(t, models.TemplateResponse{}, messageResponse)
	assert.Equal(t, expectedResp, messageResponse)
	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidCreateTemplate(t *testing.T) {
	sender := "16175551213"
	apiKey := "secret"
	whatsApp := whatsAppChannel{reqHandler: httpHandler{
		httpClient: http.Client{},
		baseURL:    "https://something.api.infobip.com",
		apiKey:     apiKey,
	}}
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "invalid",
		Structure: models.TemplateStructure{
			Body: "body {{1}} content",
			Type: "TEXT",
		},
	}

	messageResponse, respDetails, err := whatsApp.CreateTemplate(context.Background(), sender, template)
	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.TemplateResponse{}, messageResponse)
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
	apiKey := "secret"
	template := models.TemplateCreate{
		Name:     "template_name",
		Language: "en",
		Category: "ACCOUNT_UPDATE",
		Structure: models.TemplateStructure{
			Body: "body {{1}} content",
			Type: "TEXT",
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

			host := serv.URL
			whatsApp := whatsAppChannel{reqHandler: httpHandler{
				httpClient: http.Client{},
				baseURL:    host,
				apiKey:     apiKey,
			}}
			messageResponse, respDetails, err := whatsApp.CreateTemplate(context.Background(), sender, template)
			serv.Close()

			require.Nil(t, err)
			assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
			assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
			assert.Equal(t, expectedResp, respDetails.ErrorResponse)
			assert.Equal(t, tc.statusCode, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.TemplateResponse{}, messageResponse)
		})
	}
}
