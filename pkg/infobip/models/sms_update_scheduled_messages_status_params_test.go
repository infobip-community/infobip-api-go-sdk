package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledSMSStatusParams(t *testing.T) {
	t.Run("ValidUpdateScheduledSMSStatusParams", func(t *testing.T) {
		test := UpdateScheduledSMSStatusParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidUpdateScheduledSMSStatusParams(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateScheduledSMSStatusParams
	}{
		{
			name:     "empty",
			instance: UpdateScheduledSMSStatusParams{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
