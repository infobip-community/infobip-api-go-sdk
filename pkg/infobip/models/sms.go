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

type Tracking struct {
	BaseURL    string `json:"baseUrl"`
	Track      string `json:"track"`
	Type       string `json:"type" validate:"oneof=ONE_TIME_PIN SOCIAL_INVITES"`
	ProcessKey string `json:"processKey"`
}

type SendingSpeedLimit struct {
	Amount   int    `json:"amount" validate:"required"`
	TimeUnit string `json:"timeUnit" validate:"oneof=MINUTE HOUR DAY"`
}

type SendSMSRequest struct {
	BulkID             string   `json:"bulkId"`
	Messages           []SMSMsg `json:"messages" validate:"required"`
	*SendingSpeedLimit `json:"sendingSpeedLimit,omitempty"`
	*Tracking          `json:"tracking,omitempty"`
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

type GetSMSDeliveryReportsParams struct {
	BulkID    string
	MessageID string
	Limit     int
}

type GetSMSDeliveryReportsResponse struct {
	Results []struct {
		BulkID    string `json:"bulkId"`
		MessageID string `json:"messageId"`
		To        string `json:"to"`
		From      string `json:"from"`
		SentAt    string `json:"sentAt"`
		DoneAt    string `json:"doneAt"`
		SmsCount  int    `json:"smsCount"`
		MccMnc    string `json:"mccMnc"`
		Price     struct {
			PricePerMessage float64 `json:"pricePerMessage"`
			Currency        string  `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
		Error struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Permanent   bool   `json:"permanent"`
		} `json:"error"`
	} `json:"results"`
}

type GetSMSLogsResponse struct {
	Results []struct {
		BulkID    string `json:"bulkId"`
		MessageID string `json:"messageId"`
		To        string `json:"to"`
		From      string `json:"from"`
		Text      string `json:"text"`
		SentAt    string `json:"sentAt"`
		DoneAt    string `json:"doneAt"`
		SmsCount  int    `json:"smsCount"`
		MccMnc    string `json:"mccMnc"`
		Price     struct {
			PricePerMessage float64 `json:"pricePerMessage"`
			Currency        string  `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
		Error struct {
			GroupId     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Permanent   bool   `json:"permanent"`
		} `json:"error"`
	} `json:"results"`
}

type GetSMSLogsParams struct {
	From          string
	To            string
	BulkID        []string
	MessageID     []string
	GeneralStatus string `validate:"oneof=ACCEPTED PENDING UNDELIVERABLE DELIVERED REJECTED EXPIRED"`
	SentSince     string
	SentUntil     string
	Limit         int `validate:"min=1,max=1000"`
	MCC           string
	MNC           string
}
