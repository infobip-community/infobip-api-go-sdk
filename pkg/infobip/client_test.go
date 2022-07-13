package infobip

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/mms"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/whatsapp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultClient(t *testing.T) {
	apiKey := "secret"
	baseURL := "https:/something.api.infobip.com"
	client, err := NewClient(baseURL, apiKey)
	require.NoError(t, err)
	assert.Equal(t, http.Client{}, client.httpClient)

	whatsApp := client.WhatsApp
	assert.Equal(t, http.Client{}, whatsApp.(*whatsapp.Channel).ReqHandler.HTTPClient)
	assert.Equal(t, apiKey, whatsApp.(*whatsapp.Channel).ReqHandler.APIKey)
	assert.Equal(t, baseURL, whatsApp.(*whatsapp.Channel).ReqHandler.BaseURL)

	mmsChannel := client.MMS
	assert.Equal(t, http.Client{}, mmsChannel.(*mms.Channel).ReqHandler.HTTPClient)
	assert.Equal(t, apiKey, mmsChannel.(*mms.Channel).ReqHandler.APIKey)
	assert.Equal(t, baseURL, mmsChannel.(*mms.Channel).ReqHandler.BaseURL)
}

func TestClientWithOptions(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://k31ke1.api.infobip.com"
	customClient := http.Client{Timeout: 3 * time.Second}
	client, err := NewClient(baseURL, apiKey, WithHTTPClient(customClient))
	require.NoError(t, err)
	assert.Equal(t, customClient, client.httpClient)
	assert.Equal(t, customClient.Timeout, client.httpClient.Timeout)

	whatsApp := client.WhatsApp
	assert.Equal(t, customClient, whatsApp.(*whatsapp.Channel).ReqHandler.HTTPClient)
	assert.Equal(t, customClient.Timeout, whatsApp.(*whatsapp.Channel).ReqHandler.HTTPClient.Timeout)

	mmsChannel := client.MMS
	assert.Equal(t, customClient, mmsChannel.(*mms.Channel).ReqHandler.HTTPClient)
	assert.Equal(t, customClient.Timeout, mmsChannel.(*mms.Channel).ReqHandler.HTTPClient.Timeout)
}

func TestClientMissingScheme(t *testing.T) {
	apiKey := "secret"
	baseURL := "test.com"
	client, err := NewClient(baseURL, apiKey)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("https://%s", baseURL), client.baseURL)
}

func TestClientInvalidURL(t *testing.T) {
	apiKey := "secret"
	baseURL := "\\"
	client, err := NewClient(baseURL, apiKey)
	require.NotNil(t, err)
	assert.Equal(t, Client{}, client)
}
