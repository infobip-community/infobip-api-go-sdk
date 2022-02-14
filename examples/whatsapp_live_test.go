package examples

import (
	"context"
	"fmt"
	"infobip-go-client/pkg/infobip"
	"infobip-go-client/pkg/infobip/models"
	"infobip-go-client/pkg/infobip/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The following examples can also be used to test the client against a real environment.
// Replace the apiKey and baseURL params along with the From, To and Content fields of the message, then run the test.
func TestSendTextExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.TextMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.TextContent{Text: "This message was sent from the Infobip API using the Go API client."},
	}

	msgResp, respDetails, err := whatsApp.SendTextMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
	fmt.Printf("%+v", msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
}

func TestSendDocumentExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.DocumentMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.DocumentContent{MediaURL: "https://myurl.com/doc1.doc"},
	}

	msgResp, respDetails, err := whatsApp.SendDocumentMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestSendImageExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.ImageMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ImageContent{
			MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg",
		},
	}

	msgResp, respDetails, err := whatsApp.SendImageMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestAudioExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.AudioMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.AudioContent{MediaURL: "https://dl.espressif.com/dl/audio/ff-16b-2c-44100hz.aac"},
	}

	msgResp, respDetails, err := whatsApp.SendAudioMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestVideoExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.VideoMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.VideoContent{MediaURL: "https://download.samplelib.com/mp4/sample-5s.mp4"},
	}

	msgResp, respDetails, err := whatsApp.SendVideoMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestStickerExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.StickerMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.StickerContent{MediaURL: "https://myurl.com/sticker.webp"},
	}

	msgResp, respDetails, err := whatsApp.SendStickerMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestLocationExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.LocationMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.LocationContent{
			Address:   "Some Address",
			Name:      "Something",
			Latitude:  utils.Float32Ptr(73.5164),
			Longitude: utils.Float32Ptr(56.2502),
		},
	}

	msgResp, respDetails, err := whatsApp.SendLocationMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestContactExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.ContactMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ContactContent{
			Contacts: []models.Contact{{Name: models.ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
		},
	}

	msgResp, respDetails, err := whatsApp.SendContactMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestInteractiveButtonsExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.InteractiveButtonsMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.InteractiveButtonsContent{
			Body: models.InteractiveButtonsBody{Text: "Some text"},
			Action: models.InteractiveButtons{
				Buttons: []models.InteractiveButton{
					{Type: "REPLY", ID: "1", Title: "Yes"},
					{Type: "REPLY", ID: "2", Title: "No"},
				},
			},
			Header: &models.InteractiveButtonsHeader{
				Type:     "IMAGE",
				MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg",
			},
		},
	}

	msgResp, respDetails, err := whatsApp.SendInteractiveButtonsMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestInteractiveListExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.Nil(t, err)
	whatsApp := client.WhatsApp()
	message := models.InteractiveListMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.InteractiveListContent{
			Body: models.InteractiveListBody{Text: "Some text"},
			Action: models.InteractiveListAction{
				Title: "some title",
				Sections: []models.Section{
					{Title: "Title 1", Rows: []models.SectionRow{{ID: "1", Title: "Row1 Title", Description: "desc"}}},
					{Title: "Title 2", Rows: []models.SectionRow{{ID: "2", Title: "Row2 Title", Description: "desc"}}},
				},
			},
			Header: &models.InteractiveListHeader{
				Type: "TEXT",
				Text: "Header text",
			},
			Footer: &models.InteractiveListFooter{Text: "Footer text"},
		},
	}

	msgResp, respDetails, err := whatsApp.SendInteractiveListMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}
