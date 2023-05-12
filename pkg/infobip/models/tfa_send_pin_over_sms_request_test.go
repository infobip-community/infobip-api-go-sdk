package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidSendPINOverSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendPINOverSMSRequest
	}{
		{
			name: "minimum input",
			instance: SendPINOverSMSRequest{
				ApplicationID: "ABC123",
				MessageID:     "ABC123",
				To:            "555555555555",
				Placeholders:  map[string]string{"name": "John"},
			},
		},
		{
			name:     "full input",
			instance: GenerateSendPINOverSMSRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled SendPINOverSMSRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidSendPINOverSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendPINOverSMSRequest
	}{
		{
			name:     "empty",
			instance: SendPINOverSMSRequest{},
		},
		{
			name: "no application id",
			instance: SendPINOverSMSRequest{
				MessageID:    "ABC123",
				To:           "555555555555",
				Placeholders: map[string]string{"name": "John"},
			},
		},
		{
			name: "no message id",
			instance: SendPINOverSMSRequest{
				ApplicationID: "ABC123",
				To:            "555555555555",
				Placeholders:  map[string]string{"name": "John"},
			},
		},
		{
			name: "no to",
			instance: SendPINOverSMSRequest{
				ApplicationID: "ABC123",
				MessageID:     "ABC123",
				Placeholders:  map[string]string{"name": "John"},
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
