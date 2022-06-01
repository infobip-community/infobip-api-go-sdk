package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetScheduledSMSParams(t *testing.T) {
	t.Run("ValidGetScheduledMessagesParams", func(t *testing.T) {
		test := GetScheduledSMSParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetScheduledSMSParams(t *testing.T) {
	tests := []struct {
		name     string
		instance GetScheduledSMSParams
	}{
		{
			name:     "empty",
			instance: GetScheduledSMSParams{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
