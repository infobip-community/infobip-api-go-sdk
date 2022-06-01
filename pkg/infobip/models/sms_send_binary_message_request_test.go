package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidSendBinarySMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendBinarySMSRequest
	}{
		{
			name: "minimum input",
			instance: SendBinarySMSRequest{
				Messages: []BinarySMSMsg{
					{
						Destinations: []SMSDestination{
							{To: "1212345678"},
						},
					},
				},
			},
		},
		{
			name:     "full input",
			instance: GenerateSendBinarySMSRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled SendBinarySMSRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidSendBinarySMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendBinarySMSRequest
	}{
		{
			name: "missing messages",
			instance: SendBinarySMSRequest{
				Messages: []BinarySMSMsg{},
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
