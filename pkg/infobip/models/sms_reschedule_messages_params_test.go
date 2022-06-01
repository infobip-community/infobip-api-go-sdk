package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidRescheduleSMSParams(t *testing.T) {
	t.Run("ValidRescheduleSMSParams", func(t *testing.T) {
		test := RescheduleSMSParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidRescheduleSMSParams(t *testing.T) {
	tests := []struct {
		name     string
		instance RescheduleSMSParams
	}{
		{
			name:     "empty",
			instance: RescheduleSMSParams{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
