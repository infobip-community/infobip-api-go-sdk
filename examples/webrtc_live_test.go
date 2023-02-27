package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetApplications(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	resp, respDetails, err := client.WebRTC.GetApplications(context.Background())

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp[0].Name, "Name should not be empty")
	assert.NotEqual(t, models.GetWebRTCApplicationsResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSaveApplication(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	application := models.WebRTCApplication{
		Name:        "some-name",
		Description: "some-description",
	}

	resp, respDetails, err := client.WebRTC.SaveApplication(context.Background(), application)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Name, "Name should not be empty")
	assert.NotEqual(t, models.GetWebRTCApplicationResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetApplication(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	id := "e672f02e-aed7-4898-9161-9e2503b6acc8"
	resp, respDetails, err := client.WebRTC.GetApplication(context.Background(), id)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Name, "Name should not be empty")
	assert.NotEqual(t, models.GetWebRTCApplicationResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdateApplication(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	application := models.WebRTCApplication{
		Name:        "some-name",
		Description: "some-description-updated",
	}

	id := "e672f02e-aed7-4898-9161-9e2503b6acc8"

	resp, respDetails, err := client.WebRTC.UpdateApplication(context.Background(), id, application)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Name, "Name should not be empty")
	assert.NotEqual(t, models.UpdateWebRTCApplicationResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestDeleteApplication(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	id := "e672f02e-aed7-4898-9161-9e2503b6acc8"

	respDetails, err := client.WebRTC.DeleteApplication(context.Background(), id)

	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGenerateToken(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.GenerateWebRTCTokenRequest{
		Identity:      "SomeIdentity",
		ApplicationID: "fb994013-4d1e-4e75-b45a-ab4025715cfe",
		DisplayName:   "SomeDisplayName",
		Capabilities: &models.WebRTCTokenCapabilities{
			Recording: "ALWAYS",
		},
		TimeToLive: 100,
	}

	resp, respDetails, err := client.WebRTC.GenerateToken(context.Background(), req)

	fmt.Println(resp)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Token, "Token should not be empty")
	assert.NotEqual(t, models.GenerateWebRTCTokenResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
