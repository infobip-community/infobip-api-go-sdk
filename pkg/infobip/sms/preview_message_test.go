package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPreviewSMSValidReq(t *testing.T) {
	request := models.GeneratePreviewSMSRequest()
	rawJSONResp := []byte(`
		{
			"originalText": "Let's see how many characters will remain unused in this message .",
			"previews": [
				{
					 "textPreview": "Let's see how many characters will remain unused in this message .",
					 "messageCount": 1,
					 "charactersRemaining": 94,
					 "configuration": {}
				},
				{
					 "textPreview": "Let's see how many characters will remain unused in this message .",
					 "messageCount": 1,
					 "charactersRemaining": 89,
					 "configuration": {
						  "language": {
							  "languageCode": "TR"
						  }
					 }
				},
				{
					 "textPreview": "Let's see how many characters will remain unused in this message .",
					 "messageCount": 1,
					 "charactersRemaining": 86,
					 "configuration": {
						  "transliteration": "TURKISH"
					 }
				}
			]
		}
	`)

	var expectedResp models.PreviewSMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, previewSMSPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := io.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.PreviewSMSRequest
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, request)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := sms.Preview(context.Background(), request)

	require.NoError(t, err)
	assert.NotEqual(t, models.PreviewSMSResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
