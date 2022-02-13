package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidImageMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance ImageMessage
	}{
		{name: "minimum input",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/whatsappimage.jpg"},
			},
		},
		{
			name: "complete input",
			instance: ImageMessage{
				MessageCommon: GenerateTestMessageCommon(),
				Content: ImageContent{
					MediaURL: "https://www.mypath.com/whatsapp.jpg",
					Caption:  "hello world",
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

func TestImageMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name    string
		content ImageContent
	}{
		{
			name:    "missing Content field",
			content: ImageContent{},
		},
		{
			name:    "missing Content MediaURL",
			content: ImageContent{Caption: "asd"},
		},
		{
			name:    "Content MediaURL too long",
			content: ImageContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "Content invalid MediaURL",
			content: ImageContent{MediaURL: "asd"},
		},
		{
			name: "Content Caption too long",
			content: ImageContent{
				MediaURL: "https://www.mypath.com/whatsapp.jpg",
				Caption:  strings.Repeat("a", 3001),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := ImageMessage{
				MessageCommon: msgCommon,
				Content:       tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
