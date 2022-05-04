package mms

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

// MMS provides methods to interact with the Infobip MMS API.
// MMS API docs: https://www.infobip.com/docs/api#channels/mms
type MMS interface {
	SendMsg(context.Context, models.MMSMsg) (models.MMSResponse, models.ResponseDetails, error)
	GetOutboundMsgDeliveryReports(ctx context.Context, params models.OutboundMMSDeliveryReportsParams) (
		models.OutboundMMSDeliveryReportsResponse, models.ResponseDetails, error)
	GetInboundMsgs(ctx context.Context, params models.InboundMMSParams) (
		models.InboundMMSResponse, models.ResponseDetails, error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

const sendMessagePath = "/mms/1/single"
const getOutboundMMSDeliveryReportsPath = "/mms/1/reports"
const getInboundMMSPath = "/mms/1/inbox/reports"

func (mms *Channel) SendMsg(
	ctx context.Context,
	msg models.MMSMsg,
) (msgResp models.MMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = mms.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (mms *Channel) GetOutboundMsgDeliveryReports(
	ctx context.Context,
	params models.OutboundMMSDeliveryReportsParams,
) (msgResp models.OutboundMMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
	queryParams := []internal.QueryParameter{
		{Name: "bulkId", Value: params.BulkID},
		{Name: "messageId", Value: params.MessageID},
		{Name: "limit", Value: fmt.Sprint(params.Limit)},
	}
	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getOutboundMMSDeliveryReportsPath, queryParams)
	return msgResp, respDetails, err
}

func (mms *Channel) GetInboundMsgs(
	ctx context.Context,
	params models.InboundMMSParams,
) (msgResp models.InboundMMSResponse, respDetails models.ResponseDetails, err error) {
	queryParams := []internal.QueryParameter{{Name: "limit", Value: fmt.Sprint(params.Limit)}}
	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getInboundMMSPath, queryParams)
	return msgResp, respDetails, err
}
