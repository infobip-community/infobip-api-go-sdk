package webrtc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateApplicationValidReq(t *testing.T) {
	apiKey := "some-key"
	app := models.GenerateWebRTCApplication()
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
			"fcmServerKey": "AAAAtm7JlCY:-AESZ1DPdpbux24r3GaUWl_mLfjqzz1xuo"
		  },
		  "appToApp": true,
		  "appToConversations": true,
		  "appToPhone": true
		}
	`)

	var expectedResp models.UpdateWebRTCApplicationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, saveApplicationPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.WebRTCApplication
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, app)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	webRTC := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := webRTC.UpdateApplication(context.Background(), "some-app-id", app)

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateWebRTCApplicationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
