package sms

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

func TestGetDeliveryReportsValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		{
			"results": [
				{
					 "bulkId": "33640241674303572139",
					 "messageId": "33640241674303572142",
					 "to": "41793026727",
					 "from": "InfoSMS",
					 "sentAt": "2021-11-08T20:13:36.768+0000",
					 "doneAt": "2021-11-08T20:13:39.819+0000",
					 "smsCount": 1,
					 "mccMnc": "null",
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
				},
				{
					 "bulkId": "33640241674303572139",
					 "messageId": "33640241674303572141",
					 "to": "41793026727",
					 "from": "InfoSMS",
					 "sentAt": "2021-11-08T20:13:36.749+0000",
					 "doneAt": "2021-11-08T20:13:40.293+0000",
					 "smsCount": 1,
					 "mccMnc": "null",
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
				},
				{
					 "bulkId": "33644485987105283959",
					 "messageId": "33644485987105283960",
					 "to": "41793026727",
					 "from": "InfoSMS",
					 "sentAt": "2021-11-09T08:00:59.881+0000",
					 "doneAt": "2021-11-09T08:01:03.268+0000",
					 "smsCount": 1,
					 "mccMnc": "null",
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
				},
				{
					 "bulkId": "33644485987105283959",
					 "messageId": "33644485987105283961",
					 "to": "41793026727",
					 "from": "InfoSMS",
					 "sentAt": "2021-11-09T08:00:59.880+0000",
					 "doneAt": "2021-11-09T08:01:03.745+0000",
					 "smsCount": 1,
					 "mccMnc": "null",
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

	queryParams := models.GetSMSDeliveryReportsParams{
		BulkID:    "some-bulk-id",
		MessageID: "1",
		Limit:     1,
	}

	expectedParams := "bulkId=some-bulk-id&limit=1&messageId=1"

	var expectedResp models.GetSMSDeliveryReportsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, getDeliveryReportsPath))
		assert.Equal(t, expectedParams, r.URL.RawQuery)
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := sms.GetDeliveryReports(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetSMSDeliveryReportsParams{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
