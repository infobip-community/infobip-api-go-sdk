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
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

const sendMessagePath = "/mms/1/single"

func (mms *Channel) SendMsg(
	ctx context.Context,
	msg models.MMSMsg,
) (msgResp models.MMSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = mms.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}
