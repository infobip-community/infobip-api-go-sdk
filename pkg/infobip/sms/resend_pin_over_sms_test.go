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

func TestResendPINOverSMSValidReq(t *testing.T) {
	apiKey := "some-key"
	request := models.GenerateResendPINOverSMSRequest()
	rawJSONResp := []byte(`
		{
		  "pinId": "9C817C6F8AF3D48F9FE553282AFA2B67",
		  "to": "41793026727",
		  "ncStatus": "NC_DESTINATION_REACHABLE",
		  "smsStatus": "MESSAGE_SENT"
		}
	`)

	var expectedResp models.ResendPINOverSMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, resendPINOverSMSPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := io.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.ResendPINOverSMSRequest
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

	msgResp, respDetails, err := sms.ResendPINOverSMS(context.Background(),
		"9C817C6F8AF3D48F9FE553282AFA2B67", request)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResendPINOverSMSResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
