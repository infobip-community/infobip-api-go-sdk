package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	destNumber = "1234567891011"
	baseURL    = "your-base-url"
	apiKey     = "your-api-key"
)

func TestSendRCS(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	message := models.RCSMsg{
		From: "Gopher SDK",
		To:   destNumber,
		Content: &models.RCSContent{
			Type: "TEXT",
			Text: "some-text",
		},
	}

	resp, respDetails, err := client.RCS.Send(context.Background(), message)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp.Messages[0].MessageID, "Message ID should not be empty")
	assert.NotEqual(t, models.SendRCSResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestSendRCSBulk(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)

	req := models.SendRCSBulkRequest{
		Messages: []models.RCSMsg{
			{
				From: "Gopher SDK",
				To:   destNumber,
				Content: &models.RCSContent{
					Type: "TEXT",
					Text: "some-text",
				},
			},
		},
	}

	resp, respDetails, err := client.RCS.SendBulk(context.Background(), req)

	fmt.Println(resp)
	fmt.Println(respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEmptyf(t, resp[0].Messages[0].MessageID, "Message ID should not be empty")
	assert.NotEqual(t, models.SendRCSBulkResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
