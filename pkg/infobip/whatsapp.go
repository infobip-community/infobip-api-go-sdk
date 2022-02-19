package infobip

import (
	"context"
	"fmt"
	"infobip-go-client/pkg/infobip/models"
	"net/http"
)

// WhatsApp provides methods to interact with the Infobip WhatsApp API.
// WhatsApp API docs: https://www.infobip.com/docs/api#channels/whatsapp
type WhatsApp interface {
	SendTemplateMessages(context.Context, models.TemplateMessages,
	) (models.BulkMessageResponse, models.ResponseDetails, error)
	SendTextMessage(context.Context, models.TextMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendDocumentMessage(context.Context, models.DocumentMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendImageMessage(context.Context, models.ImageMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendAudioMessage(context.Context, models.AudioMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendVideoMessage(context.Context, models.VideoMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendStickerMessage(context.Context, models.StickerMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendLocationMessage(context.Context, models.LocationMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendContactMessage(context.Context, models.ContactMessage) (models.MessageResponse, models.ResponseDetails, error)
	SendInteractiveButtonsMessage(context.Context, models.InteractiveButtonsMessage,
	) (models.MessageResponse, models.ResponseDetails, error)
	SendInteractiveListMessage(context.Context, models.InteractiveListMessage,
	) (models.MessageResponse, models.ResponseDetails, error)
	SendInteractiveProductMessage(context.Context, models.InteractiveProductMessage,
	) (models.MessageResponse, models.ResponseDetails, error)
	SendInteractiveMultiproductMessage(context.Context, models.InteractiveMultiproductMessage,
	) (models.MessageResponse, models.ResponseDetails, error)
	GetTemplates(context.Context, string) (models.TemplatesResponse, models.ResponseDetails, error)
	CreateTemplate(context.Context, string, models.TemplateCreate) (models.TemplateResponse, models.ResponseDetails, error)
}

type whatsAppChannel struct {
	reqHandler httpHandler
}

func newWhatsApp(apiKey string, baseURL string, httpClient http.Client) *whatsAppChannel {
	return &whatsAppChannel{
		reqHandler: httpHandler{apiKey: apiKey, baseURL: baseURL, httpClient: httpClient},
	}
}

const sendTemplateMessagesPath = "whatsapp/1/message/template"
const sendMessagePath = "whatsapp/1/message/text"
const sendDocumentPath = "whatsapp/1/message/document"
const sendImagePath = "whatsapp/1/message/image"
const sendAudioPath = "whatsapp/1/message/audio"
const sendVideoPath = "whatsapp/1/message/video"
const sendStickerPath = "whatsapp/1/message/sticker"
const sendLocationPath = "whatsapp/1/message/location"
const sendContactPath = "whatsapp/1/message/contact"
const sendInteractiveButtonsPath = "whatsapp/1/message/interactive/buttons"
const sendInteractiveListPath = "whatsapp/1/message/interactive/list"
const sendInteractiveProductPath = "whatsapp/1/message/interactive/product"
const sendInteractiveMultiproductPath = "whatsapp/1/message/interactive/multi-product"
const templatesPath = "whatsapp/1/senders/%s/templates"

func (wap *whatsAppChannel) SendTemplateMessages(
	ctx context.Context,
	messages models.TemplateMessages,
) (msgResp models.BulkMessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &messages, &msgResp, sendTemplateMessagesPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendTextMessage(
	ctx context.Context,
	msg models.TextMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendDocumentMessage(
	ctx context.Context,
	msg models.DocumentMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendDocumentPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendImageMessage(
	ctx context.Context,
	msg models.ImageMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendImagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendAudioMessage(
	ctx context.Context,
	msg models.AudioMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendAudioPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendVideoMessage(
	ctx context.Context,
	msg models.VideoMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendVideoPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendStickerMessage(
	ctx context.Context,
	msg models.StickerMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendStickerPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendLocationMessage(
	ctx context.Context,
	msg models.LocationMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendLocationPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendContactMessage(
	ctx context.Context,
	msg models.ContactMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendContactPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveButtonsMessage(
	ctx context.Context,
	msg models.InteractiveButtonsMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveButtonsPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveListMessage(
	ctx context.Context,
	msg models.InteractiveListMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveListPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveProductMessage(
	ctx context.Context,
	msg models.InteractiveProductMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveProductPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveMultiproductMessage(
	ctx context.Context,
	msg models.InteractiveMultiproductMessage,
) (msgResp models.MessageResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveMultiproductPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) GetTemplates(
	ctx context.Context,
	sender string,
) (resp models.TemplatesResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.getRequest(ctx, &resp, fmt.Sprintf(templatesPath, sender))
	return resp, respDetails, err
}

func (wap *whatsAppChannel) CreateTemplate(
	ctx context.Context,
	sender string,
	template models.TemplateCreate,
) (resp models.TemplateResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &template, &resp, fmt.Sprintf(templatesPath, sender))
	return resp, respDetails, err
}
