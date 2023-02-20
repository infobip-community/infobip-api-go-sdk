package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidCreateTFAMessageTemplateRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance CreateTFAMessageTemplateRequest
	}{
		{
			name: "minimum input",
			instance: CreateTFAMessageTemplateRequest{
				MessageText: "Hello, {{name}}! Your code is {{code}}.",
				PINLength:   4,
				PINType:     NUMERIC,
			},
		},
		{
			name:     "full input",
			instance: GenerateCreateTFAMessageTemplateRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled CreateTFAMessageTemplateRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidCreateTFAMessageTemplateRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance CreateTFAMessageTemplateRequest
	}{
		{
			name:     "empty",
			instance: CreateTFAMessageTemplateRequest{},
		},
		{
			name: "no message text",
			instance: CreateTFAMessageTemplateRequest{
				MessageText:    "",
				PINLength:      4,
				PINPlaceholder: "{{pin}}",
				PINType:        NUMERIC,
			},
		},
		{
			name: "no pin length",
			instance: CreateTFAMessageTemplateRequest{
				MessageText:    "Hello, {{name}}! Your code is {{code}}.",
				PINLength:      0,
				PINPlaceholder: "{{pin}}",
				PINType:        NUMERIC,
			},
		},
		{
			name: "no pin type",
			instance: CreateTFAMessageTemplateRequest{
				MessageText:    "Hello, {{name}}! Your code is {{code}}.",
				PINLength:      4,
				PINPlaceholder: "{{pin}}",
				PINType:        "",
			},
		},
		{
			name: "no India principal entity ID",
			instance: CreateTFAMessageTemplateRequest{
				MessageText:    "Hello, {{name}}! Your code is {{code}}.",
				PINLength:      4,
				PINPlaceholder: "{{pin}}",
				PINType:        NUMERIC,
				Regional: &SMSRegional{
					IndiaDLT{
						ContentTemplateID: "some-id",
						PrincipalEntityID: "",
					},
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
