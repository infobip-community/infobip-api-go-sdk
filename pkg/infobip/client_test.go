package infobip

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultClient(t *testing.T) {
	apiKey := "jlf3cdef5b20acc82019482a2ce463cc-9b3d6g39-8af8-4e9b-9206-e76340dc43e5"
	baseURL := "https://k31ke1.api.infobip.com"
	client := NewClient(baseURL, apiKey)
	assert.Equal(t, http.Client{}, client.httpClient)

	whatsApp := client.WhatsApp()
	assert.Equal(t, http.Client{}, whatsApp.(whatsAppChannel).reqHandler.httpClient)
	assert.Equal(t, apiKey, whatsApp.(whatsAppChannel).reqHandler.apiKey)
	assert.Equal(t, baseURL, whatsApp.(whatsAppChannel).reqHandler.baseURL)
}

func TestClientWithOptions(t *testing.T) {
	apiKey := "jlf3cdef5b20acc82019482a2ce463cc-9b3d6g39-8af8-4e9b-9206-e76340dc43e5"
	baseURL := "https://k31ke1.api.infobip.com"
	customClient := http.Client{Timeout: 3 * time.Second}
	client := NewClient(baseURL, apiKey, WithHttpClient(customClient))
	assert.Equal(t, customClient, client.httpClient)
	assert.Equal(t, customClient.Timeout, client.httpClient.Timeout)

	whatsApp := client.WhatsApp()
	assert.Equal(t, customClient, whatsApp.(whatsAppChannel).reqHandler.httpClient)
	assert.Equal(t, customClient.Timeout, whatsApp.(whatsAppChannel).reqHandler.httpClient.Timeout)
}
