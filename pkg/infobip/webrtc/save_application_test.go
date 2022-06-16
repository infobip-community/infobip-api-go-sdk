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

	"github.com/infobip-community/infobip-api-go-sdk/v2/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveApplicationValidReq(t *testing.T) {
	apiKey := "some-key"
	app := models.GenerateWebRTCApplication()
	rawJSONResp := []byte(`
		{
		  "name": "Application Name",
		  "description": "Application Description",
		  "android": {
			"fcmServerKey": "AAAAtm7JlCYbDphXcRyBJd7UmcTZmziH6aMVqsGdDp1NRPyUHKwKRey-AESZ1DPdpbux24r3GaUWl_mLfjqzz1xuo"
		  },
		  "appToApp": true,
		  "appToConversations": false,
		  "appToPhone": false
		}
	`)

	var expectedResp models.SaveWebRTCApplicationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, saveApplicationPath))
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

	msgResp, respDetails, err := webRTC.SaveApplication(context.Background(), app)

	require.NoError(t, err)
	assert.NotEqual(t, models.SaveWebRTCApplicationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
