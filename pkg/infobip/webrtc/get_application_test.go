package webrtc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetWebRTCApplication(t *testing.T) {
	rawJSONResp := []byte(`
		{
		  "id": "894c822b-d7ba-439c-a761-141f591cace7",
		  "name": "Application Name",
		  "description": "Application Description",
		  "ios": {
			"apnsCertificateFileName": "IOS_APNS_certificate.p",
			"apnsCertificatePassword": "IOS_APNS_certificate_password"
		  },
		  "android": {
			"fcmServerKey": "H)AWbDphXcRyBJd7UmcTZmziH6aMVqsGdDp1NRPyUHKwKRey-AESZ1DPdpbux24r3GaUWl_mLfjqzz1xuo"
		  },
		  "appToApp": true,
		  "appToConversations": true,
		  "appToPhone": true
		}
	`)

	apiKey := "some-api-key"
	var expectedResp models.GetWebRTCApplicationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.Contains(r.URL.Path, getApplicationPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	webRTC := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := webRTC.GetApplication(context.Background(), "some-application-id")

	require.NoError(t, err)
	assert.NotEqual(t, models.GetWebRTCApplicationResponse{}, resp)
	assert.NotEmpty(t, resp.ID)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
