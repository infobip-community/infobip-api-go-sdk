package models

import (
	"fmt"
	"infobip-go-client/pkg/infobip/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidLocationMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance LocationMessage
	}{
		{
			name: "minimum input",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "complete input",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{
					Name:      "Some Place",
					Address:   "My Address",
					Latitude:  utils.Float32Ptr(10.5),
					Longitude: utils.Float32Ptr(10.5),
				},
			},
		},
		{
			name: "Latitude and longitude 0",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(0), Longitude: utils.Float32Ptr(0)},
			},
		},
		{
			name: "Latitude and longitude edge values",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(90), Longitude: utils.Float32Ptr(180)},
			},
		},
		{
			name: "Latitude and longitude edge values negative",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(-90), Longitude: utils.Float32Ptr(-180)},
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

func TestTextLocationConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance LocationMessage
	}{
		{
			name: "missing From field",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "missing To field",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "missing Content field",
			instance: LocationMessage{
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
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "To too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "MessageID too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "CallbackData too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "NotifyURL text too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "missing Content Latitude",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "invalid Content Latitude",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(91), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "missing Content Longitude",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "invalid Content Longitude",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(181)},
			},
		},
		{
			name: "Content Name too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Name: strings.Repeat("a", 1001)},
			},
		},
		{
			name: "Content Address too long",
			instance: LocationMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: LocationContent{Address: strings.Repeat("a", 1001)},
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
