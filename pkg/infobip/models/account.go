package models

import "bytes"

type AccountBalance struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type FreeMessagesCount struct {
	RemainingCount   int32 `json:"remainingCount"`
	SerialVersionUID int64 `json:"serialVersionUID"`
}

type TotalAccountBalance struct {
	Balance  float64 `json:"balance,omitempty"`
	Currency struct {
		Code         string `json:"code,omitempty"`
		CurrencyName string `json:"currencyMame,omitempty"`
		ID           int32  `json:"id,omitempty"`
		Symbol       string `json:"symbol,omitempty"`
	} `json:"currency,omitempty"`

	FreeMessages map[string]int32 `json:"freeMessages,omitempty"`
}

type GetAllAccountsParams struct {
	Name   string
	Limit  int32
	Enable *bool
}

type GetAllAccountsResponse struct {
	Accounts []struct {
		Enable   bool   `json:"enable,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		OwnerKey string `json:"ownerKey,omitempty"`
	} `json:"accounts,omitempty"`
}

type UpdateAccountRequest struct {
	Name   string `json:"name,omitempty"`
	Enable bool   `json:"enable,omitempty"`
}

func (u *UpdateAccountRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

func (u *UpdateAccountRequest) Validate() error {
	return validate.Struct(u)
}

type UpdateAccountResponse struct {
	Enable   bool   `json:"enable,omitempty"`
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	OwnerKey string `json:"ownerKey,omitempty"`
}

type GetAPIKeybyFilterParam struct {
	AccountID      string
	Name           string
	Page           int32
	Size           int32
	OrderBy        string
	OrderDirection string
	Enable         *bool
	APIKeystring   string
}

type GetAPIKeybyFilterResponse struct {
	APIKeys []APIKey `json:"apiKeys,omitempty"`
	Paging  struct {
		Page           int32  `json:"page,omitempty"`
		PageSize       int32  `json:"pageSize,omitempty"`
		OrderBy        string `json:"orderBy,omitempty"`
		OrderDirection string `json:"orderDirection,omitempty"`
		TotalCount     int64  `json:"totalCount,omitempty"`
		TotalPages     int32  `json:"totalPages,omitempty"`
	} `json:"paging,omitempty" validation:"required"`
}
type APIKey struct {
	ID           string     `json:"id,omitempty"`
	APIKeystring string     `json:"apiKeyString,omitempty"`
	AccountID    string     `json:"accountId,omitempty"`
	Name         string     `json:"name,omitempty" validate:"required"`
	AllowedIPs   []string   `json:"allowedIps,omitempty"`
	ValidFrom    string     `json:"validFrom,omitempty"`
	ValidTo      string     `json:"validTo,omitempty"`
	Enable       bool       `json:"enable,omitempty"`
	Permissions  []string   `json:"permissions,omitempty"`
	Platform     []Platform `json:"platform,omitempty" validate:"omitempty,dive"`
	ScopeGuide   []string   `json:"scopeGuide,omitempty"`
}

func (a *APIKey) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(a)
}

func (a *APIKey) Validate() error {
	return validate.Struct(a)
}

type Platform struct {
	Key           string `json:"key,omitempty"`
	ApplicationID string `json:"applicationId" validate:"required"`
	EntityID      string `json:"entity_id,omitempty"`
	Action        string `json:"action,omitempty" validate:"omitempty,oneof=FILL FORCE"`
}

type UpdateAPIKeyRequest struct {
	AccountID   string   `json:"accountId,omitempty"`
	Name        string   `json:"name,omitempty" validate:"required"`
	AllowedIPs  []string `json:"allowedIps,omitempty"`
	ValidFrom   string   `json:"validFrom,omitempty"`
	ValidTo     string   `json:"validTo,omitempty"`
	Enable      bool     `json:"enable,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Platform    []struct {
		Key           string `json:"key,omitempty"`
		ApplicationID string `json:"applicationId" validate:"required"`
		EntityID      string `json:"entity_id,omitempty"`
		Action        string `json:"action,omitempty" validate:"omitempty,oneof=FILL FORCE"`
	} `json:"platform,omitempty"`
	ScopeGuide []string `json:"scopeGuide,omitempty"`
}

func (u *UpdateAPIKeyRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

func (u *UpdateAPIKeyRequest) Validate() error {
	return validate.Struct(u)
}

type CreateSessionRequest struct {
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func (c *CreateSessionRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(c)
}

func (c *CreateSessionRequest) Validate() error {
	return validate.Struct(c)
}

type Token struct {
	Token string `json:"token"`
}

type CreateOauth2TokenRequest struct {
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	GrantType    string `json:"grant_type" validate:"required"`
}

func (c *CreateOauth2TokenRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(c)
}

func (c *CreateOauth2TokenRequest) Validate() error {
	return validate.Struct(c)
}

type CreateOauth2TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
}
