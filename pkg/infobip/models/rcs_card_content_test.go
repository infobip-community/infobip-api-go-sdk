package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidRCSCardContent(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSCardContent
	}{
		{
			name:     "minimum input",
			instance: RCSCardContent{},
		},
		{
			name: "full input",
			instance: RCSCardContent{
				Title:       "some-title",
				Description: "some-description",
				Media: &RCSCardContentMedia{
					File: &RCSFile{
						URL: "https://some-url",
					},
					Thumbnail: &RCSThumbnail{
						URL: "https://some-url",
					},
					Height: "MEDIUM",
				},
				Suggestions: []RCSSuggestion{
					GenerateReplyRCSSuggestion(),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled RCSCardContent
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidRCSCardContent(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSCardContent
	}{
		{
			name: "empty media",
			instance: RCSCardContent{
				Media: &RCSCardContentMedia{},
			},
		},
		{
			name: "empty media file",
			instance: RCSCardContent{
				Media: &RCSCardContentMedia{
					File: &RCSFile{},
				},
			},
		},
		{
			name: "empty media height",
			instance: RCSCardContent{
				Media: &RCSCardContentMedia{
					File: &RCSFile{
						URL: "https://some-url",
					},
				},
			},
		},
		{
			name: "invalid media height",
			instance: RCSCardContent{
				Media: &RCSCardContentMedia{
					File: &RCSFile{
						URL: "https://some-url",
					},
					Height: "INVALID",
				},
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
