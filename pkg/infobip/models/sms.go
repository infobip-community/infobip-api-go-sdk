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
	Destinations       []SMSDestination       `json:"destinations" validate:"required,min=1,dive"`
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
	Messages          []SMSMsg              `json:"messages" validate:"required,min=1,dive"`
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
	Destinations       []SMSDestination       `json:"destinations" validate:"min=1,dive"`
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
	Messages              []BinarySMSMsg `json:"messages" validate:"required,min=1,dive"`
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

type TFAApplicationConfiguration struct {
	AllowMultiplePINVerifications bool   `json:"allowMultiplePinVerifications"`
	PINAttempts                   int    `json:"pinAttempts"`
	PINTimeToLive                 string `json:"pinTimeToLive"`
	SendPINPerApplicationLimit    string `json:"sendPinPerApplicationLimit"`
	SendPINPerPhoneNumberLimit    string `json:"sendPinPerPhoneNumberLimit"`
	VerifyPINLimit                string `json:"verifyPinLimit"`
}

type TFAApplication struct {
	ApplicationID string                       `json:"applicationId"`
	Configuration *TFAApplicationConfiguration `json:"configuration,omitempty"`
	Enabled       bool                         `json:"enabled"`
	Name          string                       `json:"name" validate:"required"`
}
type GetTFAApplicationsResponse []TFAApplication

type CreateTFAApplicationRequest TFAApplication

func (c *CreateTFAApplicationRequest) Validate() error {
	return validate.Struct(c)
}

func (c *CreateTFAApplicationRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(c)
}

type CreateTFAApplicationResponse TFAApplication

type GetTFAApplicationResponse TFAApplication

type UpdateTFAApplicationRequest TFAApplication

func (u *UpdateTFAApplicationRequest) Validate() error {
	return validate.Struct(u)
}

func (u *UpdateTFAApplicationRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

type UpdateTFAApplicationResponse TFAApplication

type TFAMessageTemplate struct {
	ApplicationID  string       `json:"applicationId,omitempty"`
	Language       string       `json:"language,omitempty"`
	MessageID      string       `json:"messageId,omitempty"`
	MessageText    string       `json:"messageText,omitempty" validate:"required"`
	PINLength      int          `json:"pinLength,omitempty"`
	PINPlaceholder string       `json:"pinPlaceholder,omitempty"`
	PINType        string       `json:"pinType,omitempty" validate:"required,oneof=NUMERIC ALPHANUMERIC ALPHA HEX"`
	Regional       *SMSRegional `json:"regional,omitempty"`
	RepeatDTMF     string       `json:"repeatDTMF,omitempty"`
	SenderID       string       `json:"senderId,omitempty"`
	SpeechRate     float64      `json:"speechRate,omitempty"`
}

type GetTFAMessageTemplatesResponse []TFAMessageTemplate

type CreateTFAMessageTemplateRequest TFAMessageTemplate

func (c *CreateTFAMessageTemplateRequest) Validate() error {
	return validate.Struct(c)
}

func (c *CreateTFAMessageTemplateRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(c)
}

type CreateTFAMessageTemplateResponse TFAMessageTemplate

type GetTFAMessageTemplateResponse TFAMessageTemplate

type UpdateTFAMessageTemplateRequest TFAMessageTemplate

func (u *UpdateTFAMessageTemplateRequest) Validate() error {
	return validate.Struct(u)
}

func (u *UpdateTFAMessageTemplateRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

type UpdateTFAMessageTemplateResponse TFAMessageTemplate

type PINPlaceholders struct {
	FirstName string `json:"firstName,omitempty"` // FIXME: Members can change.
}

type SendPINRequest struct {
	ApplicationID string           `json:"applicationId" validate:"required"`
	MessageID     string           `json:"messageId" validate:"required"`
	From          string           `json:"from,omitempty"`
	To            string           `json:"to" validate:"required"`
	Placeholders  *PINPlaceholders `json:"placeholders,omitempty"`
}

type SendPINOverSMSRequest SendPINRequest

func (s *SendPINOverSMSRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendPINOverSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SendPINOverSMSParams struct {
	NCNeeded bool
}

type SendPINResponse struct {
	PINID      string `json:"pinId"`
	To         string `json:"to"`
	NCStatus   string `json:"ncStatus,omitempty"`
	SMSStatus  string `json:"smsStatus,omitempty"`
	CallStatus string `json:"callStatus,omitempty"`
}

type SendPINOverSMSResponse SendPINResponse

type ResendPINRequest struct {
	Placeholders *PINPlaceholders `json:"placeholders,omitempty"`
}

type ResendPINOverSMSRequest ResendPINRequest

func (r *ResendPINOverSMSRequest) Validate() error {
	return validate.Struct(r)
}

func (r *ResendPINOverSMSRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type ResendPINOverSMSResponse SendPINResponse

type SendPINOverVoiceRequest SendPINRequest

func (s *SendPINOverVoiceRequest) Validate() error {
	return validate.Struct(s)
}

func (s *SendPINOverVoiceRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(s)
}

type SendPINOverVoiceResponse SendPINResponse

type ResendPINOverVoiceRequest ResendPINRequest

func (r *ResendPINOverVoiceRequest) Validate() error {
	return validate.Struct(r)
}

func (r *ResendPINOverVoiceRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type ResendPINOverVoiceResponse SendPINResponse

type VerifyPhoneNumberRequest struct {
	PIN string `json:"pin" validate:"required"`
}

func (v *VerifyPhoneNumberRequest) Validate() error {
	return validate.Struct(v)
}

func (v *VerifyPhoneNumberRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(v)
}

type VerifyPhoneNumberResponse struct {
	PINId             string `json:"pinId"`
	MSISDN            string `json:"msisdn"`
	Verified          bool   `json:"verified"`
	AttemptsRemaining int    `json:"attemptsRemaining"`
}

type GetTFAVerificationStatusParams struct {
	MSISDN   string `json:"msisdn" validate:"required"`
	Verified bool   `json:"verified"`
	Sent     bool   `json:"sent"`
}

type GetTFAVerificationStatusResponse struct {
	Verifications []struct {
		Msisdn     string `json:"msisdn"`
		Verified   bool   `json:"verified"`
		VerifiedAt int    `json:"verifiedAt"`
		SentAt     int    `json:"sentAt"`
	} `json:"verifications"`
}
