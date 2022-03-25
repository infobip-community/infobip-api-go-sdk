package email

import (
	"context"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	_ = iota
	SendEmailPath = "/email/2/send"
)

type Channel struct {
	ReqHandler internal.HTTPHandler
}

type Email interface {
	Send(ctx context.Context, request models.EmailMsg) (response models.EmailResponse, responseDetails models.ResponseDetails, err error)
}

func (email *Channel) Send(
	ctx context.Context,
	msg models.EmailMsg,
) (msgResp models.EmailResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostRequest(ctx, &msg, &msgResp, SendEmailPath)
	return msgResp, respDetails, err
}
