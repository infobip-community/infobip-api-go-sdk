package sms

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

const (
	getDeliveryReportsPath       = "sms/1/reports"
	getLogsPath                  = "sms/1/logs"
	sendSMSPath                  = "sms/2/text/advanced"
	sendBinarySMSPath            = "sms/2/binary/advanced"
	sendSMSOverQueryParamsPath   = "sms/1/text/query"
	previewSMSPath               = "sms/1/preview"
	getInboundSMSPath            = "sms/1/inbox/reports"
	getScheduledSMSPath          = "sms/1/bulks"
	rescheduleSMSPath            = "sms/1/bulks"
	getScheduledSMSStatusPath    = "sms/1/bulks/status"
	updateScheduledSMSStatusPath = "sms/1/bulks/status"
	getTFAApplicationsPath       = "2fa/2/applications"
	createTFAApplicationPath     = "2fa/2/applications"
	getTFAApplicationPath        = "/2fa/2/applications"
	updateTFAApplicationPath     = "/2fa/2/applications"
)

type SMS interface {
	// Send sends everything from a simple single message to a single destination, up to batch sending of personalized
	// messages to the thousands of recipients with a single API request.
	Send(ctx context.Context, req models.SendSMSRequest) (
		resp models.SendSMSResponse, respDetails models.ResponseDetails, err error)

	// SendBinary sends single or multiple binary messages to one or more destination addresses.
	SendBinary(ctx context.Context, req models.SendBinarySMSRequest) (
		resp models.SendBinarySMSResponse, respDetails models.ResponseDetails, err error)

	// SendOverQueryParams sends messages over query parameters. All message parameters of the message can be defined
	// in the query string. Use this method only if Send message is not an option for your use case!
	// Note: Make sure that special characters and user credentials are properly encoded. Use a URL encoding reference
	// as a guide.
	SendOverQueryParams(ctx context.Context, queryParams models.SendSMSOverQueryParamsParams) (
		resp models.SendSMSOverQueryParamsResponse, respDetails models.ResponseDetails, err error)

	// Preview returns information on how different message configurations will affect your message text, number of
	// characters and message parts.
	Preview(ctx context.Context, req models.PreviewSMSRequest) (
		resp models.PreviewSMSResponse, respDetails models.ResponseDetails, err error)

	// GetDeliveryReports returns information about if and when the message has been delivered to the recipient. Each
	// request will return a batch of delivery reports only once. The request will return only new reports that arrived
	// since the last request in the last 48 hours.
	GetDeliveryReports(ctx context.Context, queryParams models.GetSMSDeliveryReportsParams) (
		resp models.GetSMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error)

	// GetLogs returns logs for the last 48 hours, and you can only retrieve maximum of 1000 logs per call. See
	// GetDeliveryReports if your use case is to verify message delivery.
	GetLogs(ctx context.Context, queryParams models.GetSMSLogsParams) (
		resp models.GetSMSLogsResponse, respDetails models.ResponseDetails, err error)

	// GetInboundMessages returns inbound messages. If for some reason you are unable to receive incoming SMS to the
	// endpoint of your choice in real time, you can use this call to fetch messages. Each request will return a batch of
	// received messages only once. The API request will only return new messages that arrived since the last request.
	GetInboundMessages(ctx context.Context, queryParams models.GetInboundSMSParams) (
		resp models.GetInboundSMSResponse, respDetails models.ResponseDetails, err error)

	// GetScheduledMessages returns the status and the scheduled time of your SMS messages.
	GetScheduledMessages(ctx context.Context, queryParams models.GetScheduledSMSParams) (
		resp models.GetScheduledSMSResponse, respDetails models.ResponseDetails, err error)

	// RescheduleMessages changes the date and time for sending scheduled messages.
	RescheduleMessages(
		ctx context.Context, req models.RescheduleSMSRequest, queryParams models.RescheduleSMSParams) (
		resp models.RescheduleSMSResponse, respDetails models.ResponseDetails, err error)

	// GetScheduledMessagesStatus returns the status of scheduled messages.
	GetScheduledMessagesStatus(ctx context.Context, queryParams models.GetScheduledSMSStatusParams) (
		resp models.GetScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error)

	// UpdateScheduledMessagesStatus changes status or completely cancels sending of scheduled messages.
	UpdateScheduledMessagesStatus(
		ctx context.Context,
		req models.UpdateScheduledSMSStatusRequest,
		queryParams models.UpdateScheduledSMSStatusParams,
	) (resp models.UpdateScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error)

	// GetTFAApplications returns your applications list.
	GetTFAApplications(
		ctx context.Context,
	) (resp models.GetTFAApplicationsResponse, respDetails models.ResponseDetails, err error)

	// CreateTFAApplication create and configure a new 2FA application.
	CreateTFAApplication(
		ctx context.Context,
		req models.CreateTFAApplicationRequest,
	) (resp models.CreateTFAApplicationResponse, respDetails models.ResponseDetails, err error)

	// GetTFAApplication returns a single 2FA message template from an application to see its configuration details.
	GetTFAApplication(
		ctx context.Context,
		appID string,
	) (resp models.GetTFAApplicationResponse, respDetails models.ResponseDetails, err error)

	// UpdateTFAApplication changes configuration options for your existing 2FA application.
	UpdateTFAApplication(
		ctx context.Context,
		req models.UpdateTFAApplicationRequest,
		appID string,
	) (resp models.UpdateTFAApplicationResponse, respDetails models.ResponseDetails, err error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

func (sms *Channel) Send(
	ctx context.Context,
	req models.SendSMSRequest,
) (resp models.SendSMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PostJSONReq(ctx, &req, &resp, sendSMSPath)
	return resp, respDetails, err
}

func (sms *Channel) SendBinary(
	ctx context.Context,
	req models.SendBinarySMSRequest,
) (resp models.SendBinarySMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PostJSONReq(ctx, &req, &resp, sendBinarySMSPath)
	return resp, respDetails, err
}

func (sms *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams models.GetSMSDeliveryReportsParams) (
	resp models.GetSMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
		{Name: "messageId", Value: queryParams.MessageID},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}
	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getDeliveryReportsPath, params)
	return resp, respDetails, err
}

