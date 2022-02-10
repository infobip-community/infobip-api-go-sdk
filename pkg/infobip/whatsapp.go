package infobip

import (
	"context"
	"infobip-go-client/pkg/infobip/models"
	"net/http"
)

// WhatsApp provides methods to interact with the Infobip WhatsApp API.
// WhatsApp API docs: https://www.infobip.com/docs/api#channels/whatsapp
type WhatsApp interface {
	SendTextMessage(context.Context, models.TextMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendDocumentMessage(context.Context, models.DocumentMessage) (models.MessageResponse, models.ResponseDetails, error)
}

type whatsAppChannel struct {
	reqHandler httpHandler
}

func newWhatsApp(apiKey string, baseURL string, httpClient http.Client) *whatsAppChannel {
	return &whatsAppChannel{
		reqHandler: httpHandler{apiKey: apiKey, baseURL: baseURL, httpClient: httpClient},
	}
}

const sendMessagePath = "whatsapp/1/message/text"
const sendDocumentPath = "whatsapp/1/message/document"

func (wap *whatsAppChannel) SendTextMessage(
	ctx context.Context,
	text models.TextMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &text, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendDocumentMessage(
	ctx context.Context,
	document models.DocumentMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendDocumentPath)
	return msgResp, respDetails, err
}
