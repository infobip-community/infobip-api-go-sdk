package numbers

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

func TestGetAvailableNumbersValidReq(t *testing.T) {

	rawJSONResp := []byte(`
	{
		"numbers": [
		  {
			"numberKey": "78D8394AC3EG0460B4CF0E723FC31B49",
			"number": "79029555551",
			"country": "RU",
			"type": "VIRTUAL_LONG_NUMBER",
			"capabilities": [
			  "SMS"
			],
			"shared": false,
			"price": {
			  "pricePerMonth": 15,
			  "setupPrice": 0,
			  "initialMonthPrice": 9.193549,
			  "currency": "EUR"
			}
		  },
		  {
			"numberKey": "3B9D1EACAB7FBDRN%EE03592BFCD6BE",
			"number": "79029555525",
			"country": "RU",
			"type": "VIRTUAL_LONG_NUMBER",
			"capabilities": [
			  "SMS"
			],
			"shared": false,
			"price": {
			  "pricePerMonth": 15,
			  "setupPrice": 0,
			  "initialMonthPrice": 9.193549,
			  "currency": "EUR"
			}
		  }
		],
		"numberCount": 2
	  }`)

	queryParams := models.GetAvailableNumbersParams{
		Capabilities: []string{"SMS"},
		Country:      "RU",
		Limit:        2,
	}

	expectedParams := "capabilities=SMS&country=RU&limit=2"

	var expectedResp models.GetAvailableNumbersResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, getAvailableNumbersPath))
		assert.Equal(t, expectedParams, r.URL.RawQuery)
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	number := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := number.GetAvailableNumbers(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetAvailableNumbersResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
