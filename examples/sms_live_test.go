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
			{To: "333333333333"},
		},
		From: "Infobip Gopher",
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

func TestGetSMSDeliveryReports(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetSMSDeliveryReportsParams{
		Limit: 10,
	}

	msgResp, respDetails, err := client.SMS.GetDeliveryReports(context.Background(), queryParams)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Results[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.GetSMSDeliveryReportsResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSMSLogs(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetSMSLogsParams{
		From:          "",
		To:            "",
		BulkID:        nil,
		MessageID:     []string{"35167777591103572672", "35158756326003571020"},
		GeneralStatus: "",
		SentSince:     "",
		SentUntil:     "",
		Limit:         10,
		MCC:           "",
		MNC:           "",
	}

	msgResp, respDetails, err := client.SMS.GetLogs(context.Background(), queryParams)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Results[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.GetSMSLogsResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
