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
		instance VideoMsg
	}{
		{
			name: "minimum input",
			instance: VideoMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsappvideo.mp4"},
			}},
		{
			name: "complete input",
			instance: VideoMsg{
				MsgCommon: GenerateTestMsgCommon(),
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
			require.NoError(t, err)
		})
	}
}

func TestVideoMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content VideoContent
	}{
		{
			name:    "missing Content field",
			content: VideoContent{},
		},
		{
			name:    "missing MediaURL",
			content: VideoContent{Caption: "asd"},
		},
		{
			name:    "invalid MediaURL length",
			content: VideoContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "invalid MediaURL format",
			content: VideoContent{MediaURL: "asd"},
		},
		{
			name: "invalid Caption length",
			content: VideoContent{
				MediaURL: "https://www.mypath.com/whatsapp.jpg",
				Caption:  strings.Repeat("a", 3001),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := VideoMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
