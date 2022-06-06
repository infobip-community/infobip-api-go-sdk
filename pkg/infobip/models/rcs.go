package models

import "bytes"

type RCSSuggestion struct {
	Text         string  `json:"text" validate:"min=1,max=25"`
	PostbackData string  `json:"postbackData" validate:"min=1,max=2048"`
	Type         string  `json:"type" validate:"oneof=REPLY OPEN_URL DIAL_PHONE SHOW_LOCATION REQUEST_LOCATION"`
	URL          string  `json:"url,omitempty" validate:"required_if=Type OPEN_URL"`
	PhoneNumber  string  `json:"phoneNumber,omitempty"`
	Latitude     float64 `json:"latitude,omitempty" validate:"min=-90,max=90,required_if=Type SHOW_LOCATION"`
	Longitude    float64 `json:"longitude,omitempty" validate:"min=-180,max=180,required_if=Type SHOW_LOCATION"`
	Label        string  `json:"label,omitempty" validate:"omitempty,min=1,max=100"`
}

func (r *RCSSuggestion) Validate() error {
	return validate.Struct(r)
}

func (r *RCSSuggestion) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type RCSFile struct {
	URL string `json:"url" validate:"required,min=1,max=1000"`
}

type RCSThumbnail struct {
	URL string `json:"url" validate:"required,min=1,max=1000"`
}

type RCSCardContent struct {
	Title       string               `json:"title" validate:"omitempty,min=1,max=200"`
	Description string               `json:"description" validate:"omitempty,min=1,max=2000"`
	Media       *RCSCardContentMedia `json:"media,omitempty"`
	Suggestions []RCSSuggestion      `json:"suggestions,omitempty" validate:"max=4"`
}

type RCSCardContentMedia struct {
	File      *RCSFile      `json:"file"`
	Thumbnail *RCSThumbnail `json:"thumbnail"`
	Height    string        `json:"height" validate:"omitempty,oneof=SHORT MEDIUM TALL"`
}

type RCSContent struct {
	Type        string           `json:"type" validate:"oneof=TEXT FILE CARD CAROUSEL"`
	File        *RCSFile         `json:"file,omitempty" validate:"required_if=Type FILE"`
	Thumbnail   *RCSThumbnail    `json:"thumbnail,omitempty"`
	Text        string           `json:"text" validator:"min=1,max=2048,required_if=Type TEXT"`
	Suggestions []RCSSuggestion  `json:"suggestions"`
	Orientation string           `json:"orientation" validate:"required_if=Type CARD,oneof=HORIZONTAL VERTICAL"`
	Alignment   string           `json:"alignment" validate:"required_if=Type CARD,oneof=LEFT RIGHT"`
	CardWidth   string           `json:"cardWidth" validate:"required_if=Type CAROUSEL,oneof=SMALL MEDIUM"`
	Contents    []RCSCardContent `json:"contents" validate:"required_if=Type CAROUSEL,min=2,max=10"`
}

type RCSSMSFailover struct {
	From                   string `json:"from" validate:"required"`
	Text                   string `json:"text" validate:"required"`
	ValidityPeriod         int    `json:"validityPeriod"`
	ValidityPeriodTimeUnit string `json:"validityPeriodTimeUnit" validate:"omitempty,oneof=SECONDS MINUTES HOURS DAYS"`
}

type RCSMsg struct {
	From                   string          `json:"from"`
	To                     string          `json:"to" validate:"required"`
	ValidityPeriod         int             `json:"validityPeriod"`
	ValidityPeriodTimeUnit string          `json:"validityPeriodTimeUnit" validate:"omitempty,oneof=SECONDS MINUTES HOURS DAYS"` //nolint:lll
	Content                *RCSContent     `json:"content" validate:"required"`
	SMSFailover            *RCSSMSFailover `json:"smsFailover"`
	NotifyURL              string          `json:"notifyUrl"`
	CallbackData           string          `json:"callbackData"`
	MessageID              string          `json:"messageId"`
}

func (r *RCSMsg) Validate() error {
	return validate.Struct(r)
}

func (r *RCSMsg) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type SendRCSResponse struct {
	Messages []struct {
		To           string `json:"to"`
		MessageCount int    `json:"messageCount"`
		MessageID    string `json:"messageId"`
		Status       struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
		} `json:"status"`
	} `json:"messages"`
}

type SendRCSBulkResponse SendRCSResponse

type SendRCSBulkRequest struct {
	Messages []RCSMsg `json:"messages"`
}

func (s *SendRCSBulkRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendRCSBulkRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}
