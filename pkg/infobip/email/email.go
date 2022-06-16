package email

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v2/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
)

const (
	sendEmailPath                     = "email/2/send"
	getDeliveryReportsPath            = "email/1/reports"
	getLogsPath                       = "email/1/logs"
	getSentEmailBulksPath             = "email/1/bulks"
	rescheduleMessagesPath            = "email/1/bulks"
	getSentEmailBulksStatusPath       = "email/1/bulks/status"
	updateScheduledMessagesStatusPath = "email/1/bulks/status"
	validateAddressesPath             = "email/2/validation"
	getDomainsPath                    = "email/1/domains"
	addDomainPath                     = "email/1/domains"
	getDomainPath                     = "email/1/domains"
	deleteDomainPath                  = "email/1/domains"
	updateDomainTrackingPath          = "email/1/domains"
	verifyDomainPath                  = "email/1/domains"
)

type Channel struct {
	ReqHandler internal.HTTPHandler
}

type Email interface {
	// GetDeliveryReports gets one-time delivery reports for all sent emails.
	GetDeliveryReports(ctx context.Context, queryParams models.GetEmailDeliveryReportsParams) (
		resp models.GetEmailDeliveryReportsResponse, respDetails models.ResponseDetails, err error)

	// GetLogs gets email logs of sent Email messagesId for request. Logs are available for the last 48 hours.
	GetLogs(ctx context.Context, queryParams models.GetEmailLogsParams) (
		resp models.GetEmailLogsResponse, respDetails models.ResponseDetails, err error)

	// GetSentBulks gets the scheduled time of your Email messages.
	GetSentBulks(ctx context.Context, queryParams models.GetSentEmailBulksParams) (
		resp models.SentEmailBulksResponse, respDetails models.ResponseDetails, err error)

	// GetSentBulksStatus returns status of scheduled email messages.
	GetSentBulksStatus(ctx context.Context, queryParams models.GetSentEmailBulksStatusParams) (
		resp models.SentEmailBulksStatusResponse, respDetails models.ResponseDetails, err error)

	// RescheduleMessages changes the date and time for scheduled messages.
	RescheduleMessages(
		ctx context.Context, req models.RescheduleEmailRequest, queryParams models.RescheduleEmailParams) (
		resp models.RescheduleEmailResponse, respDetails models.ResponseDetails, err error)

	// Send sends an email or multiple emails to a recipient or multiple recipients with CC/BCC enabled.
	Send(ctx context.Context, req models.EmailMsg) (
		resp models.SendEmailResponse, respDetails models.ResponseDetails, err error)

	// UpdateScheduledMessagesStatus updates status or completely cancels sending of scheduled messages.
	UpdateScheduledMessagesStatus(
		ctx context.Context,
		req models.UpdateScheduledEmailStatusRequest,
		queryParams models.UpdateScheduledEmailStatusParams) (
		resp models.UpdateScheduledStatusResponse, respDetails models.ResponseDetails, err error)

	// ValidateAddresses validates to identify poor quality emails to clear up your recipient list.
	ValidateAddresses(ctx context.Context, req models.ValidateEmailAddressesRequest) (
		resp models.ValidateEmailAddressesResponse, respDetails models.ResponseDetails, err error)

	// GetDomains returns all domains associated with the account. It also provides details of the retrieved domain
	// like the DNS records, tracking details, active/blocked status, etc.
	GetDomains(ctx context.Context, queryParams models.GetEmailDomainsParams) (
		resp models.GetEmailDomainsResponse, respDetails models.ResponseDetails, err error)

	// AddDomain adds new domains with a limit to create a maximum of 10 domains in a day.
	AddDomain(ctx context.Context, req models.AddEmailDomainRequest) (
		resp models.AddEmailDomainResponse, respDetails models.ResponseDetails, err error)

	// GetDomain returns the details of the domain like the DNS records, tracking details, active/blocked status, etc.
	GetDomain(ctx context.Context, domainName string) (
		resp models.GetEmailDomainResponse, respDetails models.ResponseDetails, err error)

	// DeleteDomain deletes an existing domain.
	DeleteDomain(ctx context.Context, domainName string) (
		respDetails models.ResponseDetails, err error)

	// UpdateDomainTracking updates the tracking events for the provided domain. Tracking events can be updated
	// only for CLICKS, OPENS and UNSUBSCRIBES.
	UpdateDomainTracking(ctx context.Context, domainName string, req models.UpdateEmailDomainTrackingRequest) (
		resp models.UpdateEmailDomainTrackingResponse, respDetails models.ResponseDetails, err error)

	// VerifyDomain verifies records(TXT, MX, DKIM) associated with the provided domain.
	VerifyDomain(ctx context.Context, domainName string) (
		respDetails models.ResponseDetails, err error)
}

