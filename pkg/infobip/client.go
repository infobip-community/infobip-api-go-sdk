// Package infobip provides a client library
// for interacting with the Infobip API.
// https://www.infobip.com/docs/api
package infobip

import (
	"net/http"
	"net/url"

	"github.com/pgrubacc/infobip-go-client/internal"

	"github.com/pgrubacc/infobip-go-client/pkg/infobip/whatsapp"
)

// Client is the entrypoint to all Infobip channels.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient http.Client
	WhatsApp   whatsapp.WhatsApp
}

// NewClient returns a client object using the provided baseURL and apiKey.
// If a client is not provided using options, a default one is created.
func NewClient(baseURL string, apiKey string, options ...func(*Client)) (Client, error) {
	baseURL, err := validateURL(baseURL)
	if err != nil {
		return Client{}, err
	}
	c := Client{baseURL: baseURL, apiKey: apiKey, httpClient: http.Client{}}

	for _, opt := range options {
		opt(&c)
	}

	c.WhatsApp = &whatsapp.Channel{
		ReqHandler: internal.HTTPHandler{APIKey: apiKey, BaseURL: baseURL, HTTPClient: c.httpClient},
	}

	return c, nil
}

func validateURL(baseURL string) (string, error) {
	_, err := url.ParseRequestURI(baseURL)
	if err != nil {
		baseURL = "https://" + baseURL
		_, err = url.ParseRequestURI(baseURL)
	}

	return baseURL, err
}

func WithHTTPClient(httpClient http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
