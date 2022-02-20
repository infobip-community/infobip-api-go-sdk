package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageCommon(t *testing.T) {
	tests := []struct {
		name     string
		instance MsgCommon
	}{
		{
			name: "minimum input",
			instance: MsgCommon{
				From: "16175551213",
				To:   "16175551212",
			},
		},
		{
			name: "complete input",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.instance)
			require.Nil(t, err)
		})
	}
}

func TestMessageCommonConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance MsgCommon
	}{
		{
			name: "missing From field",
			instance: MsgCommon{
				From:         "",
				To:           "16175551213",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "missing To field",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "From too long",
			instance: MsgCommon{
				From:         "1617555121333333333333333",
				To:           "16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "To too long",
			instance: MsgCommon{
				From:         "16175551212",
				To:           "1617555121333333333333333",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "MessageID too long",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "16175551212",
				MessageID:    strings.Repeat("a", 51),
				CallbackData: "some data",
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "CallbackData too long",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: strings.Repeat("a", 4001),
				NotifyURL:    "https://www.google.com",
			},
		},
		{
			name: "NotifyURL text too long",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
			},
		},
		{
			name: "NotifyURL not an url",
			instance: MsgCommon{
				From:         "16175551213",
				To:           "16175551212",
				MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
				CallbackData: "some data",
				NotifyURL:    "if only this was an url...",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.instance)
			require.NotNil(t, err)
		})
	}
}
