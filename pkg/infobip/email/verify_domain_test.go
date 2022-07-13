package email

import (
	"context"
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

func TestVerifyDomainValidReq(t *testing.T) {
	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		assert.True(t, strings.Contains(r.URL.Path, verifyDomainPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusAccepted)
	}))
	defer serv.Close()

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	respDetails, err := email.VerifyDomain(context.Background(), "test-domain.com")

	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusAccepted, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
