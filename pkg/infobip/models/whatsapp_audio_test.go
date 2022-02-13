package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidAudioMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance AudioMessage
	}{
		{name: "minimum input",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			}},
		{
			name: "complete input",
			instance: AudioMessage{
				MessageCommon: GenerateTestMessageCommon(),
				Content: AudioContent{
					MediaURL: "https://www.mypath.com/audio.mp3",
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

func TestAudioMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name    string
		content AudioContent
	}{
		{
			name:    "missing Content MediaURL",
			content: AudioContent{},
		},
		{
			name:    "Content MediaURL too long",
			content: AudioContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "Content invalid MediaURL",
			content: AudioContent{MediaURL: "asd"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := AudioMessage{
				MessageCommon: msgCommon,
				Content:       tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
