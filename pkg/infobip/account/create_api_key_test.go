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

func TestCreateAPIKeyValidReq(t *testing.T) {
	rawJSONResp := []byte(`
	{
		"accountId": "8F0792F86035A9F4290821F1EE6BC06A",
		"name": "First ApiKey on my account",
		"allowedIPs": [
		  "127.0.0.1",
		  "168.158.10.122"
		],
		"validFrom": "2023-09-01T10:00:00",
		"validTo": "2024-09-01T10:00:00",
		"permissions": [
		  "PUBLIC_API"
		],
		"scopeGuids": [
		  "account:management",
		  "2fa:manage"
		]
	}`)

	var expectedResp models.APIKey
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, createAPIKeyPath))
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

	createAPIKey := models.APIKey{
		AccountID:  "8F0792F86035A9F4290821F1EE6BC06A",
		Name:       "First ApiKey on my account",
		AllowedIPs: []string{"127.0.0.1", "168.158.10.122"},
	}

	msgResp, respDetails, err := account.CreateAPIKey(context.Background(), createAPIKey)

	require.NoError(t, err)
	assert.NotEqual(t, models.APIKey{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
