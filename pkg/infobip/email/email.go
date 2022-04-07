package email

import (
	"context"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	sendEmailPath          = "/email/2/send"
	getDeliveryReportsPath = "/email/1/reports"
	getLogsPath            = "/email/1/logs"
)

type Channel struct {
	ReqHandler internal.HTTPHandler
}

type Email interface {
	Send(ctx context.Context, request models.EmailMsg) (response models.SendEmailResponse, responseDetails models.ResponseDetails, err error)
	GetDeliveryReports(ctx context.Context, queryParams map[string]string) (result models.EmailDeliveryReportsResult, respDetails models.ResponseDetails, err error)
	GetLogs(ctx context.Context, queryParams map[string]string) (result models.EmailLogsResult, respDetails models.ResponseDetails, err error)
}

// Send sends an email message with all available features.
func (email *Channel) Send(
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
	respDetails, err = email.ReqHandler.GetRequest(ctx, &response, getDeliveryReportsPath, queryParams)
	return response, respDetails, err
}

func (email *Channel) GetLogs(ctx context.Context, queryParams map[string]string) (result models.EmailLogsResult, respDetails models.ResponseDetails, err error) {
	respDetails, err = email.ReqHandler.GetRequest(ctx, &result, getLogsPath, queryParams)
	return result, respDetails, err
}
