package models

import (
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/utils"

	"github.com/stretchr/testify/require"
)

func TestValidLocationMessage(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		instance LocationMsg
	}{
		{
			name: "minimum input",
			instance: LocationMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(10.5)},
			},
		},
		{
			name: "complete input",
			instance: LocationMsg{
				MsgCommon: msgCommon,
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
			instance: LocationMsg{
				MsgCommon: msgCommon,
				Content:   LocationContent{Latitude: utils.Float32Ptr(0), Longitude: utils.Float32Ptr(0)},
			},
		},
		{
			name: "Latitude and longitude edge values positive",
			instance: LocationMsg{
				MsgCommon: msgCommon,
				Content:   LocationContent{Latitude: utils.Float32Ptr(90), Longitude: utils.Float32Ptr(180)},
			},
		},
		{
			name: "Latitude and longitude edge values negative",
			instance: LocationMsg{
				MsgCommon: msgCommon,
				Content:   LocationContent{Latitude: utils.Float32Ptr(-90), Longitude: utils.Float32Ptr(-180)},
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
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content LocationContent
	}{
		{
			name:    "empty Content field",
			content: LocationContent{},
		},
		{
			name:    "missing Latitude",
			content: LocationContent{Longitude: utils.Float32Ptr(10.5)},
		},
		{
			name:    "invalid Latitude",
			content: LocationContent{Latitude: utils.Float32Ptr(91), Longitude: utils.Float32Ptr(10.5)},
		},
		{
			name:    "missing Longitude",
			content: LocationContent{Latitude: utils.Float32Ptr(10.5)},
		},
		{
			name:    "invalid Longitude",
			content: LocationContent{Latitude: utils.Float32Ptr(10.5), Longitude: utils.Float32Ptr(181)},
		},
		{
			name:    "invalid Name length",
			content: LocationContent{Name: strings.Repeat("a", 1001)},
		},
		{
			name:    "invalid Address length",
			content: LocationContent{Address: strings.Repeat("a", 1001)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := LocationMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
