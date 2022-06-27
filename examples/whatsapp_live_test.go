package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	apiKey  = "your-api-key"
	baseURL = "your-base-url"
	sender  = "1234567891011"
)

// The following examples can also be used to test the client against a real environment.
// Replace the apiKey and baseURL params along with the From, To and Content fields of the message, then run the test.
func TestTemplateMessagesExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WATemplateMsgs{
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

	resp, respDetails, err := client.WhatsApp.SendTemplate(context.Background(), message)
	fmt.Printf("%+v\n", resp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.BulkWAMsgResponse{}, resp)
}

func TestSendTextExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.TextContent{Text: "This message was sent from the Infobip API using the Go API client."},
	}

	msgResp, respDetails, err := client.WhatsApp.SendText(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestSendDocumentExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WADocumentMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.DocumentContent{MediaURL: "https://myurl.com/doc1.doc"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendDocument(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestSendImageExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAImageMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ImageContent{
			MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg",
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendImage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestAudioExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAAudioMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.AudioContent{MediaURL: "https://dl.espressif.com/dl/audio/ff-16b-2c-44100hz.aac"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendAudio(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestVideoExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAVideoMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.VideoContent{MediaURL: "https://download.samplelib.com/mp4/sample-5s.mp4"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendVideo(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestStickerExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAStickerMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.StickerContent{MediaURL: "https://myurl.com/sticker.webp"},
	}

	msgResp, respDetails, err := client.WhatsApp.SendSticker(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestLocationExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WALocationMsg{
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

	msgResp, respDetails, err := client.WhatsApp.SendLocation(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestContactExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAContactMsg{
		MsgCommon: models.MsgCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ContactContent{
			Contacts: []models.Contact{{Name: models.ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
		},
	}

	msgResp, respDetails, err := client.WhatsApp.SendContact(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestInteractiveButtonsExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAInteractiveButtonsMsg{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveButtons(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestInteractiveListExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAInteractiveListMsg{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveList(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestInteractiveProductExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAInteractiveProductMsg{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveProduct(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestInteractiveMultiproductExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	message := models.WAInteractiveMultiproductMsg{
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

	msgResp, respDetails, err := client.WhatsApp.SendInteractiveMultiproduct(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestGetTemplatesExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	msgResp, respDetails, err := client.WhatsApp.GetTemplates(context.Background(), sender)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetWATemplatesResponse{}, respDetails)
	assert.NotEqual(t, models.SendWAMsgResponse{}, msgResp)
}

func TestCreateTemplateExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	template := models.TemplateCreate{
		Name:     "template_name_my_test",
		Language: "en",
		Category: "ACCOUNT_UPDATE",
		Structure: models.TemplateStructure{
			Body: &models.TemplateStructureBody{Text: "body {{1}} content"},
			Type: "TEXT",
		},
	}

	msgResp, respDetails, err := client.WhatsApp.CreateTemplate(context.Background(), sender, template)
	fmt.Printf("%+v\n", msgResp)

	require.NoError(t, err)
	assert.NotEqual(t, models.CreateWATemplateResponse{}, respDetails)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
}

func TestDeleteTemplateExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)
	respDetails, err := client.WhatsApp.DeleteTemplate(context.Background(), sender, "template_name_my_test")

	require.NoError(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, http.StatusNoContent, respDetails.HTTPResponse.StatusCode)
}
