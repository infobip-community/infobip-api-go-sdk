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
	Suggestions []RCSSuggestion      `json:"suggestions,omitempty" validate:"omitempty,max=4,dive"`
}

func (r *RCSCardContent) Validate() error {
	return validate.Struct(r)
}

func (r *RCSCardContent) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type RCSCardContentMedia struct {
	File      *RCSFile      `json:"file" validate:"required"`
	Thumbnail *RCSThumbnail `json:"thumbnail"`
	Height    string        `json:"height" validate:"required,oneof=SHORT MEDIUM TALL"`
}

type RCSContent struct {
	Type        string           `json:"type" validate:"oneof=TEXT FILE CARD CAROUSEL"`
	File        *RCSFile         `json:"file,omitempty" validate:"required_if=Type FILE,omitempty"`
	Thumbnail   *RCSThumbnail    `json:"thumbnail,omitempty"`
	Text        string           `json:"text,omitempty" validator:"required_if=Type TEXT,omitempty,min=1,max=2048"`
	Suggestions []RCSSuggestion  `json:"suggestions,omitempty" validate:"omitempty,dive"`
	Orientation string           `json:"orientation,omitempty" validate:"required_if=Type CARD,omitempty,oneof=HORIZONTAL VERTICAL"` //nolint:lll
	Alignment   string           `json:"alignment,omitempty" validate:"required_if=Type CARD,omitempty,oneof=LEFT RIGHT"`            //nolint:lll
	CardWidth   string           `json:"cardWidth,omitempty" validate:"required_if=Type CAROUSEL,omitempty,oneof=SMALL MEDIUM"`      //nolint:lll
	Content     *RCSCardContent  `json:"content,omitempty" validate:"required_if=Type CARD"`
	Contents    []RCSCardContent `json:"contents,omitempty" validate:"required_if=Type CAROUSEL,omitempty,min=2,max=10,dive"` //nolint:lll
}

type RCSSMSFailover struct {
	From                   string `json:"from" validate:"required"`
	Text                   string `json:"text" validate:"required"`
	ValidityPeriod         int    `json:"validityPeriod,omitempty"`
	ValidityPeriodTimeUnit string `json:"validityPeriodTimeUnit,omitempty" validate:"omitempty,oneof=SECONDS MINUTES HOURS DAYS"` //nolint:lll
}

type RCSMsg struct {
	From                   string          `json:"from,omitempty"`
	To                     string          `json:"to" validate:"required"`
	ValidityPeriod         int             `json:"validityPeriod,omitempty"`
	ValidityPeriodTimeUnit string          `json:"validityPeriodTimeUnit,omitempty" validate:"omitempty,oneof=SECONDS MINUTES HOURS DAYS"` //nolint:lll
	Content                *RCSContent     `json:"content,omitempty" validate:"required"`
	SMSFailover            *RCSSMSFailover `json:"smsFailover,omitempty"`
	NotifyURL              string          `json:"notifyUrl,omitempty"`
	CallbackData           string          `json:"callbackData,omitempty"`
	MessageID              string          `json:"messageId,omitempty"`
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

type SendRCSBulkResponse []SendRCSResponse

type SendRCSBulkRequest struct {
	Messages []RCSMsg `json:"messages" validate:"dive"`
}

func (s *SendRCSBulkRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendRCSBulkRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}
