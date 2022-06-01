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
		instance WAAudioMsg
	}{
		{
			name: "minimum input",
			instance: WAAudioMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			}},
		{
			name: "complete input",
			instance: WAAudioMsg{
				MsgCommon: GenerateTestMsgCommon(),
				Content: AudioContent{
					MediaURL: "https://www.mypath.com/audio.mp3",
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

func TestAudioMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content AudioContent
	}{
		{
			name:    "missing MediaURL",
			content: AudioContent{},
		},
		{
			name:    "invalid MediaURL length",
			content: AudioContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
		},
		{
			name:    "invalid MediaURL format",
			content: AudioContent{MediaURL: "asd"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := WAAudioMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
