package mms

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

func TestGetOutboundMsgDeliveryReportsValidReq(t *testing.T) {
	apiKey := "secret"
	rawJSONResp := []byte(`{
		"results": [
			{
				"bulkId": "string",
				"messageId": "string",
				"to": "string",
				"from": "string",
				"sentAt": "string",
				"doneAt": "string",
				"mmsCount": 0,
				"mccMnc": "string",
				"callbackData": "string",
				"price": {
					"pricePerMessage": 0,
					"currency": "string"
				},
				"status": {
					"groupId": 1,
					"groupName": "PENDING",
					"id": 26,
					"name": "PENDING_ACCEPTED",
					"description": "Message accepted, pending for delivery."
				},
				"error": {
					"groupId": 0,
					"groupName": "string",
					"id": 0,
					"name": "string",
					"description": "string"
				}
			}
		]	
	}`)
	var expectedResp models.GetOutboundMMSDeliveryReportsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	tests := []struct {
		scenario       string
		params         models.GetMMSDeliveryReportsParams
		expectedParams string
	}{
		{scenario: "No params passed", params: models.GetMMSDeliveryReportsParams{}, expectedParams: ""},
		{
			scenario: "Params passed",
			params: models.GetMMSDeliveryReportsParams{
				BulkID:    "1",
				MessageID: "2",
				Limit:     3,
			},
			expectedParams: "bulkId=1&limit=3&messageId=2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.scenario, func(t *testing.T) {
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.True(t, strings.HasSuffix(r.URL.Path, getOutboundMMSDeliveryReportsPath))
				assert.Equal(t, tc.expectedParams, r.URL.RawQuery)
				assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

				_, servErr := w.Write(rawJSONResp)
				assert.Nil(t, servErr)
			}))
			defer serv.Close()
			mms := Channel{ReqHandler: internal.HTTPHandler{
				HTTPClient: http.Client{},
				BaseURL:    serv.URL,
				APIKey:     apiKey,
			}}

			var msgResp models.GetOutboundMMSDeliveryReportsResponse
			var respDetails models.ResponseDetails
			msgResp, respDetails, err = mms.GetDeliveryReports(
				context.Background(),
				tc.params,
			)

			require.NoError(t, err)
			assert.NotEqual(t, models.GetOutboundMMSDeliveryReportsResponse{}, msgResp)
			assert.Equal(t, expectedResp, msgResp)
			assert.NotNil(t, respDetails)
			assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
		})
	}
}
