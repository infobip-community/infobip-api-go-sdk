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

func TestUpdateAccountValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"enabled": true,
		"key": "string",
		"name": "string",
		"ownerKey": "string"
	  }`)

	var expectedResp models.UpdateAccountResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	accountKey := "string"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(updateAccountPath, accountKey)))
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

	msgResp, respDetails, err := account.UpdateAccount(context.Background(), accountKey, models.UpdateAccountRequest{})

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateAccountResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
