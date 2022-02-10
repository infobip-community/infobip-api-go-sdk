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
				Content: ImageContent{MediaURL: "https://www.mypath.com/whatsappdoc.txt"},
			}},
		{
			name: "complete input",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{
					MediaURL: "https://www.mypath.com/my_doc.txt",
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
	tests := []struct {
		name     string
		instance ImageMessage
	}{
		{
			name: "missing From field",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "missing To field",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "missing Content field",
			instance: ImageMessage{
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
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "To too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "MessageID too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "CallbackData too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "NotifyURL too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: ImageContent{MediaURL: "https://www.mypath.com/my_doc.txt"},
			},
		},
		{
			name: "missing Content MediaURL",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{Caption: "asd"},
			},
		},
		{
			name: "Content MediaURL too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: fmt.Sprintf("https://www.g%sgle.com", strings.Repeat("o", 2040))},
			},
		},
		{
			name: "Content invalid MediaURL",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{MediaURL: "asd"},
			},
		},
		{
			name: "Content Caption too long",
			instance: ImageMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ImageContent{
					MediaURL: "https://www.mypath.com/my_doc.txt",
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
