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

func TestGetTFAApplicationsValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		[
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
		  },
		  {
			"applicationId": "5F04FACFAA4978F62FCAEBA97B37E90F",
			"name": "Test application BASIC 2",
			"configuration": {
			  "pinAttempts": 12,
			  "allowMultiplePinVerifications": true,
			  "pinTimeToLive": "10m",
			  "verifyPinLimit": "2/1s",
			  "sendPinPerApplicationLimit": "10000/1d",
			  "sendPinPerPhoneNumberLimit": "5/1h"
			},
			"enabled": true
		  },
		  {
			"applicationId": "B450F966A8EF017180F148AF22C42642",
			"name": "Test application BASIC 3",
			"configuration": {
			  "pinAttempts": 15,
			  "allowMultiplePinVerifications": true,
			  "pinTimeToLive": "1h",
			  "verifyPinLimit": "30/10s",
			  "sendPinPerApplicationLimit": "10000/3d",
			  "sendPinPerPhoneNumberLimit": "10/20m"
			},
			"enabled": true
		  }
		]
	`)

	var expectedResp models.GetTFAApplicationsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, getTFAApplicationsPath))
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

	msgResp, respDetails, err := sms.GetTFAApplications(context.Background())

	require.NoError(t, err)
	assert.NotEqual(t, models.GetTFAApplicationsResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
