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

const (
	destNumber = "555555555555"
)

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

	resp, respDetails, err := client.SMS.Send(context.Background(), request)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendSMSResponse{}, resp)
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

	resp, respDetails, err := client.SMS.Send(context.Background(), request)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendSMSResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
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

	resp, respDetails, err := client.SMS.SendBinary(context.Background(), request)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.SendSMSResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
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

	resp, respDetails, err := client.SMS.SendOverQueryParams(context.Background(), paramsSMS)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Messages[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestPreviewSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	previewReq := models.PreviewSMSRequest{
		LanguageCode: "TR",
		Text:         "Mesaj gönderimi yapmadan önce önizleme seçeneğini kullanmanız doğru karar vermenize olur.",
	}

	resp, respDetails, err := client.SMS.Preview(context.Background(), previewReq)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Previews[0].TextPreview, "TextPreview should not be empty")
	assert.NotEqual(t, models.PreviewSMSResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSMSDeliveryReports(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetSMSDeliveryReportsParams{
		Limit: 10,
	}

	resp, respDetails, err := client.SMS.GetDeliveryReports(context.Background(), queryParams)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Results[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.GetSMSDeliveryReportsResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetSMSLogs(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetSMSLogsParams{
		Limit: 10,
	}

	resp, respDetails, err := client.SMS.GetLogs(context.Background(), queryParams)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Results[0].MessageID, "MessageID should not be empty")
	assert.NotEqual(t, models.GetSMSLogsResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetScheduledSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	queryParams := models.GetScheduledSMSParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}

	resp, respDetails, err := client.SMS.GetScheduledMessages(context.Background(), queryParams)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestRescheduleSMS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.RescheduleSMSParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}
	req := models.RescheduleSMSRequest{SendAt: "2022-06-01T16:00:00.000+0000"}

	resp, respDetails, err := client.SMS.RescheduleMessages(context.Background(), req, params)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetScheduledSMSStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.GetScheduledSMSStatusParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}

	resp, respDetails, err := client.SMS.GetScheduledMessagesStatus(context.Background(), params)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateScheduledSMSStatus(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	params := models.UpdateScheduledSMSStatusParams{BulkID: "f4b07b1a-a009-49d5-a94d-f8fd1bfdc985"}
	req := models.UpdateScheduledSMSStatusRequest{Status: "CANCELED"}

	resp, respDetails, err := client.SMS.UpdateScheduledMessagesStatus(context.Background(), req, params)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.BulkID, "BulkID should not be empty")
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAApplications(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	resp, respDetails, err := client.SMS.GetTFAApplications(context.Background())

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp[0].ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAApplicationsResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestCreateTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.CreateTFAApplicationRequest{
		Name: "Test Go TFA App 2",
	}

	resp, respDetails, err := client.SMS.CreateTFAApplication(context.Background(), req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.CreateTFAApplicationResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	resp, respDetails, err := client.SMS.GetTFAApplication(context.Background(), "43D78365E3257420D78752A62845A8CB")

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAApplicationResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateTFAApplication(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.UpdateTFAApplicationRequest{
		Name:    "Test Go TFA App 3",
		Enabled: true,
		Configuration: &models.TFAApplicationConfiguration{
			PINAttempts:                   6,
			AllowMultiplePINVerifications: true,
			PINTimeToLive:                 "20m",
			VerifyPINLimit:                "2/4s",
			SendPINPerApplicationLimit:    "5000/12h",
			SendPINPerPhoneNumberLimit:    "2/1d",
		},
	}

	resp, respDetails, err := client.SMS.UpdateTFAApplication(context.Background(), "43D78365E3257420D78752A62845A8CB", req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.ApplicationID, "ID should not be empty")
	assert.NotEqual(t, models.UpdateTFAApplicationResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAMessageTemplates(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	resp, respDetails, err := client.SMS.GetTFAMessageTemplates(context.Background(), "43D78365E3257420D78752A62845A8CB")

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp[0].MessageID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAMessageTemplatesResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestCreateTFAMessageTemplate(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.CreateTFAMessageTemplateRequest{
		MessageText: "The third verification code is {pin}",
		PINType:     "NUMERIC",
		PINLength:   4,
		Language:    models.English,
	}

	resp, respDetails, err := client.SMS.CreateTFAMessageTemplate(context.Background(), "18C0684EB9CC244072CB31E284A45707", req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.MessageID, "ID should not be empty")
	assert.NotEqual(t, models.CreateTFAMessageTemplateResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAMessageTemplate(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	resp, respDetails, err := client.SMS.GetTFAMessageTemplate(context.Background(), "43D78365E3257420D78752A62845A8CB", "9AD26BD115AB45657A0FEACACCC918BE")

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.MessageID, "ID should not be empty")
	assert.NotEqual(t, models.GetTFAMessageTemplateResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateTFAMessageTemplate(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.UpdateTFAMessageTemplateRequest{
		MessageText:    "Hello {{name}} the PIN is {{pin}}",
		PINType:        "NUMERIC",
		PINPlaceholder: "{{pin}}",
		PINLength:      4,
		Language:       models.English,
	}

	resp, respDetails, err := client.SMS.UpdateTFAMessageTemplate(context.Background(), "43D78365E3257420D78752A62845A8CB", "9AD26BD115AB45657A0FEACACCC918BE", req)

	fmt.Printf("%+v", resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.MessageID, "ID should not be empty")
	assert.NotEqual(t, models.UpdateTFAMessageTemplateResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSendPINOverSMS(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	params := models.SendPINOverSMSParams{NCNeeded: false}
	req := models.SendPINOverSMSRequest{
		ApplicationID: "43D78365E3257420D78752A62845A8CB",
		MessageID:     "9AD26BD115AB45657A0FEACACCC918BE",
		To:            destNumber,
		Placeholders: map[string]string{
			"name": "John",
		},
	}

	resp, respDetails, err := client.SMS.SendPINOverSMS(context.Background(), params, req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.PINID, "ID should not be empty")
	assert.NotEqual(t, models.SendPINOverSMSResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestResendPINOverSMS(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.ResendPINOverSMSRequest{
		Placeholders: map[string]string{"name": "Steve"},
	}

	resp, respDetails, err := client.SMS.ResendPINOverSMS(context.Background(), "A787EC9C153328E9D276D98861C9CEA1", req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.PINID, "ID should not be empty")
	assert.NotEqual(t, models.ResendPINOverSMSResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSendPINOverVoice(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.SendPINOverVoiceRequest{
		ApplicationID: "43D78365E3257420D78752A62845A8CB",
		MessageID:     "9AD26BD115AB45657A0FEACACCC918BE",
		To:            destNumber,
		Placeholders: map[string]string{
			"name": "John",
		},
	}

	resp, respDetails, err := client.SMS.SendPINOverVoice(context.Background(), req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.PINID, "ID should not be empty")
	assert.NotEqual(t, models.SendPINOverVoiceResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestResendPINOverVoice(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	req := models.ResendPINOverVoiceRequest{
		Placeholders: map[string]string{"name": "Steve"},
	}

	resp, respDetails, err := client.SMS.ResendPINOverVoice(context.Background(), "A787EC9C153328E9D276D98861C9CEA1", req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.PINID, "ID should not be empty")
	assert.NotEqual(t, models.ResendPINOverVoiceResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestVerifyPhoneNumber(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	pinID := "A787EC9C153328E9D276D98861C9CEA1"
	req := models.VerifyPhoneNumberRequest{
		PIN: "3240",
	}

	resp, respDetails, err := client.SMS.VerifyPhoneNumber(context.Background(), pinID, req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.PINID, "ID should not be empty")
	assert.NotEqual(t, models.VerifyPhoneNumberResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetTFAVerificationStatus(t *testing.T) {
	client, err := infobip.NewClientFromEnv()
	require.Nil(t, err)

	appID := "43D78365E3257420D78752A62845A8CB"
	params := models.GetTFAVerificationStatusParams{
		MSISDN:   destNumber,
		Verified: false,
	}
	resp, respDetails, err := client.SMS.GetTFAVerificationStatus(context.Background(), appID, params)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Verifications, "Verifications should not be empty")
	assert.NotEqual(t, models.GetTFAVerificationStatusResponse{}, resp)
	assert.NotEqual(t, models.ResponseDetails{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
