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
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
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
	tests := []struct {
		name     string
		instance VideoMessage
	}{
		{
			name: "missing From field",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "missing To field",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "missing Content field",
			instance: VideoMessage{
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
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "To too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "MessageID too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "CallbackData too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "NotifyURL too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: VideoContent{MediaURL: "https://www.mypath.com/whatsapp.jpg"},
			},
		},
		{
			name: "missing Content MediaURL",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{Caption: "asd"},
			},
		},
		{
			name: "Content MediaURL too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
			},
		},
		{
			name: "Content invalid MediaURL",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{MediaURL: "asd"},
			},
		},
		{
			name: "Content Caption too long",
			instance: VideoMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: VideoContent{
					MediaURL: "https://www.mypath.com/whatsapp.jpg",
					Caption:  strings.Repeat("a", 3001),
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
