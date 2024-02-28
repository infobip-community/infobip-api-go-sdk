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

func TestListPurchasedNumbersValidReq(t *testing.T) {

	rawJSONResp := []byte(`
	{
		"numbers": 	[
		{
			"numberKey": "6FED0BC540BFADD9B05ED7D89AAC22FA",
			"number": "447860041117",
			"country": "GB",
			"countryName": "United Kingdom",
			"type": "VIRTUAL_LONG_NUMBER",
			"capabilities": ["SMS"],
			"shared": false,
			"price": {	
				"pricePerMonth": 5,
				"setupPrice": 0,
				"currency": "EUR"
			},
			"network": "02 (Telefonica UK Ltd)",
			"keywords": [ "test", "stop" ],
			"additionalSetupRequired": false,
			"editPermissions": 
				{
					"canEditNumber": true,
					"canEditConfiguration": true
				},
				"applicationId": "default"
			}
		],
		"numberCount": 1
	}`)

	queryParams := models.ListPurchasedNumbersParam{
		Number: "447860041117",
		Limit:  1,
	}

	expectedParams := "limit=1&number=447860041117"

	var expectedResp models.ListPurchasedNumbersResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, listPurchasedNumbersPath))
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

	msgResp, respDetails, err := number.ListPurchasedNumbers(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.ListPurchasedNumbersResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
