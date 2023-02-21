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

func TestGetTFAApplicationValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	  {
		"applicationId": "0933F3BC087D2A617AC6DCB2EF5B8A61",
		"name": "Test application BASIC 1",
		"configuration": {
		  "pinAttempts": 10,
		  "allowMultiplePinVerifications": true,
		  "pinTimeToLive": "2h",
		  "verifyPinLimit": "1/3s",
		  "sendPinPerApplicationLimit": "10000/1d",
		  "sendPinPerPhoneNumberLimit": "3/1d"
		},
		"enabled": true
	  }
	`)

	var expectedResp models.GetTFAApplicationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, getTFAApplicationPath))
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

	msgResp, respDetails, err := sms.GetTFAApplication(context.Background(), "0933F3BC087D2A617AC6DCB2EF5B8A61")

	require.NoError(t, err)
	assert.NotEqual(t, models.GetTFAApplicationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
