package rcs

import (
	"context"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

const (
	sendRCSPath     = "ott/rcs/1/message"
	sendRCSBulkPath = "ott/rcs/1/message/bulk"
)

type RCS interface {
	// Send sends a single RCS message.
	Send(
		ctx context.Context,
		msg models.RCSMsg,
	) (resp models.SendRCSResponse, respDetails models.ResponseDetails, err error)

	// SendBulk sends bulk RCS messages.
	SendBulk(
		ctx context.Context,
		req models.SendRCSBulkRequest,
	) (resp models.SendRCSBulkResponse, respDetails models.ResponseDetails, err error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

func (rcs *Channel) Send(
	ctx context.Context,
	msg models.RCSMsg,
) (resp models.SendRCSResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = rcs.ReqHandler.PostJSONReq(ctx, &msg, &resp, sendRCSPath)
	return resp, respDetails, err
}

func (rcs *Channel) SendBulk(
	ctx context.Context,
	req models.SendRCSBulkRequest,
) (resp models.SendRCSBulkResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = rcs.ReqHandler.PostJSONReq(ctx, &req, &resp, sendRCSBulkPath)
	return resp, respDetails, err
}
