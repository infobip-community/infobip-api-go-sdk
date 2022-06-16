package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMMSExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.MMSMsg{
		Head: models.MMSHead{
			From: "111111111111",
			To:   "222222222222",
		},
		Text: "Some text",
	}

	msgResp, respDetails, err := client.MMS.Send(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendMMSResponse{}, msgResp)
}

func TestGetOutboundMsgDeliveryReportsExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	params := models.GetMMSDeliveryReportsParams{
		BulkID:    "1",
		MessageID: "1",
		Limit:     5,
	}

	msgResp, respDetails, err := client.MMS.GetDeliveryReports(context.Background(), params)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
}

func TestGetInboundMMSExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	params := models.GetInboundMMSParams{
		Limit: 5,
	}

	msgResp, respDetails, err := client.MMS.GetInboundMessages(context.Background(), params)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
}