func (sms *Channel) GetLogs(
	ctx context.Context,
	queryParams models.GetSMSLogsParams,
) (resp models.GetSMSLogsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "from", Value: queryParams.From},
		{Name: "to", Value: queryParams.To},
		{Name: "generalStatus", Value: queryParams.GeneralStatus},
		{Name: "sentSince", Value: queryParams.SentSince},
		{Name: "sentUntil", Value: queryParams.SentUntil},
		{Name: "mcc", Value: queryParams.MCC},
		{Name: "mnc", Value: queryParams.MNC},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}

	for _, id := range queryParams.BulkID {
		params = append(params, internal.QueryParameter{Name: "bulkId", Value: id})
	}

	for _, id := range queryParams.MessageID {
		params = append(params, internal.QueryParameter{Name: "messageId", Value: id})
	}

	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getLogsPath, params)
	return resp, respDetails, err
}

func (sms *Channel) SendOverQueryParams(
	ctx context.Context,
	queryParams models.SendSMSOverQueryParamsParams,
) (resp models.SendSMSOverQueryParamsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "username", Value: queryParams.Username},
		{Name: "password", Value: queryParams.Password},
		{Name: "bulkId", Value: queryParams.BulkID},
		{Name: "from", Value: queryParams.From},
		{Name: "text", Value: queryParams.Text},
		{Name: "flash", Value: fmt.Sprint(queryParams.Flash)},
		{Name: "transliteration", Value: queryParams.Transliteration},
		{Name: "languageCode", Value: queryParams.LanguageCode},
		{Name: "intermediateReport", Value: fmt.Sprint(queryParams.IntermediateReport)},
		{Name: "notifyUrl", Value: queryParams.NotifyURL},
		{Name: "notifyContentType", Value: queryParams.NotifyContentType},
		{Name: "callbackData", Value: queryParams.CallbackData},
		{Name: "validityPeriod", Value: fmt.Sprint(queryParams.ValidityPeriod)},
		{Name: "sendAt", Value: queryParams.SendAt},
		{Name: "track", Value: queryParams.Track},
		{Name: "processKey", Value: queryParams.ProcessKey},
		{Name: "trackingType", Value: queryParams.TrackingType},
		{Name: "indiaDltContentTemplateId", Value: queryParams.IndiaDLTContentTemplateID},
		{Name: "indiaDltPrincipalEntityId", Value: queryParams.IndiaDLTPrincipalEntityID},
	}
	for _, to := range queryParams.To {
		params = append(params, internal.QueryParameter{Name: "to", Value: to})
	}

	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, sendSMSOverQueryParamsPath, params)

	return resp, respDetails, err
}

