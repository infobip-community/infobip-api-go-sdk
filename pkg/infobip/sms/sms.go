package sms

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
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
)

// TODO:
// [x] Implement validate for all models that have validate strings.
// [ ] Check that models have all required fields (especially responses)
// [x] Check that naming is consistent for models, and short for methods.
// [ ] All tests pass
// [ ] Add tests
// [ ] Add godoc comments
// [ ] Lint passes
// [ ] No credentials in code.
// [ ] Add WebRTC and RCS endpoints.
// [ ] There aro no TODO notes.

type SMS interface {
	Send(ctx context.Context, req models.SendSMSRequest) (
		resp models.SendSMSResponse, respDetails models.ResponseDetails, err error)
	SendBinary(ctx context.Context, req models.SendBinarySMSRequest) (
		resp models.SendBinarySMSResponse, respDetails models.ResponseDetails, err error)
	SendOverQueryParams(ctx context.Context, queryParams models.SendSMSOverQueryParamsParams) (
		resp models.SendSMSOverQueryParamsResponse, respDetails models.ResponseDetails, err error)
	Preview(ctx context.Context, req models.PreviewSMSRequest) (
		resp models.PreviewSMSResponse, respDetails models.ResponseDetails, err error)
	GetDeliveryReports(ctx context.Context, queryParams models.GetSMSDeliveryReportsParams) (
		resp models.GetSMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error)
	GetLogs(ctx context.Context, queryParams models.GetSMSLogsParams) (
		resp models.GetSMSLogsResponse, respDetails models.ResponseDetails, err error)
	GetInboundMessages(ctx context.Context, queryParams models.GetInboundSMSParams) (
		resp models.GetInboundSMSResponse, respDetails models.ResponseDetails, err error)
	GetScheduledMessages(ctx context.Context, queryParams models.GetScheduledSMSParams) (
		resp models.GetScheduledSMSResponse, respDetails models.ResponseDetails, err error)
	RescheduleMessages(
		ctx context.Context, req models.RescheduleSMSRequest, queryParams models.RescheduleSMSParams) (
		resp models.RescheduleSMSResponse, respDetails models.ResponseDetails, err error)
	GetScheduledMessagesStatus(ctx context.Context, queryParams models.GetScheduledSMSStatusParams) (
		resp models.GetScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error)
	UpdateScheduledMessagesStatus(
		ctx context.Context,
		req models.UpdateScheduledSMSStatusRequest,
		queryParams models.UpdateScheduledSMSStatusParams,
	) (resp models.UpdateScheduledSMSStatusResponse, respDetails models.ResponseDetails, err error)
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
