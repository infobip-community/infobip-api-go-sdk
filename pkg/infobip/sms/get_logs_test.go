package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLogsValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		{
			"results": [
				{
					 "bulkId": "33649386024303572262",
					 "messageId": "33649386024303572263",
					 "to": "41793026727",
					 "from": "InfoSMS",
					 "text": "This is a sample message 1",
					 "sentAt": "2021-11-09T21:37:40.258+0000",
					 "doneAt": "2021-11-09T21:37:40.602+0000",
					 "smsCount": 1,
					 "mccMnc": "21910",
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
					 "error": {
						  "groupId": 0,
						  "groupName": "OK",
						  "id": 0,
						  "name": "NO_ERROR",
						  "description": "No Error",
						  "permanent": false
					 }
				}
			]
		}
	`)

	apiKey := "some-api-key"
	var expectedResp models.GetSMSLogsResponse
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

	queryParams := models.GetSMSLogsParams{
		From:      "someone",
		To:        "123456789",
		BulkID:    []string{"some-bulk-id", "some-bulk-id-2"},
		MessageID: []string{"764037642753", "764037642754"},
		Limit:     1,
	}

	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := sms.GetLogs(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetSMSLogsResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
