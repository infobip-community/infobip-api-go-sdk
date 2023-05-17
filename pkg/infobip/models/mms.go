package models

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

const MinDeliveryWindow = 60

func setupMMSValidations() {
	if validate == nil {
		validate = validator.New()
	}
	validate.RegisterStructValidation(MMSHeadValidation, MMSHead{})
}

type MMSMsg struct {
	Head                  MMSHead `validate:"required"`
	Text                  string
	Media                 *os.File
	ExternallyHostedMedia []ExternallyHostedMedia `validate:"dive"`
	SMIL                  string
	boundary              string
}

type MMSHead struct {
	From                  string              `json:"from" validate:"required"`
	To                    string              `json:"to" validate:"required"`
	ID                    string              `json:"id,omitempty"`
	Subject               string              `json:"subject,omitempty"`
	ValidityPeriodMinutes int32               `json:"validityPeriodMinutes,omitempty"`
	CallbackData          string              `json:"callbackData,omitempty" validate:"lte=200"`
	NotifyURL             string              `json:"notifyUrl,omitempty" validate:"omitempty,url"`
	SendAt                string              `json:"sendAt,omitempty"`
	IntermediateReport    *bool               `json:"intermediateReport,omitempty"`
	DeliveryTimeWindow    *DeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
}

type DeliveryTimeWindow struct {
	Days []string `json:"days" validate:"required,gte=1,dive,oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"` //nolint: lll
	From *MMSTime `json:"from,omitempty"`
	To   *MMSTime `json:"to,omitempty"`
}

type MMSTime struct {
	Hour   int32 `json:"hour" validate:"lte=23"`
	Minute int32 `json:"minute" validate:"lte=59"`
}

type ExternallyHostedMedia struct {
	ContentType string `json:"contentType" validate:"required"`
	ContentID   string `json:"contentId" validate:"required"`
	ContentURL  string `json:"contentUrl" validate:"url,required"`
}

func MMSHeadValidation(sl validator.StructLevel) {
	head, _ := sl.Current().Interface().(MMSHead)
	validateSendAt(sl, head)
	if head.DeliveryTimeWindow != nil && (head.DeliveryTimeWindow.From != nil || head.DeliveryTimeWindow.To != nil) {
		validateDeliveryTimeWindow(sl, head)
	}
}

func validateSendAt(sl validator.StructLevel, head MMSHead) {
	if head.SendAt == "" {
		return
	}
	_, err := time.Parse(time.RFC3339, head.SendAt)
	if err != nil {
		sl.ReportError(head.SendAt, "sendAt", "SendAt", "invalidformat", "")
	}
}

func validateDeliveryTimeWindow(sl validator.StructLevel, head MMSHead) {
	if head.DeliveryTimeWindow.From != nil && head.DeliveryTimeWindow.To == nil {
		sl.ReportError(head.DeliveryTimeWindow, "deliveryTimeWindow", "DeliveryTimeWindow", "missingto", "")
		return
	} else if head.DeliveryTimeWindow.To != nil && head.DeliveryTimeWindow.From == nil {
		sl.ReportError(head.DeliveryTimeWindow, "deliveryTimeWindow", "DeliveryTimeWindow", "missingfrom", "")
		return
	}

	if !deliveryWindowIsValid(*head.DeliveryTimeWindow.From, *head.DeliveryTimeWindow.To) {
		sl.ReportError(head.DeliveryTimeWindow, "deliveryTimeWindow", "DeliveryTimeWindow", "fromtonot1hourapart", "")
	}
}

func deliveryWindowIsValid(from MMSTime, to MMSTime) bool {
	fromMinutes := from.Hour*MinsPerHour + from.Minute
	toMinutes := to.Hour*MinsPerHour + to.Minute
	diff := toMinutes - fromMinutes
	return diff >= MinDeliveryWindow
}

type SendMMSResponse struct {
	BulkID       string    `json:"bulkId"`
	Messages     []SentMMS `json:"messages"`
	ErrorMessage string    `json:"errorMessage"`
}

type SentMMS struct {
	To        string    `json:"to"`
	Status    MMSStatus `json:"status"`
	MessageID string    `json:"messageId"`
}

type MMSStatus struct {
	GroupID     int32  `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *MMSMsg) Validate() error {
	return validate.Struct(t)
}

func (t *MMSMsg) Marshal() (*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	multipartWriter.Boundary()
	var partWriter io.Writer
	err := writeMultipartJSON(multipartWriter, "head", t.Head)
	if err != nil {
		return nil, err
	}

	if t.Text != "" {
		err = writeMultipartText(multipartWriter, "text", t.Text)
		if err != nil {
			return nil, err
		}
	}

	if t.Media != nil {
		defer t.Media.Close()
		if partWriter, err = multipartWriter.CreateFormFile("media", t.Media.Name()); err != nil {
			return nil, err
		}
		if _, err = io.Copy(partWriter, t.Media); err != nil {
			return nil, err
		}
	}

	if len(t.ExternallyHostedMedia) > 0 {
		err = writeMultipartJSON(multipartWriter, "externallyHostedMedia", t.ExternallyHostedMedia)
		if err != nil {
			return nil, err
		}
	}

	if t.SMIL != "" {
		err = writeMultipartXMLString(multipartWriter, "smil", t.SMIL)
		if err != nil {
			return nil, err
		}
	}

	multipartWriter.Close()
	t.boundary = multipartWriter.Boundary()
	return &buf, nil
}

func (t *MMSMsg) GetMultipartBoundary() string {
	return t.boundary
}

type GetMMSDeliveryReportsParams struct {
	BulkID    string
	MessageID string
	Limit     int
}

type GetMMSDeliveryReportsResponse struct {
	Results []OutboundMMSDeliveryResult `json:"results"`
}

type OutboundMMSDeliveryResult struct {
	BulkID       string    `json:"bulkId"`
	MessageID    string    `json:"messageId"`
	To           string    `json:"to"`
	From         string    `json:"from"`
	SentAt       string    `json:"sentAt"`
	DoneAt       string    `json:"doneAt"`
	MMSCount     int32     `json:"mmsCount"`
	MCCMNC       string    `json:"mccMnc"`
	CallbackData string    `json:"callbackData"`
	Price        MMSPrice  `json:"price"`
	Status       MMSStatus `json:"status"`
	Error        MMSStatus `json:"error"`
}

type MMSPrice struct {
	PricePerMessage float64 `json:"pricePerMessage"`
	Currency        string  `json:"currency"`
}

type GetInboundMMSResponse struct {
	Results []InboundMMSResult `json:"results"`
}

type InboundMMSResult struct {
	MessageID    string   `json:"messageId"`
	To           string   `json:"to"`
	From         string   `json:"from"`
	Message      string   `json:"message"`
	ReceivedAt   string   `json:"receivedAt"`
	MMSCount     int32    `json:"mmsCount"`
	CallbackData string   `json:"callbackData"`
	Price        MMSPrice `json:"price"`
}

type GetInboundMMSParams struct {
	Limit int
}