func (sms *Channel) Preview(
	ctx context.Context,
	req models.PreviewSMSRequest,
) (resp models.PreviewSMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PostJSONReq(ctx, &req, &resp, previewSMSPath)

	return resp, respDetails, err
}

func (sms *Channel) GetInboundMessages(
	ctx context.Context,
	queryParams models.GetInboundSMSParams,
) (resp models.GetInboundSMSResponse, respDetails models.ResponseDetails, err error) {
	var params []internal.QueryParameter
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}

	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getInboundSMSPath, params)

	return resp, respDetails, err
}

func (sms *Channel) GetScheduledMessages(
	ctx context.Context,
	queryParams models.GetScheduledSMSParams,
) (resp models.GetScheduledSMSResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
	}

	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getScheduledSMSPath, params)

	return resp, respDetails, err
}

func (sms *Channel) RescheduleMessages(
	ctx context.Context,
	req models.RescheduleSMSRequest,
	queryParams models.RescheduleSMSParams,
) (resp models.RescheduleSMSResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
	}

	respDetails, err = sms.ReqHandler.PutJSONReq(ctx, &req, &resp, rescheduleSMSPath, params)

	return resp, respDetails, err
}

func (sms *Channel) GetScheduledMessagesStatus(
	ctx context.Context,
	queryParams models.GetScheduledSMSStatusParams,
) (resp models.GetScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
	}

	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getScheduledSMSStatusPath, params)

	return resp, respDetails, err
}

func (sms *Channel) UpdateScheduledMessagesStatus(
	ctx context.Context,
	req models.UpdateScheduledSMSStatusRequest,
	queryParams models.UpdateScheduledSMSStatusParams,
) (resp models.UpdateScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
	}

	respDetails, err = sms.ReqHandler.PutJSONReq(ctx, &req, &resp, updateScheduledSMSStatusPath, params)

	return resp, respDetails, err
}

func (sms *Channel) GetTFAApplications(
	ctx context.Context,
) (resp models.GetTFAApplicationsResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, getTFAApplicationsPath, nil)

	return resp, respDetails, err
}

func (sms *Channel) CreateTFAApplication(
	ctx context.Context,
	req models.CreateTFAApplicationRequest,
) (resp models.CreateTFAApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PostJSONReq(ctx, &req, &resp, createTFAApplicationPath)

	return resp, respDetails, err
}

func (sms *Channel) GetTFAApplication(
	ctx context.Context,
	appId string,
) (resp models.GetTFAApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.GetRequest(ctx, &resp, fmt.Sprint(getTFAApplicationPath, "/", appId), nil)

	return resp, respDetails, err
}

func (sms *Channel) UpdateTFAApplication(
	ctx context.Context,
	req models.UpdateTFAApplicationRequest,
	appId string,
) (resp models.UpdateTFAApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PutJSONReq(ctx, &req, &resp, fmt.Sprint(updateTFAApplicationPath, "/", appId), nil)

	return resp, respDetails, err
}
