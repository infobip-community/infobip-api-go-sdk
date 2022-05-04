package sms

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

func TestSendMessageValidReq(t *testing.T) {
	apiKey := "some-key"
	request := models.GenerateSendSMSRequest()
	rawJSONResp := []byte(`
		{
			"bulkId": "33644485987105283959",
			"messages": [
				{
					 "to": "41793026727",
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 26,
						  "name": "PENDING_ACCEPTED",
						  "description": "Message sent to next instance"
					 },
					 "messageId": "33644485987105283960"
				},
				{
					 "to": "41793026727",
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 26,
						  "name": "PENDING_ACCEPTED",
						  "description": "Message sent to next instance"
					 },
					 "messageId": "33644485987105283961"
				}
			]
		}
	`)

	var expectedResp models.SendSMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendSmsPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.SendSMSRequest
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

	msgResp, respDetails, err := sms.Send(context.Background(), request)

	require.NoError(t, err)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
