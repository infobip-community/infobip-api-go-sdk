package email

import (
	"context"
	"fmt"

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
	GetDeliveryReports(ctx context.Context, queryParams models.GetEmailDeliveryReportsParams) (
		resp models.EmailDeliveryReportsResponse, respDetails models.ResponseDetails, err error)
	GetLogs(ctx context.Context, queryParams models.GetEmailLogsParams) (
		resp models.EmailLogsResponse, respDetails models.ResponseDetails, err error)
	GetSentBulks(ctx context.Context, queryParams models.GetSentEmailBulksParams) (
		resp models.SentEmailBulksResponse, respDetails models.ResponseDetails, err error)
	GetSentBulksStatus(ctx context.Context, queryParams models.GetSentEmailBulksStatusParams) (
		resp models.SentEmailBulksStatusResponse, respDetails models.ResponseDetails, err error)
	RescheduleMessages(
		ctx context.Context, req models.RescheduleEmailMessagesRequest, queryParams models.RescheduleEmailMessagesParams) (
		resp models.RescheduleMessagesResponse, respDetails models.ResponseDetails, err error)
	Send(ctx context.Context, req models.EmailMsg) (
		resp models.SendEmailResponse, respDetails models.ResponseDetails, err error)
	UpdateScheduledMessagesStatus(
		ctx context.Context,
		req models.UpdateScheduledEmailMessagesStatusRequest,
		queryParams models.UpdateScheduledEmailMessagesStatusParams) (
		resp models.UpdateScheduledMessagesStatusResponse, respDetails models.ResponseDetails, err error)
	ValidateAddresses(ctx context.Context, req models.ValidateEmailAddressesRequest) (
		resp models.ValidateEmailAddressesResponse, respDetails models.ResponseDetails, err error)
}

// Send sends an email or multiple emails to a recipient or multiple recipients with CC/BCC enabled.
func (email *Channel) Send(
	ctx context.Context,
	msg models.EmailMsg,
) (msgResp models.SendEmailResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendEmailPath)
	return msgResp, respDetails, err
}

// GetDeliveryReports gets one-time delivery reports for all sent emails.
func (email *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams models.GetEmailDeliveryReportsParams,
) (resp models.EmailDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
		{Name: "messageId", Value: queryParams.MessageID},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}

	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getDeliveryReportsPath, params)
	return resp, respDetails, err
}

// GetLogs gets email logs of sent Email messagesId for request. Logs are available for the last 48 hours.
func (email *Channel) GetLogs(
	ctx context.Context,
	queryParams models.GetEmailLogsParams,
) (resp models.EmailLogsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "messageId", Value: queryParams.MessageID},
		{Name: "from", Value: queryParams.From},
		{Name: "to", Value: queryParams.To},
		{Name: "bulkId", Value: queryParams.BulkID},
		{Name: "generalStatus", Value: queryParams.GeneralStatus},
		{Name: "sentSince", Value: queryParams.SentSince},
		{Name: "sentUntil", Value: queryParams.SentUntil},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}

	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getLogsPath, params)
	return resp, respDetails, err
}

// GetSentBulks gets the scheduled time of your Email messages.
func (email *Channel) GetSentBulks(ctx context.Context,
	queryParams models.GetSentEmailBulksParams,
) (resp models.SentEmailBulksResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getSentEmailBulksPath, params)
	return resp, respDetails, err
}

// RescheduleMessages changes the date and time for scheduled messages.
func (email *Channel) RescheduleMessages(
	ctx context.Context,
	req models.RescheduleEmailMessagesRequest,
	queryParams models.RescheduleEmailMessagesParams,
) (resp models.RescheduleMessagesResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.PutJSONReq(ctx, &req, &resp, rescheduleMessagesPath, params)
	return resp, respDetails, err
}

// GetSentBulksStatus returns status of scheduled email messages.
func (email *Channel) GetSentBulksStatus(ctx context.Context,
	queryParams models.GetSentEmailBulksStatusParams,
) (resp models.SentEmailBulksStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getSentEmailBulksStatusPath, params)
	return resp, respDetails, err
}

// UpdateScheduledMessagesStatus updates status or completely cancels sending of scheduled messages.
func (email *Channel) UpdateScheduledMessagesStatus(ctx context.Context,
	req models.UpdateScheduledEmailMessagesStatusRequest,
	queryParams models.UpdateScheduledEmailMessagesStatusParams,
) (resp models.UpdateScheduledMessagesStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.PutJSONReq(ctx, &req, &resp, updateScheduledMessagesStatusPath, params)
	return resp, respDetails, err
}

// ValidateAddresses validates to identify poor quality emails to clear up your recipient list.
func (email *Channel) ValidateAddresses(ctx context.Context,
	req models.ValidateEmailAddressesRequest,
) (resp models.ValidateEmailAddressesResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostJSONReq(ctx, &req, &resp, validateAddressesPath)
	return resp, respDetails, err
}
