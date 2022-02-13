package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidTextMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance TextMessage
	}{
		{name: "minimum input",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: TextContent{Text: "hello world"},
			}},
		{
			name: "complete input",
			instance: TextMessage{
				MessageCommon: GenerateTestMessageCommon(),
				Content: TextContent{
					Text:       "hello world, here's the link: https://www.google.com",
					PreviewURL: true,
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

func TestTextMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name    string
		content TextContent
	}{
		{
			name:    "missing Content field",
			content: TextContent{},
		},
		{
			name:    "missing Content text",
			content: TextContent{PreviewURL: false},
		},
		{
			name:    "Content text too long",
			content: TextContent{Text: strings.Repeat("a", 4097)},
		},
		{
			name:    "PreviewURL is true but text doesn't contain an url",
			content: TextContent{Text: "hello world", PreviewURL: true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := TextMessage{
				MessageCommon: msgCommon,
				Content:       tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
