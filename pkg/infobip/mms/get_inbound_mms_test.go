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

func TestGetInboundMMSValidReq(t *testing.T) {
	apiKey := "secret"
	rawJSONResp := []byte(`{
		"results": [
			{
				"messageId": "string",
				"to": "string",
				"from": "string",
				"message": "string",
				"receivedAt": "string",
				"mmsCount": 0,
				"callbackData": "string",
				"price": {
						"pricePerMessage": 0,
						"currency": "string"
				}
			}
		]
	}`)
	var expectedResp models.GetInboundMMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	tests := []struct {
		scenario       string
		params         models.GetInboundMMSParams
		expectedParams string
	}{
		{scenario: "No params passed", params: models.GetInboundMMSParams{}, expectedParams: ""},
		{
			scenario: "Params passed",
			params: models.GetInboundMMSParams{
				Limit: 1,
			},
			expectedParams: "limit=1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.scenario, func(t *testing.T) {
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.True(t, strings.HasSuffix(r.URL.Path, getInboundMMSPath))
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

			var msgResp models.GetInboundMMSResponse
			var respDetails models.ResponseDetails
			msgResp, respDetails, err = mms.GetInboundMessages(
				context.Background(),
				tc.params,
			)

			require.NoError(t, err)
			assert.NotEqual(t, models.GetInboundMMSResponse{}, msgResp)
			assert.Equal(t, expectedResp, msgResp)
			assert.NotNil(t, respDetails)
			assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
		})
	}
}
