package infobip

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultClient(t *testing.T) {
	apiKey := "secret"
	baseURL := "https:/something.api.infobip.com"
	client, err := NewClient(baseURL, apiKey)
	require.Nil(t, err)
	assert.Equal(t, http.Client{}, client.httpClient)

	whatsApp := client.WhatsApp()
	assert.Equal(t, http.Client{}, whatsApp.(*whatsAppChannel).reqHandler.httpClient)
	assert.Equal(t, apiKey, whatsApp.(*whatsAppChannel).reqHandler.apiKey)
	assert.Equal(t, baseURL, whatsApp.(*whatsAppChannel).reqHandler.baseURL)
}

func TestClientWithOptions(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://k31ke1.api.infobip.com"
	customClient := http.Client{Timeout: 3 * time.Second}
	client, err := NewClient(baseURL, apiKey, WithHTTPClient(customClient))
	require.Nil(t, err)
	assert.Equal(t, customClient, client.httpClient)
	assert.Equal(t, customClient.Timeout, client.httpClient.Timeout)

	whatsApp := client.WhatsApp()
	assert.Equal(t, customClient, whatsApp.(*whatsAppChannel).reqHandler.httpClient)
	assert.Equal(t, customClient.Timeout, whatsApp.(*whatsAppChannel).reqHandler.httpClient.Timeout)
}

func TestClientMissingScheme(t *testing.T) {
	apiKey := "secret"
	baseURL := "test.com"
	client, err := NewClient(baseURL, apiKey)
	require.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("https://%s", baseURL), client.baseURL)
}

func TestClientInvalidURL(t *testing.T) {
	apiKey := "secret"
	baseURL := "\\"
	client, err := NewClient(baseURL, apiKey)
	require.NotNil(t, err)
	assert.Equal(t, Client{}, client)
}
