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

func TestGetNumberConfigurationsValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"configurations": [
		  {
			"key": "6336C2CCF10E74B705340E70D8E06BD6",
			"keyword": "KEYWORD1",
			"action": {
			  "type": "HTTP_FORWARD",
			  "url": "http://something.com"
			},
			"otherActionsDetails": [
			  {
				"message": "Auto response message text.",
				"editable": true,
				"type": "AUTORESPONSE"
			  }
			],
			"otherActions": [
			  "AUTORESPONSE"
			]
		  },
		  {
			"key": "8F0792F86035A9F4290821F1EE6BC06A",
			"keyword": "KEYWORD2",
			"action": {
			  "type": "MAIL_FORWARD",
			  "mail": "someone@something.com"
			},
			"otherActionsDetails": [
			  {
				"blockType": "FROM_SENDER",
				"editable": true,
				"type": "BLOCK"
			  }
			],
			"otherActions": [
			  "BLOCK"
			]
		  }
		],
		"totalCount": 2
	  }`)

	queryParams := models.GetAllNumberConfigurationParam{
		Limit: 2,
	}

	expectedParams := "limit=2"

	var expectedResp models.GetAllNumberConfigurationResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	numberKey := ""

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(getAllNumberConfigurationsPath, numberKey)))
		assert.Equal(t, expectedParams, r.URL.RawQuery)
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

	msgResp, respDetails, err := number.GetAllNumberConfigurations(context.Background(), numberKey, queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetAllNumberConfigurationResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
