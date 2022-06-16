package rcs

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

func TestSendMessageValidReq(t *testing.T) {
	apiKey := "some-key"
	msg := models.GenerateRCSFileMsg()
	rawJSONResp := []byte(`
		{
		  "messages": [
			{
			  "to": "385977666618",
			  "messageCount": 1,
			  "messageId": "06df139a-7eb5-4a6e-902e-40e892210455",
			  "status": {
				"groupId": 1,
				"groupName": "PENDING",
				"id": 7,
				"name": "PENDING_ENROUTE",
				"description": "Message sent to next instance",
				"action": "string"
			  }
			}
		  ]
		}
	`)

	var expectedResp models.SendRCSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendRCSPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.RCSMsg
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, msg)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	rcs := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := rcs.Send(context.Background(), msg)

	require.NoError(t, err)
	assert.NotEqual(t, models.SendRCSResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
