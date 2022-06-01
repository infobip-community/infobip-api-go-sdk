package models

import (
	"bytes"
)

type SMSDestination struct {
	MessageID string `json:"messageId"`
	To        string `json:"to" validate:"required"`
}

type SMSLanguage struct {
	LanguageCode string `json:"languageCode"`
}

type SMSTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type SMSDeliveryTimeWindow struct {
	Days []string `json:"days" validate:"required,min=1"`
	From SMSTime  `json:"from,omitempty"`
	To   SMSTime  `json:"to,omitempty"`
}

type SMSMsg struct {
	CallbackData       string                 `json:"callbackData,omitempty"`
	DeliveryTimeWindow *SMSDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	Destinations       []SMSDestination       `json:"destinations" validate:"required,min=1"`
	Flash              bool                   `json:"flash,omitempty"`
	From               string                 `json:"from"`
	IntermediateReport bool                   `json:"intermediateReport,omitempty"`
	Language           *SMSLanguage           `json:"language,omitempty"`
	NotifyContentType  string                 `json:"notifyContentType,omitempty"`
	NotifyURL          string                 `json:"notifyUrl,omitempty"`
	Regional           *SMSRegional           `json:"regional,omitempty"`
	SendAt             string                 `json:"sendAt,omitempty"`
	Text               string                 `json:"text"`
	Transliteration    string                 `json:"transliteration,omitempty"`
	ValidityPeriod     int                    `json:"validityPeriod,omitempty"`
}

func (s *SMSMsg) Validate() error {
	return validate.Struct(s)
}

func (s *SMSMsg) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SMSTracking struct {
	BaseURL    string `json:"baseUrl"`
	Track      string `json:"track"`
	Type       string `json:"type" validate:"oneof=ONE_TIME_PIN SOCIAL_INVITES"`
	ProcessKey string `json:"processKey"`
}

type SMSSendingSpeedLimit struct {
	Amount   int    `json:"amount" validate:"required"`
	TimeUnit string `json:"timeUnit" validate:"oneof=MINUTE HOUR DAY"`
}

type SendSMSRequest struct {
	BulkID            string                `json:"bulkId"`
	Messages          []SMSMsg              `json:"messages" validate:"required,min=1"`
	SendingSpeedLimit *SMSSendingSpeedLimit `json:"sendingSpeedLimit,omitempty"`
	Tracking          *SMSTracking          `json:"tracking,omitempty"`
}

