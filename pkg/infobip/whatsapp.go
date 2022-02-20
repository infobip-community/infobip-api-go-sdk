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
	SendTemplateMsgs(context.Context, models.TemplateMsgs) (models.BulkMsgResponse, models.ResponseDetails, error)
	SendTextMsg(context.Context, models.TextMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendDocumentMsg(context.Context, models.DocumentMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendImageMsg(context.Context, models.ImageMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendAudioMsg(context.Context, models.AudioMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendVideoMsg(context.Context, models.VideoMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendStickerMsg(context.Context, models.StickerMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendLocationMsg(context.Context, models.LocationMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendContactMsg(context.Context, models.ContactMsg) (models.MsgResponse, models.ResponseDetails, error)
	SendInteractiveButtonsMsg(context.Context, models.InteractiveButtonsMsg,
	) (models.MsgResponse, models.ResponseDetails, error)
	SendInteractiveListMsg(context.Context, models.InteractiveListMsg,
	) (models.MsgResponse, models.ResponseDetails, error)
	SendInteractiveProductMsg(context.Context, models.InteractiveProductMsg,
	) (models.MsgResponse, models.ResponseDetails, error)
	SendInteractiveMultiproductMsg(context.Context, models.InteractiveMultiproductMsg,
	) (models.MsgResponse, models.ResponseDetails, error)
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

func (wap *whatsAppChannel) SendTemplateMsgs(
	ctx context.Context,
	messages models.TemplateMsgs,
) (msgResp models.BulkMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &messages, &msgResp, sendTemplateMessagesPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendTextMsg(
	ctx context.Context,
	msg models.TextMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendDocumentMsg(
	ctx context.Context,
	msg models.DocumentMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendDocumentPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendImageMsg(
	ctx context.Context,
	msg models.ImageMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendImagePath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendAudioMsg(
	ctx context.Context,
	msg models.AudioMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendAudioPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendVideoMsg(
	ctx context.Context,
	msg models.VideoMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendVideoPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendStickerMsg(
	ctx context.Context,
	msg models.StickerMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendStickerPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendLocationMsg(
	ctx context.Context,
	msg models.LocationMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendLocationPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendContactMsg(
	ctx context.Context,
	msg models.ContactMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendContactPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveButtonsMsg(
	ctx context.Context,
	msg models.InteractiveButtonsMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveButtonsPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveListMsg(
	ctx context.Context,
	msg models.InteractiveListMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveListPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveProductMsg(
	ctx context.Context,
	msg models.InteractiveProductMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.reqHandler.postRequest(ctx, &msg, &msgResp, sendInteractiveProductPath)
	return msgResp, respDetails, err
}

func (wap *whatsAppChannel) SendInteractiveMultiproductMsg(
	ctx context.Context,
	msg models.InteractiveMultiproductMsg,
) (msgResp models.MsgResponse, respDetails models.ResponseDetails, err error) {
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
