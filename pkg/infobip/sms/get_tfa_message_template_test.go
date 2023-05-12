package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTFAMessageTemplateValidReq(t *testing.T) {
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

	var expectedResp models.GetTFAMessageTemplateResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, getTFAMessageTemplatePath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := sms.GetTFAMessageTemplate(context.Background(),
		"0933F3BC087D2A617AC6DCB2EF5B8A61", "16A8B5FE2BCD6CA716A2D780CB3F3390")

	require.NoError(t, err)
	assert.NotEqual(t, models.GetTFAMessageTemplateResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
