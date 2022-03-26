package mms

import (
	"context"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

// MMS provides methods to interact with the Infobip MMS API.
// MMS API docs: https://www.infobip.com/docs/api#channels/mms
type MMS interface {
	SendMsg(context.Context, models.MMSMsg) (models.MMSResponse, models.ResponseDetails, error)
	GetOutboundMsgDeliveryReports(ctx context.Context, opts models.OutboundMMSDeliveryReportsOpts) (
		models.OutboundMMSDeliveryReportsResponse, models.ResponseDetails, error)
	GetInboundMsgs(ctx context.Context, opts models.InboundMMSOpts) (
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
	opts models.OutboundMMSDeliveryReportsOpts,
) (msgResp models.OutboundMMSDeliveryReportsResponse, respDetails models.ResponseDetails, err error) {
	params := map[string]string{"bulkId": opts.BulkID, "messageId": opts.MessageID, "limit": opts.Limit}
	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getOutboundMMSDeliveryReportsPath, params)
	return msgResp, respDetails, err
}

func (mms *Channel) GetInboundMsgs(
	ctx context.Context,
	opts models.InboundMMSOpts,
) (msgResp models.InboundMMSResponse, respDetails models.ResponseDetails, err error) {
	params := map[string]string{"limit": opts.Limit}
	respDetails, err = mms.ReqHandler.GetRequest(ctx, &msgResp, getInboundMMSPath, params)
	return msgResp, respDetails, err
}
