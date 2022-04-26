package models

import (
	"bytes"
)

type Destination struct {
	MessageID string `json:"messageId"`
	To        string `json:"to" validate:"required"`
}

type Language struct {
	LanguageCode string `json:"languageCode"`
}

type SMSTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type SMSDeliveryTimeWindow struct {
	Days []string `json:"days" validate:"required"`
	From SMSTime  `json:"from,omitempty"`
	To   SMSTime  `json:"to,omitempty"`
}

type SMSMsg struct {
	CallbackData           string        `json:"callbackData,omitempty"`
	Destinations           []Destination `json:"destinations"`
	Flash                  bool          `json:"flash,omitempty"`
	From                   string        `json:"from"`
	IntermediateReport     bool          `json:"intermediateReport,omitempty"`
	Language               `json:"language,omitempty"`
	NotifyContentType      string `json:"notifyContentType,omitempty"`
	NotifyURL              string `json:"notifyUrl,omitempty"`
	Text                   string `json:"text"`
	Transliteration        string `json:"transliteration,omitempty"`
	ValidityPeriod         int    `json:"validityPeriod,omitempty"`
	*SMSDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	SendAt                 string `json:"sendAt,omitempty"`
}

func (s *SMSMsg) Validate() error {
	return validate.Struct(s)
}

func (s *SMSMsg) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SendSMSRequest struct {
	BulkID   string   `json:"bulkId"`
	Messages []SMSMsg `json:"messages" validate:"required"`
	Tracking struct {
		Track string `json:"track"`
		Type  string `json:"type"`
	} `json:"tracking"`
}

func (s *SendSMSRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SendSMSResponse struct {
	BulkID   string `json:"bulkId"`
	Messages []struct {
		MessageID string `json:"messageId"`
		Status    struct {
			Description string `json:"description"`
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"status"`
		To string `json:"to"`
	} `json:"messages"`
}
