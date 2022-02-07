package infobip

import (
	"context"
	"infobip-go-client/pkg/infobip/models"
	"net/http"
)

// WhatsApp provides methods to interact with the Infobip WhatsApp API.
// WhatsApp API docs: https://www.infobip.com/docs/api#channels/whatsapp
type WhatsApp interface {
	SendTextMessage(context.Context, models.TextMessageRequest) (models.TextMessageResponse, models.ResponseDetails, error)
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

func (wap *whatsAppChannel) SendTextMessage(ctx context.Context,
	message models.TextMessageRequest,
) (msgResp models.TextMessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &message, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}
