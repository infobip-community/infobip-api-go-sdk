package whatsapp

import (
	"context"
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

func TestDeleteWATemplateValidReq(t *testing.T) {
	apiKey := "some-api-key"
	sender := "111111111111"
	templateName := "template-name"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, fmt.Sprintf(deleteTemplatePath, sender, templateName)))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer serv.Close()

	whatsapp := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	respDetails, err := whatsapp.DeleteTemplate(context.Background(), sender, templateName)

	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusNoContent, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
