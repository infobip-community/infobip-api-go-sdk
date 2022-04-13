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
	apiKey  = "YOURKEY"
	baseURL = "YOURURL.api.infobip.com"
)

func TestSendEmail(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	mail := models.EmailMsg{
		From:    "somemail@somedomain.com",
		To:      "somemail@somedomain.com",
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

func TestSendEmailBulk(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	mail := models.EmailMsg{
		From:    "@selfserviceib.com",
		To:      "@gmail.com",
		Subject: "Some subject",
		Text:    "Some text",
		BulkId:  "test-bulk-78",
		SendAt:  "2022-04-13T11:35:39.214+00:00",
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
	queryParams["bulkId"] = "test-bulk-78"

	bulks, respDetails, err := client.Email.GetSentBulks(context.Background(), queryParams)

	fmt.Println(bulks)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.SentEmailBulksResponse{}, bulks)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSentBulksStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["bulkId"] = "test-bulk-78"

	bulks, respDetails, err := client.Email.GetSentBulksStatus(context.Background(), queryParams)

	fmt.Println(bulks)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.SentEmailBulksStatusResponse{}, bulks)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestRescheduleMessages(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["bulkId"] = "test-bulk-78"

	req := models.RescheduleMessagesRequest{
		SendAt: "2022-04-13T17:56:07Z",
	}

	rescheduleResp, respDetails, err := client.Email.RescheduleMessages(context.Background(), req, queryParams)

	fmt.Println(rescheduleResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.RescheduleMessagesResponse{}, rescheduleResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateScheduledMessagesStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := make(map[string]string)
	queryParams["bulkId"] = "test-bulk-78"

	req := models.UpdateScheduledMessagesStatusRequest{
		Status: "CANCELED",
	}

	updateResp, respDetails, err := client.Email.UpdateScheduledMessagesStatus(context.Background(), req, queryParams)

	fmt.Println(updateResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.UpdateScheduledMessagesStatusResponse{}, updateResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestValidateAddresses(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.ValidateAddressesRequest{
		To: "somemail@domain.com",
	}

	validateResp, respDetails, err := client.Email.ValidateAddresses(context.Background(), req)

	fmt.Println(validateResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.ValidateAddressesResponse{}, validateResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
