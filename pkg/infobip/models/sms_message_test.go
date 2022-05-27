package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidSMSMsg(t *testing.T) {
	tests := []struct {
		name     string
		instance SMSMsg
	}{
		{
			name: "minimum input",
			instance: SMSMsg{
				Destinations: []SMSDestination{
					{To: "1212345678"},
				},
			},
		},
		{
			name:     "full input",
			instance: GenerateSMSMsg(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)
		})
	}
}

func TestInvalidSMSMsg(t *testing.T) {
	tests := []struct {
		name     string
		instance SMSMsg
	}{
		{
			name:     "empty input",
			instance: SMSMsg{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
