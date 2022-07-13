package whatsapp

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

// WhatsApp provides methods to interact with the Infobip WhatsApp API.
// WhatsApp API docs: https://www.infobip.com/docs/api#channels/whatsapp
type WhatsApp interface {
	SendTemplate(context.Context, models.WATemplateMsgs) (models.BulkWAMsgResponse, models.ResponseDetails, error)
	SendText(context.Context, models.WATextMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendDocument(context.Context, models.WADocumentMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendImage(context.Context, models.WAImageMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendAudio(context.Context, models.WAAudioMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendVideo(context.Context, models.WAVideoMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendSticker(context.Context, models.WAStickerMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendLocation(context.Context, models.WALocationMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendContact(context.Context, models.WAContactMsg) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendInteractiveButtons(context.Context, models.WAInteractiveButtonsMsg,
	) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendInteractiveList(context.Context, models.WAInteractiveListMsg,
	) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendInteractiveProduct(context.Context, models.WAInteractiveProductMsg,
	) (models.SendWAMsgResponse, models.ResponseDetails, error)
	SendInteractiveMultiproduct(context.Context, models.WAInteractiveMultiproductMsg,
	) (models.SendWAMsgResponse, models.ResponseDetails, error)
	GetTemplates(context.Context, string) (models.GetWATemplatesResponse, models.ResponseDetails, error)
	CreateTemplate(context.Context, string, models.TemplateCreate,
	) (models.CreateWATemplateResponse, models.ResponseDetails, error)
	DeleteTemplate(context.Context, string, string,
	) (models.ResponseDetails, error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

const (
	sendTemplateMessagesPath        = "whatsapp/1/message/template"
	sendMessagePath                 = "whatsapp/1/message/text"
	sendDocumentPath                = "whatsapp/1/message/document"
	sendImagePath                   = "whatsapp/1/message/image"
	sendAudioPath                   = "whatsapp/1/message/audio"
	sendVideoPath                   = "whatsapp/1/message/video"
	sendStickerPath                 = "whatsapp/1/message/sticker"
	sendLocationPath                = "whatsapp/1/message/location"
	sendContactPath                 = "whatsapp/1/message/contact"
	sendInteractiveButtonsPath      = "whatsapp/1/message/interactive/buttons"
	sendInteractiveListPath         = "whatsapp/1/message/interactive/list"
	sendInteractiveProductPath      = "whatsapp/1/message/interactive/product"
	sendInteractiveMultiproductPath = "whatsapp/1/message/interactive/multi-product"
	templatesPath                   = "whatsapp/2/senders/%s/templates"
	deleteTemplatePath              = "whatsapp/2/senders/%s/templates/%s"
)

func (wap *Channel) SendTemplate(
	ctx context.Context,
	messages models.WATemplateMsgs,
) (msgResp models.BulkWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &messages, &msgResp, sendTemplateMessagesPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendText(
	ctx context.Context,
	msg models.WATextMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendMessagePath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendDocument(
	ctx context.Context,
	msg models.WADocumentMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendDocumentPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendImage(
	ctx context.Context,
	msg models.WAImageMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendImagePath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendAudio(
	ctx context.Context,
	msg models.WAAudioMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendAudioPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendVideo(
	ctx context.Context,
	msg models.WAVideoMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendVideoPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendSticker(
	ctx context.Context,
	msg models.WAStickerMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendStickerPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendLocation(
	ctx context.Context,
	msg models.WALocationMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendLocationPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendContact(
	ctx context.Context,
	msg models.WAContactMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendContactPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendInteractiveButtons(
	ctx context.Context,
	msg models.WAInteractiveButtonsMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendInteractiveButtonsPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendInteractiveList(
	ctx context.Context,
	msg models.WAInteractiveListMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendInteractiveListPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendInteractiveProduct(
	ctx context.Context,
	msg models.WAInteractiveProductMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendInteractiveProductPath)
	return msgResp, respDetails, err
}

func (wap *Channel) SendInteractiveMultiproduct(
	ctx context.Context,
	msg models.WAInteractiveMultiproductMsg,
) (msgResp models.SendWAMsgResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &msg, &msgResp, sendInteractiveMultiproductPath)
	return msgResp, respDetails, err
}

func (wap *Channel) GetTemplates(
	ctx context.Context,
	sender string,
) (resp models.GetWATemplatesResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.GetRequest(ctx, &resp, fmt.Sprintf(templatesPath, sender), nil)
	return resp, respDetails, err
}

func (wap *Channel) CreateTemplate(
	ctx context.Context,
	sender string,
	template models.TemplateCreate,
) (resp models.CreateWATemplateResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.PostJSONReq(ctx, &template, &resp, fmt.Sprintf(templatesPath, sender))
	return resp, respDetails, err
}

func (wap *Channel) DeleteTemplate(
	ctx context.Context,
	sender string,
	templateName string,
) (respDetails models.ResponseDetails, err error) {
	respDetails, err = wap.ReqHandler.DeleteRequest(ctx, fmt.Sprintf(deleteTemplatePath, sender, templateName), nil)
	return respDetails, err
}
