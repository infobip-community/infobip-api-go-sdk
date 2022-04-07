package examples

import (
	"context"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test email sending.
func TestSendEmailExample(t *testing.T) {
	apiKey := "6397c2cb673df2c823bc1b9f022067e9-00933425-bd44-402c-94d2-d2ddb66eb0b1"
	baseURL := "3v6kj1.api.infobip.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	mail := models.EmailMsg{
		From:    "edcoronag@selfserviceib.com",
		To:      "edcoronag@gmail.com",
		Subject: "Some subject",
		Text:    "Some text",
	}

	msgResp, respDetails, err := client.Email.Send(context.Background(), mail)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageId, "MessageId should not be empty")
	assert.NotEqual(t, models.SendEmailResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, 200, respDetails.HTTPResponse.StatusCode)
}
