package email

import (
	"context"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	sendEmailPath = "/email/2/send"
	getDeliveryReportsPath = "/email/1/reports"
)

type Channel struct {
	ReqHandler internal.HTTPHandler
}

type Email interface {
	SendFullyFeatured(ctx context.Context, request models.EmailMsg) (response models.SendEmailResponse, responseDetails models.ResponseDetails, err error)
	GetDeliveryReports(ctx context.Context, queryParams map[string]string) (responseDetails models.ResponseDetails, err error)
}

// SendFullyFeatured sends an email message with all available features.
func (email *Channel) SendFullyFeatured(
	ctx context.Context,
	msg models.EmailMsg,
) (msgResp models.SendEmailResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.PostMultipartReq(ctx, &msg, &msgResp, sendEmailPath)
	return msgResp, respDetails, err
}

// GetDeliveryReports returns delivery reports for sent emails.
func (email *Channel) GetDeliveryReports(
	ctx context.Context,
	queryParams map[string]string,
) (response models.EmailDeliveryReportsResult, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.GetRequest(ctx, &response, sendEmailPath, queryParams)
	return response, respDetails, err
}
