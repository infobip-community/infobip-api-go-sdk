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
		instance ImageMsg
	}{
		{name: "minimum input",
			instance: ImageMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/whatsappimage.jpg"},
			},
		},
		{
			name: "complete input",
			instance: ImageMsg{
				MsgCommon: GenerateTestMsgCommon(),
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
			require.NoError(t, err)
		})
	}
}

func TestImageMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content ImageContent
	}{
		{
			name:    "missing Content field",
			content: ImageContent{},
		},
		{
			name:    "missing MediaURL",
			content: ImageContent{Caption: "asd"},
		},
		{
			name:    "invalid MediaURL length",
			content: ImageContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "invalid MediaURL format",
			content: ImageContent{MediaURL: "asd"},
		},
		{
			name: "invalid Caption length",
			content: ImageContent{
				MediaURL: "https://www.mypath.com/whatsapp.jpg",
				Caption:  strings.Repeat("a", 3001),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := ImageMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
