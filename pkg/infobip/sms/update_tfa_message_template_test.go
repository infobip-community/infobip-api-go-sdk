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

func TestUpdateTFAMessageTemplateValidReq(t *testing.T) {
	apiKey := "some-key"
	request := models.GenerateUpdateTFAMessageTemplateRequest()
	rawJSONResp := []byte(`
		{
		  "pinPlaceholder": "{{pin}}",
		  "messageText": "Your pin is {{pin}}",
		  "pinLength": 4,
		  "pinType": "ALPHANUMERIC",
		  "language": "en",
		  "senderId": "Infobip 2FA",
		  "repeatDTMF": "1#",
		  "speechRate": 1
		}	
	`)

	var expectedResp models.UpdateTFAMessageTemplateResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, createTFAMessageTemplatePath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := io.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.UpdateTFAMessageTemplateRequest
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

	msgResp, respDetails, err := sms.UpdateTFAMessageTemplate(context.Background(), "1234567",
		"1234567", request)

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateTFAApplicationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
