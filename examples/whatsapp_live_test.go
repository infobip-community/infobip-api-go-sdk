package examples

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"infobip-go-client/pkg/infobip"
	"infobip-go-client/pkg/infobip/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSendMessageExample can be used to test the client against a real environment.
// Replace the apiKey, baseURL and From / To fields of the message and run the test.
func TestSendMessageExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	whatsApp := client.WhatsApp()
	message := models.TextMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.TextContent{Text: "This message was sent from the Infobip API using the Go API client."},
	}

	msgResp, respDetails, err := whatsApp.SendTextMessage(context.Background(), message)

	require.Nil(t, err)
	assert.NotEqual(t, models.TextMessageResponse{}, msgResp)
	fmt.Printf("%+v", msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
}
