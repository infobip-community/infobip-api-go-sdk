package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The following examples can also be used to test the client against a real environment.
// Replace the apiKey and baseURL params along with the From, To and Content fields of the message, then run the test.
func TestTemplateMessagesExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.TemplateMsgs{
		Messages: []models.TemplateMsg{
			{
				MsgCommon: models.MsgCommon{From: "111111111111 ", To: "222222222222"},
				Content: models.TemplateMsgContent{
					TemplateName: "template_name",
					TemplateData: models.TemplateData{
						Body: models.TemplateBody{Placeholders: []string{}},
					},
					Language: "en_GB",
				},
			},
			{
				MsgCommon: models.MsgCommon{From: "111111111111 ", To: "222222222222"},
				Content: models.TemplateMsgContent{
					TemplateName: "template_name",
					TemplateData: models.TemplateData{
						Body: models.TemplateBody{Placeholders: []string{}},
					},
					Language: "en_GB",
				},
			},
		},
	}

	resp, respDetails, err := client.WhatsApp.SendTemplateMsgs(context.Background(), message)
	fmt.Printf("%+v\n", resp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.BulkMsgResponse{}, resp)
}

func TestSendTextExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.TextMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.TextContent{Text: "This message was sent from the Infobip API using the Go API client."},
	}

	msgResp, respDetails, err := client.WhatsApp.SendTextMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestSendDocumentExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.DocumentMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.DocumentContent{MediaURL: "https://myurl.com/doc1.doc"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendDocumentMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestSendImageExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.ImageMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ImageContent{
			MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg",
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendImageMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestAudioExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.AudioMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.AudioContent{MediaURL: "https://dl.espressif.com/dl/audio/ff-16b-2c-44100hz.aac"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendAudioMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestVideoExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.VideoMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.VideoContent{MediaURL: "https://download.samplelib.com/mp4/sample-5s.mp4"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendVideoMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestStickerExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.StickerMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.StickerContent{MediaURL: "https://myurl.com/sticker.webp"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendStickerMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestLocationExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.LocationMsg{
		MsgCommon: models.MsgCommon{
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

	msgResp, respDetails, err := client.WhatsApp.SendLocationMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestContactExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.ContactMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ContactContent{
			Contacts: []models.Contact{{Name: models.ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendContactMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestInteractiveButtonsExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.InteractiveButtonsMsg{
		MsgCommon: models.MsgCommon{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveButtonsMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestInteractiveListExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.InteractiveListMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.InteractiveListContent{
			Body: models.InteractiveListBody{Text: "Some text"},
			Action: models.InteractiveListAction{
				Title: "some title",
				Sections: []models.InteractiveListSection{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveListMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestInteractiveProductExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.InteractiveProductMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.InteractiveProductContent{
			Action: models.InteractiveProductAction{
				CatalogID:         "1",
				ProductRetailerID: "2",
			},
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveProductMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestInteractiveMultiproductExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.InteractiveMultiproductMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.InteractiveMultiproductContent{
			Header: models.InteractiveMultiproductHeader{Type: "TEXT", Text: "Header"},
			Body:   models.InteractiveMultiproductBody{Text: "Some Text"},
			Action: models.InteractiveMultiproductAction{
				CatalogID: "1",
				Sections: []models.InteractiveMultiproductSection{
					{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
				},
			},
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveMultiproductMsg(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestGetTemplatesExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	sender := "111111111111"

	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	msgResp, respDetails, err := client.WhatsApp.GetTemplates(context.Background(), sender)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.TemplatesResponse{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}

func TestCreateTemplateExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	sender := "111111111111"

	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	template := models.TemplateCreate{
		Name:     "template_name_mytest",
		Language: "en",
		Category: "ACCOUNT_UPDATE",
		Structure: models.TemplateStructure{
			Body: "body {{1}} content",
			Type: "TEXT",
		},
	}

	msgResp, respDetails, err := client.WhatsApp.CreateTemplate(context.Background(), sender, template)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.TemplateResponse{}, respDetails)
	assert.NotEqual(t, models.MsgResponse{}, msgResp)
}
