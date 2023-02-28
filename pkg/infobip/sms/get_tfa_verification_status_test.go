package sms

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

func TestGetTFAVerificationStatusValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		{
		  "verifications": [
			{
			  "msisdn": "41793026727",
			  "verified": true,
			  "verifiedAt": 1418364366,
			  "sentAt": 1418364246
			},
			{
			  "msisdn": "41793026746",
			  "verified": false,
			  "verifiedAt": 1418364226,
			  "sentAt": 1418333246
			}
		  ]
		}
	`)

	queryParams := models.GetTFAVerificationStatusParams{
		MSISDN:   "555555555",
		Verified: false,
		Sent:     false,
	}

	expectedParams := "msisdn=555555555&sent=false&verified=false"

	var expectedResp models.GetTFAVerificationStatusResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, getTFAVerificationStatusPath))
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

	msgResp, respDetails, err := sms.GetTFAVerificationStatus(context.Background(), "0933F3BC087D2A617AC6DCB2EF5B8A61", queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetTFAVerificationStatusResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
