package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	apiKey  = "you-api-key"
	baseURL = "your-base-url"
)

func TestSendSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	sms := models.SMSMsg{
		Destinations: []models.Destination{
			models.Destination{To: "523311800428"},
		},
		From: "Gopher",
		Text: "Hello from Go SDK",
	}
	request := models.SendSMSRequest{
		Messages: []models.SMSMsg{sms},
	}

	msgResp, respDetails, err := client.SMS.Send(context.Background(), request)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendSMSResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
