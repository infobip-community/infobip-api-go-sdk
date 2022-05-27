package models

import (
	"testing"

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.NoError(t, err)
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
