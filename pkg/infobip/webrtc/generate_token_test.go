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

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateTokenValidReq(t *testing.T) {
	apiKey := "some-key"
	req := models.GenerateWebRTCTokenRequest{
		Identity: "some-identity",
	}
	rawJSONResp := []byte(`
		{
		  "token": "dvbmRlcmxhbmQiLCJleHAiOjE1NzkyOTA2MzgsImNhcHMiOlsyXX0.QyCMqjH8DsftChibW2Rw4EByH-eEviUp3-kHVKuJpKg",
		  "expirationTime": "2020-01-17T19:50:38.488589Z"
		}
	`)

	var expectedResp models.GenerateWebRTCTokenResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, generateTokenPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.GenerateWebRTCTokenRequest
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, req)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	webRTC := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := webRTC.GenerateToken(context.Background(), req)

	require.NoError(t, err)
	assert.NotEqual(t, models.GenerateWebRTCTokenResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
