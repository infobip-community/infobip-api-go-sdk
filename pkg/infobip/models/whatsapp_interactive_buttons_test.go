package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidInteractiveButtonsMessage(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		instance InteractiveButtonsMsg
	}{
		{
			name: "minimum input, no header",
			instance: InteractiveButtonsMsg{
				MsgCommon: MsgCommon{
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
			instance: InteractiveButtonsMsg{
				MsgCommon: msgCommon,
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
			instance: InteractiveButtonsMsg{
				MsgCommon: msgCommon,
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
			instance: InteractiveButtonsMsg{
				MsgCommon: msgCommon,
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
			instance: InteractiveButtonsMsg{
				MsgCommon: msgCommon,
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
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content InteractiveButtonsContent
	}{
		{
			name:    "empty Content field",
			content: InteractiveButtonsContent{},
		},
		{
			name: "missing Body",
			content: InteractiveButtonsContent{
				Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
			},
		},
		{
			name: "missing Body Text",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
			},
		},
		{
			name: "invalid Body Text length",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{Text: strings.Repeat("a", 1025)},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", ID: "1", Title: "Yes"}}},
			},
		},
		{
			name: "missing Action Buttons",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{},
			},
		},
		{
			name: "Action Buttons longer than 3",
			content: InteractiveButtonsContent{
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
		{
			name: "missing Action Button type",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "3", Title: "Yes"}}},
			},
		},
		{
			name: "invalid Action Button type",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "invalid", ID: "3", Title: "Yes"}}},
			},
		},
		{
			name: "missing Action Button ID",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{Type: "REPLY", Title: "Yes"}}},
			},
		},
		{
			name: "invalid Action Button ID length",
			content: InteractiveButtonsContent{
				Body: InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{
					Buttons: []InteractiveButton{
						{ID: strings.Repeat("a", 257), Type: "REPLY", Title: "Yes"},
					},
				},
			},
		},
		{
			name: "missing Action Button Title",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY"}}},
			},
		},
		{
			name: "invalid Action Button Title length",
			content: InteractiveButtonsContent{
				Body: InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{
					Buttons: []InteractiveButton{
						{ID: "1", Type: "REPLY", Title: strings.Repeat("a", 21)},
					},
				},
			},
		},
		{
			name: "missing Header Type",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{},
			},
		},
		{
			name: "invalid Header Type",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "invalid"},
			},
		},
		{
			name: "missing Text for Header Type TEXT",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "TEXT"},
			},
		},
		{
			name: "invalid Text for Header Type TEXT",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "TEXT", Text: strings.Repeat("a", 61)},
			},
		},
		{
			name: "missing MediaURL for Header Type VIDEO",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "VIDEO"},
			},
		},
		{
			name: "invalid MediaURL format for Header Type VIDEO",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "VIDEO", MediaURL: "asd"},
			},
		},
		{
			name: "invalid MediaURL length for Header Type VIDEO",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{
					Type:     "VIDEO",
					MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
				},
			},
		},
		{
			name: "missing MediaURL for Header Type IMAGE",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "IMAGE"},
			},
		},
		{
			name: "invalid MediaURL length for Header Type IMAGE",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{
					Type:     "IMAGE",
					MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
				},
			},
		},
		{
			name: "missing MediaURL for Header Type DOCUMENT",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{Type: "DOCUMENT"},
			},
		},
		{
			name: "invalid MediaURL length for Header Type DOCUMENT",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{
					Type:     "DOCUMENT",
					MediaURL: fmt.Sprintf("https://ww.g%sgle.com", strings.Repeat("o", 2048)),
				},
			},
		},
		{
			name: "invalid Filename length for Header Type DOCUMENT",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Header: &InteractiveButtonsHeader{
					Type:     "DOCUMENT",
					MediaURL: "https://ww.google.com",
					Filename: strings.Repeat("a", 241),
				},
			},
		},
		{
			name: "missing Text field for Footer",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Footer: &InteractiveButtonsFooter{},
			},
		},
		{
			name: "invalid Text length for Footer",
			content: InteractiveButtonsContent{
				Body:   InteractiveButtonsBody{"Some text"},
				Action: InteractiveButtons{Buttons: []InteractiveButton{{ID: "1", Type: "REPLY", Title: "YES"}}},
				Footer: &InteractiveButtonsFooter{Text: strings.Repeat("a", 61)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := InteractiveButtonsMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
