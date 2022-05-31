package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidSendSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendSMSRequest
	}{
		{
			name: "minimum input",
			instance: SendSMSRequest{
				Messages: []SMSMsg{
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
			instance: GenerateSendSMSRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled SendSMSRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidSendSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendSMSRequest
	}{
		{
			name: "missing messages",
			instance: SendSMSRequest{
				Messages: []SMSMsg{},
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
