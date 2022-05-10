package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidRescheduleMessagesParams(t *testing.T) {
	t.Run("ValidRescheduleMessagesParams", func(t *testing.T) {
		test := RescheduleEmailMessagesParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidRescheduleMessagesParams(t *testing.T) {
	t.Run("InvalidRescheduleMessagesParams", func(t *testing.T) {
		test := RescheduleEmailMessagesParams{}
		err := test.Validate()
		require.Error(t, err)
	})
}
