package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetScheduledSMSStatusParams(t *testing.T) {
	t.Run("ValidGetScheduledSMSStatusParams", func(t *testing.T) {
		test := GetScheduledSMSStatusParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetScheduledSMSStatusParams(t *testing.T) {
	tests := []struct {
		name     string
		instance GetScheduledSMSStatusParams
	}{
		{
			name:     "empty",
			instance: GetScheduledSMSStatusParams{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
