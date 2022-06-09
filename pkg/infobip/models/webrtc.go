package models

import (
	"bytes"
)

type WebRTCIOS struct {
	ApnsCertificateFileName    string `json:"apnsCertificateFileName" validate:"required"`
	ApnsCertificateFileContent string `json:"apnsCertificateFileContent" validate:"required"`
	ApnsCertificatePassword    string `json:"apnsCertificatePassword"`
}

type WebRTCAndroid struct {
	FcmServerKey string `json:"fcmServerKey" validate:"required"`
}

type GetWebRTCApplicationsResponse []WebRTCApplication

type WebRTCApplication struct {
	ID                 string         `json:"id"`
	Name               string         `json:"name" validate:"required"`
	Description        string         `json:"description"`
	IOS                *WebRTCIOS     `json:"ios,omitempty"`
	Android            *WebRTCAndroid `json:"android,omitempty"`
	AppToApp           bool           `json:"appToApp"`
	AppToConversations bool           `json:"appToConversations"`
	AppToPhone         bool           `json:"appToPhone"`
}

func (w *WebRTCApplication) Validate() error {
	return validate.Struct(w)
}

func (w *WebRTCApplication) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(w)
}

type SaveWebRTCApplicationResponse WebRTCApplication

type GetWebRTCApplicationResponse WebRTCApplication

type UpdateWebRTCApplicationResponse WebRTCApplication

type WebRTCTokenCapabilities struct {
	Recording string `json:"recording" validate:"omitempty,oneof=ALWAYS ON_DEMAND DISABLED"`
}

type GenerateWebRTCTokenRequest struct {
	Identity      string                   `json:"identity" validate:"required,min=3,max=64"`
	ApplicationID string                   `json:"applicationId"`
	DisplayName   string                   `json:"displayName"`
	Capabilities  *WebRTCTokenCapabilities `json:"capabilities"`
	TimeToLive    int64                    `json:"timeToLive"`
}

func (g *GenerateWebRTCTokenRequest) Validate() error {
	return validate.Struct(g)
}

func (g *GenerateWebRTCTokenRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(g)
}

type GenerateWebRTCTokenResponse struct {
	Token          string `json:"token"`
	ExpirationTime string `json:"expirationTime"`
}
