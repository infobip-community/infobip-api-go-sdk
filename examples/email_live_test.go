// Package examples provides some real usage examples. Some of them depend on the server state, and need custom configuration.
package examples

import (
	"context"
	"fmt"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

const (
	apiKey  = "secret"
	baseURL = "YOURURL.api.infobip.com"
)

func TestSendEmail(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	mail := models.EmailMsg{
		From:    "edcoronag@selfserviceib.com",
		To:      "edcoronag@gmail.com",
		Subject: "Some subject",
		Text:    "Some text",
	}

	msgResp, respDetails, err := client.Email.Send(context.Background(), mail)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageId, "MessageId should not be empty")
	assert.NotEqual(t, models.SendEmailResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetEmailDeliveryReports(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["bulkId"] = ""
	queryParams["messageId"] = ""
	queryParams["limit"] = "1000"
	deliveryReports, respDetails, err := client.Email.GetDeliveryReports(context.Background(), queryParams)

	fmt.Println(deliveryReports)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, deliveryReports.Results[0].MessageId, "MessageId should not be empty")
	assert.NotEqual(t, models.EmailDeliveryReportsResponse{}, deliveryReports)
	assert.NotEqual(t, models.ResponseDetails{}, deliveryReports)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetLogs(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["messageId"] = ""
	queryParams["from"] = ""
	queryParams["to"] = ""
	queryParams["bulkId"] = ""
	queryParams["generalStatus"] = ""
	queryParams["sentSince"] = ""
	queryParams["sentUntil"] = ""
	queryParams["limit"] = "1000"

	logs, respDetails, err := client.Email.GetLogs(context.Background(), queryParams)

	fmt.Println(logs)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.EmailLogsResponse{}, logs)
	assert.NotEqual(t, models.ResponseDetails{}, logs)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSentBulks(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["bulkId"] = ""

	bulks, respDetails, err := client.Email.GetSentBulks(context.Background(), queryParams)

	fmt.Println(bulks)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.SentEmailBulksResponse{}, bulks)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
