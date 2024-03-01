package account

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

func TestGetAllAccountsValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"accounts": [
		  {
			"key": "8F0792F86035A9F4290821F1EE6BC06A",
			"ownerKey": "37B93F4D2BA3C58B58526EAEAA1AB35C",
			"name": "First account",
			"enabled": true
		  },
		  {
			"key": "6298AA7707903A4ED680B436929681AD",
			"ownerKey": "37B93F4D2BA3C58B58526EAEAA1AB35C",
			"name": "Second account",
			"enabled": false
		  }
		]
	  }`)

	queryParam := models.GetAllAccountsParams{
		Limit: 2,
	}

	expectedParams := "limit=2"

	var expectedResp models.GetAllAccountsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, getAllAccountsPath))
		assert.Equal(t, expectedParams, r.URL.RawQuery)
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	account := Platform{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := account.GetAllAccounts(context.Background(), queryParam)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetAllAccountsResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
