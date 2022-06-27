package mms

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

// MMS provides methods to interact with the Infobip MMS API.
// MMS API docs: https://www.infobip.com/docs/api#channels/mms
type MMS interface {
	Send(context.Context, models.MMSMsg) (models.SendMMSResponse, models.ResponseDetails, error)
	GetDeliveryReports(ctx context.Context, queryParams models.GetMMSDeliveryReportsParams) (
		models.GetMMSDeliveryReportsResponse, models.ResponseDetails, error)
	GetInboundMessages(ctx context.Context, queryParams models.GetInboundMMSParams) (
		models.GetInboundMMSResponse, models.ResponseDetails, error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

const (
	sendMessagePath                   = "mms/1/single"
	getOutboundMMSDeliveryReportsPath = "mms/1/reports"
	getInboundMMSPath                 = "mms/1/inbox/reports"
)

func (mms *Channel) Send(
	ctx context.Context,
	msg models.MMSMsg,
) (msgResp models.SendMMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = mms.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (mms *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams models.GetMMSDeliveryReportsParams,
) (msgResp models.GetMMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "bulkId", Value: queryParams.BulkID},
		{Name: "messageId", Value: queryParams.MessageID},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}

	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getOutboundMMSDeliveryReportsPath, params)
	return msgResp, respDetails, err
}

func (mms *Channel) GetInboundMessages(
	ctx context.Context,
	queryParams models.GetInboundMMSParams,
) (msgResp models.GetInboundMMSResponse, respDetails models.ResponseDetails, err error) {
	var params []internal.QueryParameter
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}
	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getInboundMMSPath, params)
	return msgResp, respDetails, err
}
