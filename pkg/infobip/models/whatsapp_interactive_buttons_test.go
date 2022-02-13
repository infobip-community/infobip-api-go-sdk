package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidInteractiveButtonsMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance InteractiveButtonsMessage
	}{
		{
			name: "minimum input, no header",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: InteractiveButtonsContent{
					Body: InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{
						Buttons: []InteractiveButton{
							{Type: "REPLY", ID: "1", Title: "Yes"},
							{Type: "REPLY", ID: "2", Title: "No"},
						},
					},
				},
			},
		},
		{
			name: "complete input, text header",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
					Header: &InteractiveButtonsHeader{Type: "TEXT", Text: "Some header"},
					Footer: &InteractiveButtonsFooter{Text: "Footer"},
				},
			},
		},
		{
			name: "complete input, video header",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
					Header: &InteractiveButtonsHeader{Type: "VIDEO", MediaURL: "https://myurl.com/asd.mp4"},
					Footer: &InteractiveButtonsFooter{Text: "Footer"},
				},
			},
		},
		{
			name: "complete input, image header",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
					Header: &InteractiveButtonsHeader{Type: "IMAGE", MediaURL: "https://myurl.com/asd.jpg"},
					Footer: &InteractiveButtonsFooter{Text: "Footer"},
				},
			},
		},
		{
			name: "complete input, document header",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
					Header: &InteractiveButtonsHeader{
						Type:     "DOCUMENT",
						MediaURL: "https://myurl.com/asd.pdf",
						Filename: "asd.pdf",
					},
					Footer: &InteractiveButtonsFooter{Text: "Footer"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Nil(t, err)
		})
	}
}

func TestTextInteractiveButtonsConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance InteractiveButtonsMessage
	}{
		{
			name: "missing From field",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing To field",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing Content field",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
			},
		},
		{
			name: "From too long",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "To too long",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "MessageID too long",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "CallbackData too long",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "NotifyURL text too long",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: "Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing Content Body",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing Content Body Text",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "invalid Content Body Text",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{Text: strings.Repeat("a", 1025)},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing Content Action Buttons",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{},
				},
			},
		},
		{
			name: "Content Action Buttons longer than 3",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body: InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{
						Buttons: []InteractiveButton{
							{Type: "REPLY", ID: "1", Title: "Yes"},
							{Type: "REPLY", ID: "2", Title: "No"},
							{Type: "REPLY", ID: "3", Title: "Maybe"},
							{Type: "REPLY", ID: "4", Title: "Too long"},
						},
					},
				},
			},
		},
		{
			name: "missing Content Action Button type",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "3", Title: "Yes"}}},
				},
			},
		},
		{
			name: "invalid Content Action Button type",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "invalid", ID: "3", Title: "Yes"}}},
				},
			},
		},
		{
			name: "missing Content Action Button ID",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", Title: "Yes"}}},
				},
			},
		},
		{
			name: "invalid Content Action Button ID",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body: InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{
						Buttons: []InteractiveButton{
							{ID: strings.Repeat("a", 257), Type: "REPLY", Title: "Yes"},
						},
					},
				},
			},
		},
		{
			name: "missing Content Action Button Title",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY"}}},
				},
			},
		},
		{
			name: "invalid Content Action Button Title",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body: InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{
						Buttons: []InteractiveButton{
							{ID: "1", Type: "REPLY", Title: strings.Repeat("a", 21)},
						},
					},
				},
			},
		},
		{
			name: "missing Content Header Type",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{},
				},
			},
		},
		{
			name: "invalid Content Header Type",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "invalid"},
				},
			},
		},
		{
			name: "missing Text field for Content Header Type TEXT",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "TEXT"},
				},
			},
		},
		{
			name: "invalid Text field for Content Header Type TEXT",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "TEXT", Text: strings.Repeat("a", 61)},
				},
			},
		},
		{
			name: "missing MediaURL field for Content Header Type VIDEO",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "VIDEO"},
				},
			},
		},
		{
			name: "invalid MediaURL field for Content Header Type VIDEO",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "VIDEO", MediaURL: "asd"},
				},
			},
		},
		{
			name: "invalid MediaURL field for Content Header Type VIDEO",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{
						Type:     "VIDEO",
						MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
					},
				},
			},
		},
		{
			name: "missing MediaURL field for Content Header Type IMAGE",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "IMAGE"},
				},
			},
		},
		{
			name: "invalid MediaURL field for Content Header Type IMAGE",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{
						Type:     "IMAGE",
						MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
					},
				},
			},
		},
		{
			name: "missing MediaURL field for Content Header Type DOCUMENT",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{Type: "DOCUMENT"},
				},
			},
		},
		{
			name: "invalid MediaURL field for Content Header Type DOCUMENT",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{
						Type:     "DOCUMENT",
						MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
					},
				},
			},
		},
		{
			name: "invalid Filename field for Content Header Type DOCUMENT",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Header: &InteractiveButtonsHeader{
						Type:     "DOCUMENT",
						MediaURL: "https://ww.google.com",
						Filename: strings.Repeat("a", 241),
					},
				},
			},
		},
		{
			name: "missing Text field for Content Footer",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Footer: &InteractiveButtonsFooter{},
				},
			},
		},
		{
			name: "invalid Text field for Content Footer",
			instance: InteractiveButtonsMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: InteractiveButtonsContent{
					Body:   InteractiveButtonsBody{"Some text"},
					Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
					Footer: &InteractiveButtonsFooter{Text: strings.Repeat("a", 61)},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NotNil(t, err)
		})
	}
}
