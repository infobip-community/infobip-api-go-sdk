package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidRCSMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSMsg
	}{
		{
			name: "minimum input",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type: "TEXT",
					Text: "some-text",
				},
			},
		},
		{
			name:     "full FILE	input",
			instance: GenerateRCSFileMsg(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled RCSMsg
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidRCSMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance RCSMsg
	}{
		{
			name:     "empty input",
			instance: RCSMsg{},
		},
		{
			name: "missing to",
			instance: RCSMsg{
				Content: &RCSContent{
					Type: "TEXT",
					Text: "some-text",
				},
			},
		},
		{
			name: "missing content",
			instance: RCSMsg{
				To: "123456789",
			},
		},
		{
			name: "missing content type",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Text: "some-text",
				},
			},
		},
		{
			name: "invalid content type",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type: "INVALID",
				},
			},
		},
		{
			name: "missing file",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type: "FILE",
				},
			},
		},
		{
			name: "missing file url",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type: "FILE",
					File: &RCSFile{},
				},
			},
		},
		{
			name: "card missing orientation",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type:      "CARD",
					Alignment: "LEFT",
					Content:   GenerateRCSCardContent(),
				},
			},
		},
		{
			name: "card missing alignment",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type:        "CARD",
					Orientation: "HORIZONTAL",
					Content:     GenerateRCSCardContent(),
				},
			},
		},
		{
			name: "card with invalid alignment",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type:        "CARD",
					Alignment:   "INVALID",
					Orientation: "HORIZONTAL",
					Content:     GenerateRCSCardContent(),
				},
			},
		},
		{
			name: "card with invalid orientation",
			instance: RCSMsg{
				To: "123456789",
				Content: &RCSContent{
					Type:        "CARD",
					Alignment:   "LEFT",
					Orientation: "INVALID",
					Content:     GenerateRCSCardContent(),
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