func (email *Channel) Send(
	ctx context.Context,
	msg models.EmailMsg,
) (msgResp models.SendEmailResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendEmailPath)
	return msgResp, respDetails, err
}

func (email *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams models.GetEmailDeliveryReportsParams,
) (resp models.GetEmailDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
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

func (email *Channel) GetLogs(
	ctx context.Context,
	queryParams models.GetEmailLogsParams,
) (resp models.GetEmailLogsResponse, respDetails models.ResponseDetails, err error) {
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

func (email *Channel) GetSentBulks(
	ctx context.Context,
	queryParams models.GetSentEmailBulksParams,
) (resp models.SentEmailBulksResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getSentEmailBulksPath, params)
	return resp, respDetails, err
}

func (email *Channel) RescheduleMessages(
	ctx context.Context,
	req models.RescheduleEmailRequest,
	queryParams models.RescheduleEmailParams,
) (resp models.RescheduleEmailResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.PutJSONReq(ctx, &req, &resp, rescheduleMessagesPath, params)
	return resp, respDetails, err
}

func (email *Channel) GetSentBulksStatus(
	ctx context.Context,
	queryParams models.GetSentEmailBulksStatusParams,
) (resp models.SentEmailBulksStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getSentEmailBulksStatusPath, params)
	return resp, respDetails, err
}

func (email *Channel) UpdateScheduledMessagesStatus(
	ctx context.Context,
	req models.UpdateScheduledEmailStatusRequest,
	queryParams models.UpdateScheduledEmailStatusParams,
) (resp models.UpdateScheduledStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{{Name: "bulkId", Value: queryParams.BulkID}}
	respDetails, err = email.ReqHandler.PutJSONReq(ctx, &req, &resp, updateScheduledMessagesStatusPath, params)
	return resp, respDetails, err
}

func (email *Channel) ValidateAddresses(
	ctx context.Context,
	req models.ValidateEmailAddressesRequest,
) (resp models.ValidateEmailAddressesResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostJSONReq(ctx, &req, &resp, validateAddressesPath)
	return resp, respDetails, err
}

func (email *Channel) GetDomains(
	ctx context.Context,
	queryParams models.GetEmailDomainsParams,
) (resp models.GetEmailDomainsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "size", Value: fmt.Sprint(queryParams.Size)},
		{Name: "page", Value: fmt.Sprint(queryParams.Page)},
	}
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, getDomainsPath, params)
	return resp, respDetails, err
}

func (email *Channel) AddDomain(
	ctx context.Context,
	req models.AddEmailDomainRequest,
) (resp models.AddEmailDomainResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostJSONReq(ctx, &req, &resp, addDomainPath)
	return resp, respDetails, err
}

func (email *Channel) GetDomain(
	ctx context.Context,
	domainName string,
) (resp models.GetEmailDomainResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.GetRequest(ctx, &resp, fmt.Sprint(getDomainPath, "/", domainName), nil)
	return resp, respDetails, err
}

func (email *Channel) DeleteDomain(
	ctx context.Context,
	domainName string,
) (respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.DeleteRequest(ctx, fmt.Sprint(deleteDomainPath, "/", domainName), nil)
	return respDetails, err
}

func (email *Channel) UpdateDomainTracking(
	ctx context.Context,
	domainName string,
	req models.UpdateEmailDomainTrackingRequest,
) (resp models.UpdateEmailDomainTrackingResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PutJSONReq(ctx, &req, &resp,
		fmt.Sprint(updateDomainTrackingPath, "/", domainName, "/tracking"), nil)
	return resp, respDetails, err
}

func (email *Channel) VerifyDomain(
	ctx context.Context,
	domainName string,
) (respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostNoBodyReq(ctx, nil,
		fmt.Sprint(verifyDomainPath, "/", domainName, "/verify"))
	return respDetails, err
}
