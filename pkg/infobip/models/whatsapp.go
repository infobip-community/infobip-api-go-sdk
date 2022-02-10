package models

import (
	"github.com/go-playground/validator/v10"
	"mvdan.cc/xurls/v2"
)

type MessageCommon struct {
	From         string `json:"from" validate:"required,lte=24"`
	To           string `json:"to" validate:"required,lte=24"`
	MessageID    string `json:"messageId,omitempty" validate:"lte=50"`
	CallbackData string `json:"callbackData,omitempty" validate:"lte=4000"`
	NotifyURL    string `json:"notifyUrl,omitempty" validate:"omitempty,url,lte=2048"`
}

type TextMessage struct {
	MessageCommon
	Content TextContent `json:"content" validate:"required"`
}

type TextContent struct {
	Text       string `json:"text" validate:"required,gte=1,lte=4096"`
	PreviewURL bool   `json:"previewURL,omitempty"`
}

func (t *TextMessage) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(PreviewURLValidation, TextContent{})
	return validate.Struct(t)
}

func PreviewURLValidation(sl validator.StructLevel) {
	content, _ := sl.Current().Interface().(TextContent)
	containsURL := xurls.Relaxed().FindString(content.Text)
	if content.PreviewURL && containsURL == "" {
		sl.ReportError(content.Text, "text", "Text", "missingurlintext", "")
	}
}

type TextMessageResponse struct {
	To           string `json:"to"`
	MessageCount int32  `json:"messageCount"`
	MessageID    string `json:"messageId"`
	Status       Status `json:"status"`
}

type Status struct {
	GroupID     int32  `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

type DocumentMessage struct {
	MessageCommon
	Content DocumentContent `json:"content" validate:"required"`
}

type DocumentContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=240"`
	Filename string `json:"filename,omitempty" validate:"lte=240"`
}

func (t *DocumentMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}
