package models

import (
	"time"

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

type MessageResponse struct {
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
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
	Filename string `json:"filename,omitempty" validate:"lte=240"`
}

func (t *DocumentMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type ImageMessage struct {
	MessageCommon
	Content ImageContent `json:"content" validate:"required"`
}

type ImageContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
}

func (t *ImageMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type AudioMessage struct {
	MessageCommon
	Content AudioContent `json:"content" validate:"required"`
}

type AudioContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
}

func (t *AudioMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type VideoMessage struct {
	MessageCommon
	Content VideoContent `json:"content" validate:"required"`
}

type VideoContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
}

func (t *VideoMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type StickerMessage struct {
	MessageCommon
	Content StickerContent `json:"content" validate:"required"`
}

type StickerContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
}

func (t *StickerMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type LocationMessage struct {
	MessageCommon
	Content LocationContent `json:"content" validate:"required"`
}

type LocationContent struct {
	Latitude  *float32 `json:"latitude" validate:"required,latitude"`
	Longitude *float32 `json:"longitude" validate:"required,longitude"`
	Name      string   `json:"name" validate:"lte=1000"`
	Address   string   `json:"address" validate:"lte=1000"`
}

func (t *LocationMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type ContactMessage struct {
	MessageCommon
	Content ContactContent `json:"content" validate:"required"`
}

type ContactContent struct {
	Contacts []Contact `json:"contacts" validate:"required,dive"`
}

type Contact struct {
	Addresses []ContactAddress `json:"addresses,omitempty" validate:"omitempty,dive"`
	Birthday  string           `json:"birthday,omitempty"`
	Emails    []ContactEmail   `json:"emails,omitempty" validate:"omitempty,dive"`
	Name      ContactName      `json:"name" validate:"required"`
	Org       ContactOrg       `json:"org,omitempty"`
	Phones    []ContactPhone   `json:"phones,omitempty" validate:"omitempty,dive"`
	Urls      []ContactURL     `json:"urls,omitempty" validate:"omitempty,dive"`
}

func (t *ContactMessage) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(BirthdayValidation, Contact{})
	return validate.Struct(t)
}

func BirthdayValidation(sl validator.StructLevel) {
	contact, _ := sl.Current().Interface().(Contact)
	if contact.Birthday == "" {
		return
	}
	_, err := time.Parse("2006-01-02", contact.Birthday)
	if err != nil {
		sl.ReportError(contact.Birthday, "birthday", "Contact", "invalidbirthdayformat", "")
	}
}

type ContactAddress struct {
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Type        string `json:"type,omitempty" validate:"omitempty,oneof=HOME WORK"`
}

type ContactEmail struct {
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	Type  string `json:"type,omitempty" validate:"omitempty,oneof=HOME WORK"`
}

type ContactName struct {
	FirstName     string `json:"firstName" validate:"required"`
	LastName      string `json:"lastName,omitempty"`
	MiddleName    string `json:"middleName,omitempty"`
	NameSuffix    string `json:"nameSuffix,omitempty"`
	NamePrefix    string `json:"namePrefix,omitempty"`
	FormattedName string `json:"formattedName" validate:"required"`
}

type ContactOrg struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type ContactPhone struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty" validate:"omitempty,oneof=CELL MAIN IPHONE HOME WORK"`
	WaID  string `json:"waId,omitempty"`
}

type ContactURL struct {
	URL  string `json:"url,omitempty" validate:"omitempty,url"`
	Type string `json:"type,omitempty" validate:"omitempty,oneof=HOME WORK"`
}
