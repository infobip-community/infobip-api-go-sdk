package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidRescheduleMessagesOpts(t *testing.T) {
	t.Run("ValidRescheduleMessagesOpts", func(t *testing.T) {
		test := RescheduleMessagesOpts{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidRescheduleMessagesOpts(t *testing.T) {
	t.Run("InvalidRescheduleMessagesOpts", func(t *testing.T) {
		test := RescheduleMessagesOpts{}
		err := test.Validate()
		require.Error(t, err)
	})
}
