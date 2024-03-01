package models

import "bytes"

type GetAvailableNumbersParams struct {
	Capabilities []string `validate:"omitempty,dive,oneof=SMS VOICE MMS WHATSAPP WHATSAPP_VOICE"`
	Country      string
	State        string
	NPA          int32
	Nxx          int32
	Limit        int `validate:"omitempty,min=1,max=1000"`
	Number       string
	Page         int `validate:"omitempty,min=1,max=1000"`
}

func (g *GetAvailableNumbersParams) Validate() error {
	return validate.Struct(g)
}

type GetPurchasedNumberParam struct {
	NumberKey string `validate:"required"`
}

func (g *GetPurchasedNumberParam) Validate() error {
	return validate.Struct(g)
}

type UpdatePurchasedNumberRequest struct {
	ApplicationID string `json:"applicationId" validate:"required"`
	EntityID      string `json:"entityId,omitempty"`
}

func (u *UpdatePurchasedNumberRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

func (u *UpdatePurchasedNumberRequest) Validate() error {
	return validate.Struct(u)
}

type CancelPurchasedNumberParam struct {
	NumberKey string `validate:"required"`
}

func (g *CancelPurchasedNumberParam) Validate() error {
	return validate.Struct(g)
}

type ListPurchasedNumbersParam struct {
	Limit  int `validate:"omitempty,min=1,max=1000"`
	Number string
	Page   int `validate:"omitempty,min=0,max=1000"`
}

func (g *ListPurchasedNumbersParam) Validate() error {
	return validate.Struct(g)
}

type PurchaseNumberRequest struct {
	NumberKey     string `json:"numberKey,omitempty"`
	Number        string `json:"number,omitempty"`
	ApplicationID string `json:"applicationId,omitempty"`
	EntityID      string `json:"entityId,omitempty"`
}

func (p *PurchaseNumberRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(p)
}
func (p *PurchaseNumberRequest) Validate() error {
	return validate.Struct(p)
}

type GetAvailableNumbersResponse struct {
	Numbers     []Number `json:"numbers"`
	NumberCount int32    `json:"numberCount"`
}

type ListPurchasedNumbersResponse struct {
	Numbers     []Number `json:"numbers"`
	NumberCount int32    `json:"numberCount"`
}

type Number struct {
	NumberKey              string          `json:"numberKey,omitempty"`
	Number                 string          `json:"number,omitempty"`
	Country                string          `json:"country,omitempty"`
	CountryName            string          `json:"countryName,omitempty"`
	Type                   string          `json:"type,omitempty"`
	Capabilities           []string        `json:"capabilities,omitempty"`
	Shared                 bool            `json:"shared,omitempty"`
	Price                  *NumberPrice    `json:"price,omitempty"`
	Network                string          `json:"network,omitempty"`
	Keywords               []string        `json:"keywords,omitempty"`
	VoiceSetup             *VoiceSetup     `json:"voiceSetup,omitempty"`
	ReservationStatus      string          `json:"reservationStatus,omitempty"`
	AdditionalSetupRequest bool            `json:"additionalSetupRequest,omitempty"`
	EditPermissions        EditPermissions `json:"edit_permissions,omitempty"`
	Note                   string          `json:"note,omitempty"`
	ApplicationID          string          `json:"applicationId,omitempty"`
	EntityID               string          `json:"entityId,omitempty"`
}

type NumberPrice struct {
	PricePerMonth     float64 `json:"pricePerMonth"`
	SetupPrice        float64 `json:"setupPrice"`
	InitialMonthPrice float64 `json:"initialMonthPrice"`
	Cuurency          string  `json:"cuurency,omitempty"`
}

type VoiceSetup struct {
	ApplicationID string `json:"applicationId,omitempty"`
	EntityID      string `json:"entityId,omitempty"`
	Action        Action `json:"action"`
}

type Action struct {
	Description                 string `json:"description,omitempty"`
	Type                        string `json:"type,omitempty"`
	VoiceNumberMaskingConfigKey string `json:"voiceNumberMaskingConfigKey"`
}

type EditPermissions struct {
	CanEditNumber        bool `json:"canEditNumber,omitempty"`
	CanEditConfiguration bool `json:"canEditConfiguration,omitempty"`
}

type GetAllNumberConfigurationParam struct {
	Limit int
	Page  int
}

type GetAllNumberConfigurationResponse struct {
	Configurations []NumberConfiguration `json:"configurations,omitempty"`
	TotalCount     int32                 `json:"totalCount,omitempty"`
}

type NumberConfiguration struct {
	Keywork         string               `json:"keywork,omitempty"`
	Action          *ActionConfiguration `json:"action" validate:"required,dive"`
	UseConversation struct {
		Enable bool `json:"enable,omitempty"`
	} `json:"useConversation"`

	// Only for GET requests.
	OtherActions  []string `json:"otherActions,omitempty"`
	ApplicationID string   `json:"applicationId,omitempty"`
	EntityID      string   `json:"entityId,omitempty"`
	// Only Present in response
	Key string `json:"key,omitempty"`
}

func (n *NumberConfiguration) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(n)
}
func (n *NumberConfiguration) Validate() error {
	return validate.Struct(n)
}

type UpdateNumberConfigurationRequest struct {
	Keywork         string               `json:"keywork,omitempty"`
	Action          *ActionConfiguration `json:"action" validate:"required,dive"`
	UseConversation struct {
		Enable bool `json:"enable,omitempty"`
	} `json:"useConversation"`
	ApplicationID string `json:"applicationId,omitempty"`
	EntityID      string `json:"entityId,omitempty"`
	Key           string `json:"key" validate:"required"`
}

func (n *UpdateNumberConfigurationRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(n)
}
func (n *UpdateNumberConfigurationRequest) Validate() error {
	return validate.Struct(n)
}

type ActionConfiguration struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty" validate:"omitempty,oneof=PULL HTTP_FORWARD SMPP_FORWARD MAIL_FORWARD NO_ACTION"` //nolint:lll
	URL         string `json:"url,omitempty" validate:"required_if=Type HTTP_FORWARD,omitempty,url"`
	HTTPMethod  string `json:"httpMethod,omitempty" validate:"omitempty,oneof=GET POST"`
	ContentType string `json:"contentType,omitempty" validate:"omitempty,oneof=JSON XML"`
	Mail        string `json:"mail,omitempty" validate:"required_if=Type MAIL_FORWARD,omitempty,email"`
}

type OtherActionsDetails struct {
	// Only for GET request
	Editable           bool                      `json:"editable,omitempty"`
	Type               string                    `json:"type,omitempty"`
	Message            string                    `json:"message,omitempty"`
	Sender             string                    `json:"sender,omitempty"`
	DeliveryTimeWindow *NumberDeliveryTimeWindow `json:"deliveryTimeWindow,omitempty" validate:"omitempty,dive"`
}

type NumberDeliveryTimeWindow struct {
	Days             []string `json:"days,omitempty" validate:"required,gte=1,dive,oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"` //nolint:lll
	From             string   `json:"from" validate:"required"`
	To               string   `json:"to" validate:"required"`
	DeliveryTimeZone string   `json:"deliveryTimeZone,omitempty"`
}
