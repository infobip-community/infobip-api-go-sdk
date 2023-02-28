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

func TestUpdateTFAApplicationValidReq(t *testing.T) {
	apiKey := "some-key"
	request := models.GenerateUpdateTFAApplicationRequest()
	rawJSONResp := []byte(`
		{
		  "applicationId": "1234567",
		  "name": "Application name",
		  "configuration": {
			"pinAttempts": 5,
			"allowMultiplePinVerifications": true,
			"pinTimeToLive": "10m",
			"verifyPinLimit": "2/4s",
			"sendPinPerApplicationLimit": "5000/12h",
			"sendPinPerPhoneNumberLimit": "2/1d"
		  },
		  "enabled": true
		}
	`)

	var expectedResp models.UpdateTFAApplicationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, createTFAApplicationPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := io.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.UpdateTFAApplicationRequest
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

	msgResp, respDetails, err := sms.UpdateTFAApplication(context.Background(), "1234567", request)

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateTFAApplicationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
