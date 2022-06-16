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

func TestGetInboundMessagesValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		{
			"results": [
				{
					 "messageId": "817790313235066447",
					 "from": "385916242493",
					 "to": "385921004026",
					 "text": "QUIZ Correct answer is Paris",
					 "cleanText": "Correct answer is Paris",
					 "keyword": "QUIZ",
					 "receivedAt": "2019-11-09T16:00:00.0000000+00:00",
					 "smsCount": 1,
					 "price": {
						  "pricePerMessage": 0,
						  "currency": "EUR"
					 },
					 "callbackData": "callbackData"
				}
			],
			"messageCount": 1,
			"pendingMessageCount": 0
		}
	`)

	apiKey := "some-api-key"
	var expectedResp models.GetInboundSMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, getInboundSMSPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	queryParams := models.GetInboundSMSParams{
		Limit: 1,
	}

	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := sms.GetInboundMessages(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetInboundSMSResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
