package models

import "bytes"

type GetAvailableNumbersParams struct {
	Capabilities []string
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
