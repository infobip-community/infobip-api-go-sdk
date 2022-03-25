package examples

import (
	"context"
	"fmt"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test email sending.
func TestSendEmailExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	mail := models.EmailMsg{
		From:                    "someone@infobip.com",
		To:                      "ecorona@infobip.com",
		Subject:                 "Test email",
		Text:                    "Test email body",
	}

	msgResp, err := client.Email.Send(context.Background(), mail)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}
