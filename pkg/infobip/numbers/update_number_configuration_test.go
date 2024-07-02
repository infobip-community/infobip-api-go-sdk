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

func TestUpdateNumberConfigurationValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"key": "E9FCDCA496035F08EEA5933702EDF745",
		"keyword": "KEYWORD1",
		"action": {
		  "url": "http://something.com",
		  "httpMethod": "POST",
		  "contentType": "JSON",
		  "type": "HTTP_FORWARD"
		},
		"useConversation": {
		  "enabled": false
		},
		"otherActionsDetails": [],
		"otherActions": []
	  }`)

	var expectedResp models.NumberConfiguration
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	numberKey := "number"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(updateNumberConfigurationsPath, numberKey)))
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

	msgResp, respDetails, err := number.UpdateNumberConfiguration(context.Background(), numberKey,
		models.UpdateNumberConfigurationRequest{
			Key: "E9FCDCA496035F08EEA5933702EDF745",
			Action: &models.ActionConfiguration{
				URL:  "http://something.com",
				Type: "HTTP_FORWARD",
			},
		})

	require.NoError(t, err)
	assert.NotEqual(t, models.NumberConfiguration{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
