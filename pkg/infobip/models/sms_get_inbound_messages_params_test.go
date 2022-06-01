package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetInboundMessagesParams(t *testing.T) {
	t.Run("ValidGetInboundMessagesParams", func(t *testing.T) {
		test := GetInboundSMSParams{
			Limit: 1,
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetInboundSMSParams(t *testing.T) {
	tests := []struct {
		name     string
		instance GetInboundSMSParams
	}{
		{
			name:     "low limit",
			instance: GetInboundSMSParams{Limit: -1},
		},
		{
			name:     "high limit",
			instance: GetInboundSMSParams{Limit: 2000},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
