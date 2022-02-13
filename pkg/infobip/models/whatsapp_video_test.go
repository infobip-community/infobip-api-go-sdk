package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidVideoMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance VideoMessage
	}{
		{name: "minimum input",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsappvideo.mp4"},
			}},
		{
			name: "complete input",
			instance: VideoMessage{
				MessageCommon: GenerateTestMessageCommon(),
				Content: VideoContent{
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

func TestVideoMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name    string
		content VideoContent
	}{
		{
			name:    "missing Content field",
			content: VideoContent{},
		},
		{
			name:    "missing Content MediaURL",
			content: VideoContent{Caption: "asd"},
		},
		{
			name:    "Content MediaURL too long",
			content: VideoContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "Content invalid MediaURL",
			content: VideoContent{MediaURL: "asd"},
		},
		{
			name: "Content Caption too long",
			content: VideoContent{
				MediaURL: "https://www.mypath.com/whatsapp.jpg",
				Caption:  strings.Repeat("a", 3001),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := VideoMessage{
				MessageCommon: msgCommon,
				Content:       tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