func (s *SendSMSRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SMSStatus struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type SMSPrice struct {
	PricePerMessage float64 `json:"pricePerMessage"`
	Currency        string  `json:"currency"`
}

type SendSMSResponse struct {
	BulkID   string `json:"bulkId"`
	Messages []struct {
		MessageID string     `json:"messageId"`
		Status    *SMSStatus `json:"status"`
		To        string     `json:"to"`
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

type SMSError struct {
	Description string `json:"description"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permanent   bool   `json:"permanent"`
}

type GetSMSDeliveryReportsResponse struct {
	Results []struct {
		BulkID       string    `json:"bulkId"`
		CallbackData string    `json:"callbackData"`
		DoneAt       string    `json:"doneAt"`
		Error        SMSError  `json:"error"`
		From         string    `json:"from"`
		MccMnc       string    `json:"mccMnc"`
		MessageID    string    `json:"messageId"`
		Price        SMSPrice  `json:"price"`
		SentAt       string    `json:"sentAt"`
		SmsCount     int       `json:"smsCount"`
		Status       SMSStatus `json:"status"`
		To           string    `json:"to"`
	} `json:"results"`
}

type GetSMSLogsResponse struct {
	Results []struct {
		BulkID    string    `json:"bulkId"`
		MessageID string    `json:"messageId"`
		To        string    `json:"to"`
		From      string    `json:"from"`
		Text      string    `json:"text"`
		SentAt    string    `json:"sentAt"`
		DoneAt    string    `json:"doneAt"`
		SmsCount  int       `json:"smsCount"`
		MccMnc    string    `json:"mccMnc"`
		Price     SMSPrice  `json:"price"`
		Status    SMSStatus `json:"status"`
		Error     SMSError  `json:"error"`
	} `json:"results"`
}

type GetSMSLogsParams struct {
	From          string
	To            string
	BulkID        []string
	MessageID     []string
	GeneralStatus string `validate:"omitempty,oneof=ACCEPTED PENDING UNDELIVERABLE DELIVERED REJECTED EXPIRED"`
	SentSince     string
	SentUntil     string
	Limit         int `validate:"omitempty,min=1,max=1000"`
	MCC           string
	MNC           string
}

func (g *GetSMSLogsParams) Validate() error {
	return validate.Struct(g)
}

type SMSBinary struct {
	Hex        string `json:"hex" validate:"required"`
	DataCoding int    `json:"dataCoding"`
	EsmClass   int    `json:"esmClass"`
}

type IndiaDLT struct {
	ContentTemplateID string `json:"contentTemplateId"`
	PrincipalEntityID string `json:"principalEntityId" validate:"required"`
}

type SMSRegional struct {
	IndiaDLT `json:"indiaDlt"`
}

type BinarySMSMsg struct {
	From               string                 `json:"from"`
	Destinations       []SMSDestination       `json:"destinations" validate:"min=1"`
	Binary             *SMSBinary             `json:"binary"`
	IntermediateReport bool                   `json:"intermediateReport,omitempty"`
	NotifyURL          string                 `json:"notifyUrl,omitempty"`
	NotifyContentType  string                 `json:"notifyContentType,omitempty"`
	CallbackData       string                 `json:"callbackData,omitempty"`
	ValidityPeriod     int                    `json:"validityPeriod,omitempty"`
	SendAt             string                 `json:"sendAt,omitempty"`
	DeliveryTimeWindow *SMSDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	Regional           *SMSRegional           `json:"regional,omitempty"`
}

func (b *BinarySMSMsg) Validate() error {
	return validate.Struct(b)
}

func (b *BinarySMSMsg) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(b)
}

type SendBinarySMSRequest struct {
	BulkID                string         `json:"bulkId"`
	Messages              []BinarySMSMsg `json:"messages" validate:"required,min=1"`
	*SMSSendingSpeedLimit `json:"sendingSpeedLimit,omitempty"`
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
		To        string    `json:"to"`
		Status    SMSStatus `json:"status"`
		MessageID string    `json:"messageId"`
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
		MessageID string    `json:"messageId"`
		Status    SMSStatus `json:"status,omitempty"`
		To        string    `json:"to"`
	} `json:"messages"`
}

type PreviewSMSRequest struct {
	LanguageCode    string `json:"languageCode,omitempty" validate:"omitempty,oneof=TR ES PT AUTODETECT"`
	Text            string `json:"text" validate:"required"`
	Transliteration string `json:"transliteration,omitempty" validate:"omitempty,oneof=TURKISH GREEK CYRILLIC SERBIAN_CYRILLIC CENTRAL_EUROPEAN BALTIC NON_UNICODE"` //nolint:lll
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
			Language        SMSLanguage `json:"language"`
			Transliteration string      `json:"transliteration"`
		} `json:"configuration"`
		MessageCount int    `json:"messageCount"`
		TextPreview  string `json:"textPreview"`
	} `json:"previews"`
}

type GetInboundSMSParams struct {
	Limit int `validate:"omitempty,min=1,max=1000"`
}

func (g *GetInboundSMSParams) Validate() error {
	return validate.Struct(g)
}

type GetInboundSMSResponse struct {
	MessageCount        int `json:"messageCount"`
	PendingMessageCount int `json:"pendingMessageCount"`
	Results             []struct {
		CallbackData string   `json:"callbackData"`
		CleanText    string   `json:"cleanText"`
		From         string   `json:"from"`
		Keyword      string   `json:"keyword"`
		MessageID    string   `json:"messageId"`
		Price        SMSPrice `json:"price"`
		ReceivedAt   string   `json:"receivedAt"`
		SmsCount     int      `json:"smsCount"`
		Text         string   `json:"text"`
		To           string   `json:"to"`
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
	Status string `json:"status"`
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
