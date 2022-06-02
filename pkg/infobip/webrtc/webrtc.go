package webrtc

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

const (
	getApplicationsPath   = "webrtc/1/applications"
	saveApplicationPath   = "webrtc/1/applications"
	getApplicationPath    = "webrtc/1/applications"
	updateApplicationPath = "webrtc/1/applications"
	generateTokenPath     = "webrtc/1/token"
)

type WebRTC interface {
	// GetApplications returns a list of all applications for WebRTC channel.
	GetApplications(
		ctx context.Context,
	) (resp models.GetWebRTCApplicationsResponse, respDetails models.ResponseDetails, err error)

	// SaveApplication creates and configures a new WebRTC application.
	SaveApplication(
		ctx context.Context,
		application models.WebRTCApplication,
	) (resp models.SaveWebRTCApplicationResponse, respDetails models.ResponseDetails, err error)

	// GetApplication returns a single WebRTC application to see its configuration details.
	GetApplication(
		ctx context.Context,
		applicationID string,
	) (resp models.GetWebRTCApplicationResponse, respDetails models.ResponseDetails, err error)

	// UpdateApplication changes configuration details of your existing WebRTC application.
	UpdateApplication(
		ctx context.Context,
		applicationID string,
		application models.WebRTCApplication,
	) (resp models.UpdateWebRTCApplicationResponse, respDetails models.ResponseDetails, err error)

	// DeleteApplication deletes WebRTC application for a given id.
	DeleteApplication(
		ctx context.Context,
		applicationID string,
	) (respDetails models.ResponseDetails, err error)

	// GetToken generates and returns token for WebRTC channel.
	GenerateToken(
		ctx context.Context,
		req models.GenerateWebRTCTokenRequest,
	) (resp models.GenerateWebRTCTokenResponse, respDetails models.ResponseDetails, err error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

func (wrtc *Channel) SaveApplication(
	ctx context.Context,
	application models.WebRTCApplication,
) (resp models.SaveWebRTCApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.PostJSONReq(ctx, &application, &resp, saveApplicationPath)
	return resp, respDetails, err
}

func (wrtc *Channel) GetApplications(
	ctx context.Context,
) (resp models.GetWebRTCApplicationsResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.GetRequest(ctx, &resp, getApplicationsPath, nil)
	return resp, respDetails, err
}

func (wrtc *Channel) GetApplication(
	ctx context.Context,
	applicationID string,
) (resp models.GetWebRTCApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.GetRequest(
		ctx, &resp, fmt.Sprint(getApplicationPath, "/", applicationID), nil)
	return resp, respDetails, err
}

func (wrtc *Channel) UpdateApplication(
	ctx context.Context,
	applicationID string,
	application models.WebRTCApplication,
) (resp models.UpdateWebRTCApplicationResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.PutJSONReq(
		ctx, &application, &resp, fmt.Sprint(updateApplicationPath, "/", applicationID), nil)
	return resp, respDetails, err
}

func (wrtc *Channel) DeleteApplication(
	ctx context.Context,
	applicationID string,
) (respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.DeleteRequest(ctx, fmt.Sprint(getApplicationPath, "/", applicationID), nil)
	return respDetails, err
}

func (wrtc *Channel) GenerateToken(
	ctx context.Context,
	req models.GenerateWebRTCTokenRequest,
) (resp models.GenerateWebRTCTokenResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = wrtc.ReqHandler.PostJSONReq(ctx, &req, &resp, generateTokenPath)
	return resp, respDetails, err
}
