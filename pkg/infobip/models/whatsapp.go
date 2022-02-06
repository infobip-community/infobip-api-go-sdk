package models

import (
	"github.com/go-playground/validator/v10"
	"mvdan.cc/xurls/v2"
)

type TextMessageRequest struct {
	From         string  `json:"from" validate:"required,lte=24"`
	To           string  `json:"to" validate:"required,lte=24"`
	MessageID    string  `json:"messageId,omitempty" validate:"lte=50"`
	Content      Content `json:"content" validate:"required"`
	CallbackData string  `json:"callbackData,omitempty" validate:"lte=4000"`
	NotifyURL    string  `json:"notifyUrl,omitempty" validate:"url,lte=2048"`
}

type Content struct {
	Text       string `json:"text" validate:"required,gte=1,lte=4096"`
	PreviewURL bool   `json:"previewURL,omitempty"`
}

func (t *TextMessageRequest) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(PreviewURLValidation, Content{})
	return validate.Struct(t)
}

func PreviewURLValidation(sl validator.StructLevel) {
	content, _ := sl.Current().Interface().(Content)
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
	GroupID     string `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
}
