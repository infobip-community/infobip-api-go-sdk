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

func TestGetWebRTCApplications(t *testing.T) {
	rawJSONResp := []byte(`
		[
		  {
			"id": "894c822b-d7ba-439c-a761-141f591cace7",
			"name": "Application Name 1",
			"description": "Application Description",
			"ios": {
			  "apnsCertificateFileName": "IOS_APNS_certificate.p",
			  "apnsCertificatePassword": "IOS_APNS_certificate_password"
			},
			"appToApp": true,
			"appToConversations": false,
			"appToPhone": true
		  },
		  {
			"id": "988c411a-a2db-227a-c563-245a159dcbe2",
			"name": "Application Name 2",
			"description": "Application Description",
			"android": {
			  "fcmServerKey": "AAAAtm7JlCY:-AESZ1DPdpbux24r3GaUWl_mLfjqzz1xuo"
			},
			"appToApp": true,
			"appToConversations": false,
			"appToPhone": true
		  },
		  {
			"id": "454d142b-a1ad-239a-d231-227fa335aadc3",
			"name": "Application Name 3",
			"description": "Application Description",
			"appToApp": true,
			"appToConversations": true,
			"appToPhone": false
		  }
		]
	`)

	apiKey := "some-api-key"
	var expectedResp models.GetWebRTCApplicationsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, getApplicationsPath))
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

	resp, respDetails, err := webRTC.GetApplications(context.Background())

	require.NoError(t, err)
	assert.NotEqual(t, models.GetWebRTCApplicationsResponse{}, resp)
	assert.NotEmpty(t, resp[0].ID)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
