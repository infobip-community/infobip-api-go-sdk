package models

import (
	"fmt"
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

func (t *TextMessage) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(PreviewURLValidation, TextContent{})
	return validate.Struct(t)
}

type TextMessage struct {
	MessageCommon
	Content TextContent `json:"content" validate:"required"`
}

type TextContent struct {
	Text       string `json:"text" validate:"required,gte=1,lte=4096"`
	PreviewURL bool   `json:"previewURL,omitempty"`
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

func (t *DocumentMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type DocumentContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
	Filename string `json:"filename,omitempty" validate:"lte=240"`
}

type ImageMessage struct {
	MessageCommon
	Content ImageContent `json:"content" validate:"required"`
}

func (t *ImageMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type ImageContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
}

type AudioMessage struct {
	MessageCommon
	Content AudioContent `json:"content" validate:"required"`
}

func (t *AudioMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type AudioContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
}

type VideoMessage struct {
	MessageCommon
	Content VideoContent `json:"content" validate:"required"`
}

func (t *VideoMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type VideoContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
	Caption  string `json:"caption,omitempty" validate:"lte=3000"`
}

type StickerMessage struct {
	MessageCommon
	Content StickerContent `json:"content" validate:"required"`
}

func (t *StickerMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type StickerContent struct {
	MediaURL string `json:"mediaUrl" validate:"required,url,lte=2048"`
}

type LocationMessage struct {
	MessageCommon
	Content LocationContent `json:"content" validate:"required"`
}

func (t *LocationMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type LocationContent struct {
	Latitude  *float32 `json:"latitude" validate:"required,latitude"`
	Longitude *float32 `json:"longitude" validate:"required,longitude"`
	Name      string   `json:"name" validate:"lte=1000"`
	Address   string   `json:"address" validate:"lte=1000"`
}

type ContactMessage struct {
	MessageCommon
	Content ContactContent `json:"content" validate:"required"`
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

type InteractiveButtonsMessage struct {
	MessageCommon
	Content InteractiveButtonsContent `json:"content" validate:"required"`
}

func (t *InteractiveButtonsMessage) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(ButtonHeaderValidation, InteractiveButtonsHeader{})
	return validate.Struct(t)
}

func ButtonHeaderValidation(sl validator.StructLevel) {
	header, _ := sl.Current().Interface().(InteractiveButtonsHeader)
	switch header.Type {
	case "TEXT":
		if header.Text == "" {
			sl.ReportError(header.Text, "text", "Text", "missingtext", "")
		}
	case "VIDEO", "IMAGE", "DOCUMENT":
		if header.MediaURL == "" {
			sl.ReportError(header.MediaURL, "mediaUrl", "MediaURL", "missingmediaurl", "")
		}
	}
}

type InteractiveButtonsContent struct {
	Body   InteractiveButtonsBody    `json:"body" validate:"required"`
	Action InteractiveButtons        `json:"action" validate:"required"`
	Header *InteractiveButtonsHeader `json:"header,omitempty" validate:"omitempty"`
	Footer *InteractiveButtonsFooter `json:"footer,omitempty"`
}

type InteractiveButtonsBody struct {
	Text string `json:"text" validate:"required,lte=1024"`
}

type InteractiveButtons struct {
	Buttons []InteractiveButton `json:"buttons" validate:"required,min=1,max=3,dive"`
}

type InteractiveButton struct {
	Type  string `json:"type" validate:"required,oneof=REPLY"`
	ID    string `json:"id" validate:"required,lte=256"`
	Title string `json:"title" validate:"required,lte=20"`
}

type InteractiveButtonsHeader struct {
	Type     string `json:"type" validate:"required,oneof=TEXT VIDEO IMAGE DOCUMENT"`
	Text     string `json:"text,omitempty" validate:"lte=60"`
	MediaURL string `json:"mediaUrl,omitempty" validate:"omitempty,url,lte=2048"`
	Filename string `json:"filename,omitempty" validate:"lte=240"`
}

type InteractiveButtonsFooter struct {
	Text string `json:"text" validate:"required,lte=60"`
}

type InteractiveListMessage struct {
	MessageCommon
	Content InteractiveListContent `json:"content" validate:"required"`
}

func (t *InteractiveListMessage) Validate() error {
	validate = validator.New()
	validate.RegisterStructValidation(InteractiveListActionValidation, InteractiveListAction{})
	return validate.Struct(t)
}

func InteractiveListActionValidation(sl validator.StructLevel) {
	action, _ := sl.Current().Interface().(InteractiveListAction)
	validateDuplicateRows(sl, action)
	validateSectionTitles(sl, action)
}

func validateDuplicateRows(sl validator.StructLevel, action InteractiveListAction) {
	rowIDs := make(map[string]int)
	for _, section := range action.Sections {
		for _, row := range section.Rows {
			rowIDs[row.ID]++
			if rowIDs[row.ID] > 1 {
				sl.ReportError(
					action.Sections,
					"sections",
					"Sections",
					fmt.Sprintf("duplicaterowID%s", row.ID),
					"",
				)
			}
		}
	}
}

func validateSectionTitles(sl validator.StructLevel, action InteractiveListAction) {
	if len(action.Sections) > 1 {
		for _, section := range action.Sections {
			if section.Title == "" {
				sl.ReportError(
					action.Sections,
					"sections",
					"Sections",
					"missingtitlemultiplesections",
					"",
				)
			}
		}
	}
}

type InteractiveListContent struct {
	Body   InteractiveListBody    `json:"body" validate:"required"`
	Action InteractiveListAction  `json:"action" validate:"required"`
	Header *InteractiveListHeader `json:"header,omitempty" validate:"omitempty"`
	Footer *InteractiveListFooter `json:"footer,omitempty"`
}

type InteractiveListBody struct {
	Text string `json:"text" validate:"required,lte=1024"`
}

type InteractiveListAction struct {
	Title    string    `json:"title" validate:"required,lte=20"`
	Sections []Section `json:"sections" validate:"required,min=1,max=10,dive"`
}

type Section struct {
	Title string       `json:"title,omitempty" validate:"lte=24"`
	Rows  []SectionRow `json:"rows" validate:"required,min=1,max=10,dive"`
}

type SectionRow struct {
	ID          string `json:"id" validate:"required,lte=200"`
	Title       string `json:"title" validate:"required,lte=24"`
	Description string `json:"description,omitempty" validate:"lte=72"`
}

type InteractiveListHeader struct {
	Type string `json:"type" validate:"required,oneof=TEXT"`
	Text string `json:"text" validate:"required,lte=60"`
}

type InteractiveListFooter struct {
	Text string `json:"text" validate:"required,lte=60"`
}

type InteractiveProductMessage struct {
	MessageCommon
	Content InteractiveProductContent `json:"content" validate:"required"`
}

func (t *InteractiveProductMessage) Validate() error {
	validate = validator.New()
	return validate.Struct(t)
}

type InteractiveProductContent struct {
	Action InteractiveProductAction  `json:"action" validate:"required"`
	Body   *InteractiveProductBody   `json:"body,omitempty"`
	Footer *InteractiveProductFooter `json:"footer,omitempty"`
}

type InteractiveProductAction struct {
	CatalogID         string `json:"catalogId" validate:"required"`
	ProductRetailerID string `json:"productRetailerId" validate:"required"`
}

type InteractiveProductBody struct {
	Text string `json:"text" validate:"required,lte=1024"`
}

type InteractiveProductFooter struct {
	Text string `json:"text" validate:"required,lte=60"`
}
