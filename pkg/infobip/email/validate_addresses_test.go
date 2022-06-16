package email

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateAddressesValidReq(t *testing.T) {
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
			"to": "joan.doe0@example.com",
			"validMailbox": "true",
			"validSyntax": true,
			"catchAll": false,
			"disposable": false,
			"roleBased": false
		}
	`)

	var expectedResp models.ValidateEmailAddressesResponse

	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, validateAddressesPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	req := models.ValidateEmailAddressesRequest{
		To: "someone@infobip.com",
	}

	resp, respDetails, err := email.ValidateAddresses(context.Background(), req)

	require.NoError(t, err)
	assert.NotEqual(t, models.ValidateEmailAddressesResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
