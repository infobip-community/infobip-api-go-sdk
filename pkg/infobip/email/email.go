package email

import (
	"context"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	sendEmailPath                     = "/email/2/send"
	getDeliveryReportsPath            = "/email/1/reports"
	getLogsPath                       = "/email/1/logs"
	getSentEmailBulksPath             = "/email/1/bulks"
	rescheduleMessagesPath            = "/email/1/bulks"
	getSentEmailBulksStatusPath       = "/email/1/bulks/status"
	updateScheduledMessagesStatusPath = "/email/1/bulks/status"
	validateAddressesPath             = "/email/2/validation"
)

type Channel struct {
	ReqHandler internal.HTTPHandler
}

type Email interface {
	Send(ctx context.Context, request models.EmailMsg) (response models.SendEmailResponse, responseDetails models.ResponseDetails, err error)
	GetDeliveryReports(ctx context.Context, queryParams map[string]string) (result models.EmailDeliveryReportsResult, responseDetails models.ResponseDetails, err error)
	GetLogs(ctx context.Context, queryParams map[string]string) (result models.EmailLogsResult, responseDetails models.ResponseDetails, err error)
	GetSentBulks(ctx context.Context, queryParams map[string]string) (result models.SentEmailBulksResult, responseDetails models.ResponseDetails, err error)
	RescheduleMessages(ctx context.Context, request models.RescheduleMessagesRequest, queryParams map[string]string) (response models.RescheduleMessagesResponse, responseDetails models.ResponseDetails, err error)
	GetSentBulksStatus(ctx context.Context, queryParams map[string]string) (result models.SentEmailBulksStatusResult, responseDetails models.ResponseDetails, err error)
	UpdateScheduledMessagesStatus(ctx context.Context, request models.UpdateScheduledMessagesStatusRequest, queryParams map[string]string) (response models.UpdateScheduledMessagesStatusResponse, responseDetails models.ResponseDetails, err error)
	ValidateAddresses(ctx context.Context, request models.ValidateAddressesRequest) (result models.ValidateAddressesResult, responseDetails models.ResponseDetails, err error)
}

// Send sends an email message with all available features.
func (email *Channel) Send(
	ctx context.Context,
	msg models.EmailMsg,
) (msgResp models.SendEmailResponse, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendEmailPath)
	return msgResp, responseDetails, err
}

// GetDeliveryReports returns delivery reports for sent emails.
func (email *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams map[string]string,
) (response models.EmailDeliveryReportsResult, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.GetRequest(ctx, &response, getDeliveryReportsPath, queryParams)
	return response, responseDetails, err
}

func (email *Channel) GetLogs(ctx context.Context,
	queryParams map[string]string,
) (result models.EmailLogsResult, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.GetRequest(ctx, &result, getLogsPath, queryParams)
	return result, responseDetails, err
}

func (email *Channel) GetSentBulks(ctx context.Context,
	queryParams map[string]string,
) (result models.SentEmailBulksResult, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.GetRequest(ctx, &result, getSentEmailBulksPath, queryParams)
	return result, responseDetails, err
}

// RescheduleMessages changes the date and time for scheduled messages.
func (email *Channel) RescheduleMessages(
	ctx context.Context,
	request models.RescheduleMessagesRequest,
	queryParams map[string]string,
) (response models.RescheduleMessagesResponse, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.PutJSONReq(ctx, &request, &response, rescheduleMessagesPath, queryParams)
	return response, responseDetails, err
}

// GetSentBulksStatus returns status of scheduled email messages.
func (email *Channel) GetSentBulksStatus(ctx context.Context,
	queryParams map[string]string,
) (result models.SentEmailBulksStatusResult, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.GetRequest(ctx, &result, getSentEmailBulksStatusPath, queryParams)
	return result, responseDetails, err
}

// UpdateScheduledMessagesStatus updates status or completely cancels sending of scheduled messages.
func (email *Channel) UpdateScheduledMessagesStatus(ctx context.Context,
	request models.UpdateScheduledMessagesStatusRequest,
	queryParams map[string]string,
) (response models.UpdateScheduledMessagesStatusResponse, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.PutJSONReq(ctx, &request, &response, updateScheduledMessagesStatusPath, queryParams)
	return response, responseDetails, err
}

// ValidateAddresses validates to identify poor quality emails to clear up your recipient list.
func (email *Channel) ValidateAddresses(ctx context.Context,
	request models.ValidateAddressesRequest,
) (result models.ValidateAddressesResult, responseDetails models.ResponseDetails, err error) {
	responseDetails, err = email.ReqHandler.PostJSONReq(ctx, &request, &result, validateAddressesPath)
	return result, responseDetails, err
}
