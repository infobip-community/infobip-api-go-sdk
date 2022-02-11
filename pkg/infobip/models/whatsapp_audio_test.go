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
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
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
	tests := []struct {
		name     string
		instance AudioMessage
	}{
		{
			name: "missing From field",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "missing To field",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "missing Content field",
			instance: AudioMessage{
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
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "To too long",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "MessageID too long",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "CallbackData too long",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "NotifyURL too long",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: AudioContent{MediaURL: "https://www.mypath.com/audio.mp3"},
			},
		},
		{
			name: "missing Content MediaURL",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{},
			},
		},
		{
			name: "Content MediaURL too long",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
			},
		},
		{
			name: "Content invalid MediaURL",
			instance: AudioMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: AudioContent{MediaURL: "asd"},
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
