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

func TestSendBinaryMessageValidReq(t *testing.T) {
	request := models.GenerateSendBinarySMSRequest()

	rawJSONResp := []byte(`
		{
			"bulkId": "test-bulk-509",
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
					 "messageId": "test-message-361"
				},
				{
					 "to": "+41793026727",
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 26,
						  "name": "PENDING_ACCEPTED",
						  "description": "Message sent to next instance"
					 },
					 "messageId": "33644717836103574271"
				},
				{
					 "to": "+41793026727",
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 26,
						  "name": "PENDING_ACCEPTED",
						  "description": "Message sent to next instance"
					 },
					 "messageId": "33644717836103574272"
				}
			]
		}
	`)

	var expectedResp models.SendBinarySMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "apiKey"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendBinarySMSPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := io.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.SendBinarySMSRequest
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, models.GenerateSendBinarySMSRequest())

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := sms.SendBinary(context.Background(), request)

	require.NoError(t, err)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
