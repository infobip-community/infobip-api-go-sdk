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

func TestSendSMSOverQueryParamsValidReq(t *testing.T) {
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
		  "bulkId": "1478260834465349756",
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
			  "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"
			}
		  ]
		}	
	`)

	var expectedResp models.SendSMSOverQueryParamsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, sendSMSOverQueryParamsPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	queryParams := models.SendSMSOverQueryParamsParams{
		Username:           "username",
		Password:           "pass",
		BulkID:             "bulk-12503542",
		From:               "someone",
		To:                 []string{"41793026727", "41793026728"},
		Text:               "some text",
		Flash:              false,
		Transliteration:    "something",
		LanguageCode:       "EN",
		IntermediateReport: false,
		NotifyURL:          "http://some.url",
		NotifyContentType:  "application/json",
		CallbackData:       "callback data",
	}

	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := sms.SendOverQueryParams(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.SendSMSOverQueryParamsResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
