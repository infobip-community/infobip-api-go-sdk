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
	CallbackData           string `json:"callbackData,omitempty"`
	*SMSDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	Destinations           []Destination `json:"destinations"`
	Flash                  bool          `json:"flash,omitempty"`
	From                   string        `json:"from"`
	IntermediateReport     bool          `json:"intermediateReport,omitempty"`
	*Language              `json:"language,omitempty"`
	NotifyContentType      string `json:"notifyContentType,omitempty"`
	NotifyURL              string `json:"notifyUrl,omitempty"`
	*Regional              `json:"regional,omitempty"`
	SendAt                 string `json:"sendAt,omitempty"`
	Text                   string `json:"text"`
	Transliteration        string `json:"transliteration,omitempty"`
	ValidityPeriod         int    `json:"validityPeriod,omitempty"`
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
			Action      string `json:"action"`
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

func (g *GetSMSDeliveryReportsParams) Validate() error {
	return validate.Struct(g)
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
			GroupID     int    `json:"groupId"`
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

func (g *GetSMSLogsParams) Validate() error {
	return validate.Struct(g)
}

type Binary struct {
	Hex        string `json:"hex" validate:"required"`
	DataCoding int    `json:"dataCoding"`
	EsmClass   int    `json:"esmClass"`
}

type IndiaDLT struct {
	ContentTemplateID string `json:"contentTemplateId"`
	PrincipalEntityID string `json:"principalEntityId" validate:"required"`
}

type Regional struct {
	IndiaDLT `json:"indiaDlt"`
}

type BinarySMSMsg struct {
	From               string        `json:"from"`
	Destinations       []Destination `json:"destinations"`
	*Binary            `json:"binary"`
	IntermediateReport bool                   `json:"intermediateReport,omitempty"`
	NotifyURL          string                 `json:"notifyUrl,omitempty"`
	NotifyContentType  string                 `json:"notifyContentType,omitempty"`
	CallbackData       string                 `json:"callbackData,omitempty"`
	ValidityPeriod     int                    `json:"validityPeriod,omitempty"`
	SendAt             string                 `json:"sendAt,omitempty"`
	DeliveryTimeWindow *SMSDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	*Regional          `json:"regional,omitempty"`
}

type SendBinarySMSRequest struct {
	BulkID             string         `json:"bulkId"`
	Messages           []BinarySMSMsg `json:"messages" validate:"required"`
	*SendingSpeedLimit `json:"sendingSpeedLimit,omitempty"`
}

func (s *SendBinarySMSRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendBinarySMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SendBinarySMSResponse struct {
	BulkID   string `json:"bulkId"`
	Messages []struct {
		To     string `json:"to"`
		Status struct {
			Action      string `json:"action"`
			Description string `json:"description"`
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"status"`
		MessageID string `json:"messageId"`
	} `json:"messages"`
}

type SendSMSOverQueryParamsParams struct {
	Username                  string `validate:"required"`
	Password                  string `validate:"required"`
	BulkID                    string
	From                      string
	To                        []string `validate:"required"`
	Text                      string
	Flash                     bool
	Transliteration           string
	LanguageCode              string
	IntermediateReport        bool
	NotifyURL                 string
	NotifyContentType         string
	CallbackData              string
	ValidityPeriod            int
	SendAt                    string
	Track                     string
	ProcessKey                string
	TrackingType              string
	IndiaDLTContentTemplateID string
	IndiaDLTPrincipalEntityID string
}

func (s *SendSMSOverQueryParamsParams) Validate() error {
	return validate.Struct(s)
}

type SendSMSOverQueryParamsResponse struct {
	BulkID   string `json:"bulkId"`
	Messages []struct {
		To     string `json:"to"`
		Status struct {
			Action      string `json:"action"`
			Description string `json:"description"`
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"status"`
		MessageID string `json:"messageId"`
	} `json:"messages"`
}

type PreviewSMSRequest struct {
	LanguageCode    string `json:"languageCode" validation:"oneof=TR ES PT AUTODETECT"`
	Text            string `json:"text"`
	Transliteration string `json:"transliteration" validation:"oneof=TURKISH GREEK CYRILLIC SERBIAN_CYRILLIC CENTRAL_EUROPEAN BALTIC NON_UNICODE"`
}

func (r *PreviewSMSRequest) Validate() error {
	return validate.Struct(r)
}

func (r *PreviewSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type PreviewSMSResponse struct {
	OriginalText string `json:"originalText"`
	Previews     []struct {
		CharactersRemaining int `json:"charactersRemaining"`
		Configuration       struct {
			Language        `json:"language"`
			Transliteration string `json:"transliteration"`
		} `json:"configuration"`
		MessageCount int    `json:"messageCount"`
		TextPreview  string `json:"textPreview"`
	} `json:"previews"`
}

type GetInboundSMSParams struct {
	Limit int `validate:"min=1,max=1000"`
}

func (g *GetInboundSMSParams) Validate() error {
	return validate.Struct(g)
}

type GetInboundSMSResponse struct {
	MessageCount        int `json:"messageCount"`
	PendingMessageCount int `json:"pendingMessageCount"`
	Results             []struct {
		CallbackData string `json:"callbackData"`
		CleanText    string `json:"cleanText"`
		From         string `json:"from"`
		Keyword      string `json:"keyword"`
		MessageID    string `json:"messageId"`
		Price        struct {
			PricePerMessage float64 `json:"pricePerMessage"`
			Currency        string  `json:"currency"`
		} `json:"price"`
		ReceivedAt string `json:"receivedAt"`
		SmsCount   int    `json:"smsCount"`
		Text       string `json:"text"`
		To         string `json:"to"`
	} `json:"results"`
}

type GetScheduledSMSParams struct {
	BulkID string `validate:"required"`
}

func (g *GetScheduledSMSParams) Validate() error {
	return validate.Struct(g)
}

type GetScheduledSMSResponse struct {
	BulkID string `json:"bulkId"`
	SendAt string `json:"sendAt"`
}

type RescheduleSMSParams struct {
	BulkID string `json:"bulkId" validate:"required"`
}

func (r *RescheduleSMSParams) Validate() error {
	return validate.Struct(r)
}

type RescheduleSMSRequest struct {
	SendAt string `json:"sendAt" validate:"required"`
}

func (r *RescheduleSMSRequest) Validate() error {
	return validate.Struct(r)
}

func (r *RescheduleSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type RescheduleSMSResponse struct {
	BulkID string `json:"bulkId"`
	SendAt string `json:"sendAt"`
}

type GetScheduledSMSStatusParams struct {
	BulkID string `json:"bulkId" validate:"required"`
}

func (g *GetScheduledSMSStatusParams) Validate() error {
	return validate.Struct(g)
}

type GetScheduledSMSStatusResponse struct {
	BulkID string `json:"bulkId"`
	Status string `json:"sendAt"`
}

type UpdateScheduledSMSStatusParams struct {
	BulkID string `json:"bulkId" validate:"required"`
}

func (u *UpdateScheduledSMSStatusParams) Validate() error {
	return validate.Struct(u)
}

type UpdateScheduledSMSStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=PENDING PAUSED PROCESSING CANCELED FINISHED FAILED"`
}

func (u *UpdateScheduledSMSStatusRequest) Validate() error {
	return validate.Struct(u)
}

func (u *UpdateScheduledSMSStatusRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

type UpdateScheduledSMSStatusResponse struct {
	BulkID string `json:"bulkId"`
	Status string `json:"status"`
}
