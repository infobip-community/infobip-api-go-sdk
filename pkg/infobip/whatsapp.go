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
	SendImageMessage(context.Context, models.ImageMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendAudioMessage(context.Context, models.AudioMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendVideoMessage(context.Context, models.VideoMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendStickerMessage(context.Context, models.StickerMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendLocationMessage(context.Context, models.LocationMessage) (models.MessageResponse, models.ResponseDetails, error)
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
const sendImagePath = "whatsapp/1/message/image"
const sendAudioPath = "whatsapp/1/message/audio"
const sendVideoPath = "whatsapp/1/message/video"
const sendStickerPath = "whatsapp/1/message/sticker"
const sendLocationPath = "whatsapp/1/message/location"

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

func (wap *whatsAppChannel) SendImageMessage(
	ctx context.Context,
	document models.ImageMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendImagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendAudioMessage(
	ctx context.Context,
	document models.AudioMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendAudioPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendVideoMessage(
	ctx context.Context,
	document models.VideoMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendVideoPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendStickerMessage(
	ctx context.Context,
	document models.StickerMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendStickerPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendLocationMessage(
	ctx context.Context,
	document models.LocationMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &document, &msgResp, sendLocationPath)
	return msgResp, respDetails, err
}
