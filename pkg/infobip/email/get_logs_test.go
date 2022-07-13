package email

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLogsValidReq(t *testing.T) {
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
			"results": [
				{
					 "messageId": "ykk9rmnnx79kfqk9hbme",
					 "to": "joan.doe0@example.com",
					 "from": "postman@selfserviceib.com",
					 "text": "This is the message text",
					 "sentAt": "2021-11-08T21:39:21.866+0000",
					 "doneAt": "2021-11-08T21:39:22.146+0000",
					 "messageCount": 1,
					 "price": {
						  "pricePerMessage": 0,
						  "currency": "EUR"
					 },
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 3,
						  "name": "PENDING_WAITING_DELIVERY",
						  "description": "Message sent, waiting for delivery report"
					 },
					 "channel": "EMAIL"
				},
				{
					 "messageId": "yvk4d88q7onz73li2edz",
					 "to": "joan.doe0@example.com",
					 "from": "postman@selfserviceib.com",
					 "text": "This is the message text",
					 "sentAt": "2021-11-08T21:39:20.848+0000",
					 "doneAt": "2021-11-08T21:39:21.118+0000",
					 "messageCount": 1,
					 "price": {
						  "pricePerMessage": 0,
						  "currency": "EUR"
					 },
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 3,
						  "name": "PENDING_WAITING_DELIVERY",
						  "description": "Message sent, waiting for delivery report"
					 },
					 "channel": "EMAIL"
				},
				{
					 "messageId": "bybtz598bogn48d1voyj",
					 "to": "joan.doe0@example.com",
					 "from": "postman@selfserviceib.com",
					 "text": "This is the message text",
					 "sentAt": "2021-11-08T21:39:19.801+0000",
					 "doneAt": "2021-11-08T21:39:20.070+0000",
					 "messageCount": 1,
					 "price": {
						  "pricePerMessage": 0,
						  "currency": "EUR"
					 },
					 "status": {
						  "groupId": 3,
						  "groupName": "DELIVERED",
						  "id": 5,
						  "name": "DELIVERED_TO_HANDSET",
						  "description": "Message delivered to handset"
					 },
					 "channel": "EMAIL"
				},
				{
					 "messageId": "jcpgh35t8genyuj3zjdm",
					 "to": "joan.doe0@example.com",
					 "from": "postman@selfserviceib.com",
					 "text": "This is the message text",
					 "sentAt": "2021-11-08T21:39:18.591+0000",
					 "doneAt": "2021-11-08T21:39:18.870+0000",
					 "messageCount": 1,
					 "price": {
						  "pricePerMessage": 0,
						  "currency": "EUR"
					 },
					 "status": {
						  "groupId": 1,
						  "groupName": "PENDING",
						  "id": 3,
						  "name": "PENDING_WAITING_DELIVERY",
						  "description": "Message sent, waiting for delivery report"
					 },
					 "channel": "EMAIL"
				}
			]
		}
	`)

	var expectedResp models.GetEmailLogsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, getLogsPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	queryParams := models.GetEmailLogsParams{
		Limit: 1,
	}

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := email.GetLogs(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetEmailLogsResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
