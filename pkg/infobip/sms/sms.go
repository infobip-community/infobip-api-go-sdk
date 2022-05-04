package sms

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	getDeliveryReportsPath = "/sms/1/reports"
	getLogsPath            = "/sms/1/logs"
	sendSmsPath            = "/sms/2/text/advanced"
)

type SMS interface {
	Send(ctx context.Context, req models.SendSMSRequest) (
		resp models.SendSMSResponse, respDetails models.ResponseDetails, err error)
	GetDeliveryReports(ctx context.Context, params models.GetSMSDeliveryReportsParams) (
		resp models.GetSMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error)
	GetLogs(ctx context.Context, params models.GetSMSLogsParams) (
		resp models.GetSMSLogsResponse, respDetails models.ResponseDetails, err error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

func (sms *Channel) Send(
	ctx context.Context,
	req models.SendSMSRequest,
) (resp models.SendSMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = sms.ReqHandler.PostJSONReq(ctx, &req, &resp, sendSmsPath)
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
