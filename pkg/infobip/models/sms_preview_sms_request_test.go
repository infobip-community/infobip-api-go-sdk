package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPreviewSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance PreviewSMSRequest
	}{
		{
			name: "minimum input",
			instance: PreviewSMSRequest{
				Text: "some text",
			},
		},
		{
			name:     "full input",
			instance: GeneratePreviewSMSRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestInvalidPreviewSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance PreviewSMSRequest
	}{
		{
			name:     "empty input",
			instance: PreviewSMSRequest{},
		},
		{
			name: "bad transliteration",
			instance: PreviewSMSRequest{
				Text:            "some text",
				Transliteration: "invalid",
			},
		},
		{
			name: "bad language code",
			instance: PreviewSMSRequest{
				Text:         "some text",
				LanguageCode: "invalid",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			assert.Error(t, err)
		})
	}
}
