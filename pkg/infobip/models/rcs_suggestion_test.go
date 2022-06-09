package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidRCSSuggestion(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSSuggestion
	}{
		{
			name: "minimum input",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "REPLY",
			},
		},
		{
			name:     "full REPLY input",
			instance: GenerateReplyRCSSuggestion(),
		},
		{
			name:     "full OPEN_URL input",
			instance: GenerateOpenURLRCSSuggestion(),
		},
		{
			name:     "full DIAL_PHONE input",
			instance: GenerateDialPhoneRCSSuggestion(),
		},
		{
			name:     "full SHOW_LOCATION input",
			instance: GenerateShowLocationRCSSuggestion(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled RCSSuggestion
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidRCSSuggestion(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSSuggestion
	}{
		{
			name:     "empty input",
			instance: RCSSuggestion{},
		},
		{
			name: "empty text",
			instance: RCSSuggestion{
				Text:         "",
				PostbackData: "some-postback-data",
				Type:         "REPLY",
			},
		},
		{
			name: "empty postback data",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "",
				Type:         "REPLY",
			},
		},
		{
			name: "empty type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "",
			},
		},
		{
			name: "invalid type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "INVALID",
			},
		},
		{
			name: "missing URL for OPEN_URL type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "OPEN_URL",
			},
		},
		{
			name: "missing latitude for SHOW_LOCATION type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "SHOW_LOCATION",
				Longitude:    20.0,
			},
		},
		{
			name: "missing longitude for SHOW_LOCATION type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "SHOW_LOCATION",
				Latitude:     20.0,
			},
		},
		{
			name: "latitude out of range for SHOW_LOCATION type",
			instance: RCSSuggestion{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "SHOW_LOCATION",
				Latitude:     -91.0,
				Longitude:    20.0,
			},
		},
		{
			name: "text too long",
			instance: RCSSuggestion{
				Text:         "some-text-that-is-longer-than-25-characters-so-it-should-be-invalid",
				PostbackData: "some-postback-data",
				Type:         "REPLY",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
