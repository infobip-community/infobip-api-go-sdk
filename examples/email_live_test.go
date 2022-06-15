// Package examples provides some real usage examples.
// Some of them depend on the server state, and need custom configuration.
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
	apiKey  = "your-api-key"
	baseURL = "your-base-url"
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
	assert.NotEmptyf(t, msgResp.Messages[0].MessageID, "MessageID should not be empty")
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
		BulkID:  "test-bulk-78",
		SendAt:  "2022-04-13T11:35:39.214+00:00",
	}

	msgResp, respDetails, err := client.Email.Send(context.Background(), mail)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendEmailResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetEmailDeliveryReports(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetEmailDeliveryReportsParams{
		BulkID:    "",
		MessageID: "",
		Limit:     1,
	}
	deliveryReports, respDetails, err := client.Email.GetDeliveryReports(context.Background(), queryParams)

	fmt.Println(deliveryReports)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, deliveryReports.Results[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.GetEmailDeliveryReportsResponse{}, deliveryReports)
	assert.NotEqual(t, models.ResponseDetails{}, deliveryReports)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetLogs(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetEmailLogsParams{
		MessageID:     "",
		From:          "",
		To:            "",
		BulkID:        "",
		GeneralStatus: "",
		SentSince:     "",
		SentUntil:     "",
		Limit:         1,
	}

	logs, respDetails, err := client.Email.GetLogs(context.Background(), queryParams)

	fmt.Println(logs)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.GetEmailLogsResponse{}, logs)
	assert.NotEqual(t, models.ResponseDetails{}, logs)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSentBulks(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetSentEmailBulksParams{
		BulkID: "test-bulk-78",
	}

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

	queryParams := models.GetSentEmailBulksStatusParams{
		BulkID: "test-bulk-78",
	}

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

	queryParams := models.RescheduleEmailParams{
		BulkID: "test-bulk-78",
	}

	req := models.RescheduleEmailRequest{
		SendAt: "2022-04-13T17:56:07Z",
	}

	rescheduleResp, respDetails, err := client.Email.RescheduleMessages(context.Background(), req, queryParams)

	fmt.Println(rescheduleResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.RescheduleEmailResponse{}, rescheduleResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateScheduledMessagesStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.UpdateScheduledEmailStatusParams{
		BulkID: "test-bulk-78",
	}

	req := models.UpdateScheduledEmailStatusRequest{
		Status: "CANCELED",
	}

	updateResp, respDetails, err := client.Email.UpdateScheduledMessagesStatus(context.Background(), req, queryParams)

	fmt.Println(updateResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.UpdateScheduledStatusResponse{}, updateResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestValidateAddresses(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.ValidateEmailAddressesRequest{
		To: "somemail@domain.com",
	}

	validateResp, respDetails, err := client.Email.ValidateAddresses(context.Background(), req)

	fmt.Println(validateResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.ValidateEmailAddressesResponse{}, validateResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetDomains(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.GetEmailDomainsParams{
		Size: 10,
		Page: 0,
	}
	resp, respDetails, err := client.Email.GetDomains(context.Background(), params)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.GetEmailDomainsResponse{}, resp)
	assert.Greater(t, len(resp.Results), 0)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestAddDomain(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.AddEmailDomainRequest{
		DomainName: "test-domain2.com",
	}

	domain, respDetails, err := client.Email.AddDomain(context.Background(), req)

	fmt.Println(domain)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.AddEmailDomainResponse{}, domain)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetDomain(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	resp, respDetails, err := client.Email.GetDomain(context.Background(), "test-domain.com")

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.GetEmailDomainResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestDeleteDomain(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	respDetails, err := client.Email.DeleteDomain(context.Background(), "test-domain2.com")

	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusNoContent, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateDomainTracking(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.UpdateEmailDomainTrackingRequest{
		Opens:       false,
		Clicks:      true,
		Unsubscribe: false,
	}

	resp, respDetails, err := client.Email.UpdateDomainTracking(context.Background(), "test-domain.com", req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.UpdateEmailDomainTrackingResponse{}, resp)
	assert.Falsef(t, resp.Tracking.Opens, "Opens should be false")
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestVerifyDomain(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	respDetails, err := client.Email.VerifyDomain(context.Background(), "test-domain.com")

	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.Equal(t, http.StatusAccepted, respDetails.HTTPResponse.StatusCode)
}
