package examples

import (
	"context"
	"fmt"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendMMSExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	message := models.MMSMsg{
		Head: models.MMSHead{
			From: "111111111111",
			To:   "222222222222",
		},
		Text: "Some text",
	}

	msgResp, respDetails, err := client.MMS.SendMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MMSResponse{}, msgResp)
}
