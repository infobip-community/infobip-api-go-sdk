package sms

import (
	"context"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	sendSmsPath = "/sms/2/text/advanced"
)

type SMS interface {
	Send(ctx context.Context, req models.SendSMSRequest) (
		resp models.SendSMSResponse, respDetails models.ResponseDetails, err error)
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
