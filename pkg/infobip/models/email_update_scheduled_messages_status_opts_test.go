package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledMessagesStatusParams(t *testing.T) {
	t.Run("ValidUpdateScheduledMessagesStatusParams", func(t *testing.T) {
		test := UpdateScheduledEmailMessagesStatusParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidUpdateScheduledMessagesStatusParams(t *testing.T) {
	t.Run("InvalidUpdateScheduledMessagesStatusParams", func(t *testing.T) {
		test := UpdateScheduledEmailMessagesStatusParams{}
		err := test.Validate()
		require.Error(t, err)
	})
}
