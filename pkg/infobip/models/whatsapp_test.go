package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextMessageRequestConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance TextMessageRequest
	}{
		{
			name: "missing From field",
			instance: TextMessageRequest{
				From:         "",
				To:           "+16175551213",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "missing To field",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "missing Content field",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "missing Content text",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{PreviewURL: false},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "From too long",
			instance: TextMessageRequest{
				From:         "+161755512133333333333333",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "To too long",
			instance: TextMessageRequest{
				From:         "+16175551212",
				To:           "+161755512133333333333333",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "MessageID too long",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    strings.Repeat("a", 51),
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "CallbackData too long",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: strings.Repeat("a", 4001),
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "Content text too long",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: strings.Repeat("a", 4097)},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "NotifyURL text too long",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
			},
		},
		{
			name: "NotifyURL not an url",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world"},
				CallbackData: "some data",
				NotifyURL:    "if only this was an url...",
			},
		},
		{
			name: "PreviewURL is true but text doesn't contain an url",
			instance: TextMessageRequest{
				From:         "+16175551213",
				To:           "+16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				Content:      Content{Text: "hello world", PreviewURL: true},
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			assert.NotNil(t, err)
		})
	}
}
