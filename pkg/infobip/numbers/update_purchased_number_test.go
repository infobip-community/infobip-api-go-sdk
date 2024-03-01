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

func TestUpdatePurchasedNumberValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"numberKey": "58B3840032C7774BAC840EEEA2C23A44",
		"number": "447860041117",
		"country": "GB",
		"type": "VIRTUAL_LONG_NUMBER",
		"capabilities": ["SMS"],
		"shared": false,
		"price": 	{
			"pricePerMonth": 5,
			"setupPrice": 0,
			"currency": "EUR"
		},
		"applicationId": "appId",
		"entityId": "entId"
	}`)

	numberKey := "58B3840032C7774BAC840EEEA2C23A44"

	var expectedResp models.Number
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(updatepurchasedNumberPath, numberKey)))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	number := Platform{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := number.UpdatePurshasedNumbers(context.Background(), numberKey,
		models.UpdatePurchasedNumberRequest{
			ApplicationID: "appId",
			EntityID:      "entId",
		})

	require.NoError(t, err)
	assert.NotEqual(t, models.Number{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
