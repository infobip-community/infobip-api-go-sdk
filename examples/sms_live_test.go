package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//const (
//	apiKey     = "your-api-key"
//	baseURL    = "your-base-url"
//  destNumber = "123456789012"
//)

func TestSendSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	sms := models.SMSMsg{
		Destinations: []models.SMSDestination{
			{To: destNumber},
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
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSendSMSBulk(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	sms := models.SMSMsg{
		Destinations: []models.SMSDestination{
			{To: destNumber},
		},
		From:   "Infobip Gopher",
		Text:   "Hello from Go SDK",
		SendAt: "2022-06-02T16:00:00.000+0000",
	}
	sms2 := models.SMSMsg{
		Destinations: []models.SMSDestination{
			{To: destNumber},
		},
		From: "Infobip Gopher",
		Text: "Hello (2) from Go SDK",
	}
	request := models.SendSMSRequest{
		BulkID:   "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985",
		Messages: []models.SMSMsg{sms, sms2},
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

func TestSendBinarySMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	binSMS := models.BinarySMSMsg{
		Destinations: []models.SMSDestination{
			{To: destNumber},
		},
		From:   "Infobip Gopher",
		Binary: &models.SMSBinary{Hex: "0f c2 4a bf 34 13 ba"},
	}
	request := models.SendBinarySMSRequest{
		Messages: []models.BinarySMSMsg{binSMS},
	}

	msgResp, respDetails, err := client.SMS.SendBinary(context.Background(), request)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendSMSResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSendSMSOverQueryParameters(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	paramsSMS := models.SendSMSOverQueryParamsParams{
		Username: "your-username",
		Password: "your-password",
		From:     "Infobip Gopher",
		To:       []string{destNumber},
		Text:     "Hello from Go SDK",
	}

	msgResp, respDetails, err := client.SMS.SendOverQueryParams(context.Background(), paramsSMS)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestPreviewSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	previewReq := models.PreviewSMSRequest{
		LanguageCode: "TR",
		Text:         "Mesaj gönderimi yapmadan önce önizleme seçeneğini kullanmanız doğru karar vermenize olur.",
	}

	msgResp, respDetails, err := client.SMS.Preview(context.Background(), previewReq)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.Previews[0].TextPreview, "TextPreview should not be empty")
	assert.NotEqual(t, models.PreviewSMSResponse{}, msgResp)
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
		Limit: 10,
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

func TestGetScheduledSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetScheduledSMSParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}

	msgResp, respDetails, err := client.SMS.GetScheduledMessages(context.Background(), queryParams)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestRescheduleSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.RescheduleSMSParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}
	req := models.RescheduleSMSRequest{SendAt: "2022-06-01T16:00:00.000+0000"}

	msgResp, respDetails, err := client.SMS.RescheduleMessages(context.Background(), req, params)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetScheduledSMSStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.GetScheduledSMSStatusParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}

	msgResp, respDetails, err := client.SMS.GetScheduledMessagesStatus(context.Background(), params)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateScheduledSMSStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.UpdateScheduledSMSStatusParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}
	req := models.UpdateScheduledSMSStatusRequest{Status: "CANCELED"}

	msgResp, respDetails, err := client.SMS.UpdateScheduledMessagesStatus(context.Background(), req, params)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAApplications(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	msgResp, respDetails, err := client.SMS.GetTFAApplications(context.Background())

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp[0].ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAApplicationsResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestCreateTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.CreateTFAApplicationRequest{
		Name: "Test Go TFA App 2",
	}

	msgResp, respDetails, err := client.SMS.CreateTFAApplication(context.Background(), req)

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.CreateTFAApplicationResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	msgResp, respDetails, err := client.SMS.GetTFAApplication(context.Background(), "43D78365E3257420D78752A62845A8CB")

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAApplicationResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.UpdateTFAApplicationRequest{
		Name:    "Test Go TFA App 3",
		Enabled: true,
		Configuration: &models.TFAApplicationConfiguration{
			PinAttempts:                   6,
			AllowMultiplePINVerifications: true,
			PINTimeToLive:                 "42m",
			VerifyPINLimit:                "2/4s",
			SendPINPerApplicationLimit:    "5000/12h",
			SendPINPerPhoneNumberLimit:    "2/1d",
		},
	}

	msgResp, respDetails, err := client.SMS.UpdateTFAApplication(context.Background(), req, "43D78365E3257420D78752A62845A8CB")

	fmt.Println(msgResp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, msgResp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.UpdateTFAApplicationResponse{}, msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, msgResp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
