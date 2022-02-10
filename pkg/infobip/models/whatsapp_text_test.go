package models

import (
	"fmt"
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
					From: "+16175551213",
					To:   "+16175551212",
				},
				Content: TextContent{Text: "hello world"},
			}},
		{
			name: "complete input",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
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
	tests := []struct {
		name     string
		instance TextMessage
	}{
		{
			name: "missing From field",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "+16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "missing To field",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "missing Content field",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
			},
		},
		{
			name: "missing Content text",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{PreviewURL: false},
			},
		},
		{
			name: "From too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+161755512133333333333333",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "To too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551212",
					To:           "+161755512133333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "MessageID too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "CallbackData too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "Content text too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: strings.Repeat("a", 4097)},
			},
		},
		{
			name: "NotifyURL text too long",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: TextContent{Text: "hello world"},
			},
		},
		{
			name: "PreviewURL is true but text doesn't contain an url",
			instance: TextMessage{
				MessageCommon: MessageCommon{
					From:         "+16175551213",
					To:           "+16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: TextContent{Text: "hello world", PreviewURL: true},
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
