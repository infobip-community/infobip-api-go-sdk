package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledMessagesStatusOpts(t *testing.T) {
	t.Run("ValidUpdateScheduledMessagesStatusOpts", func(t *testing.T) {
		test := UpdateScheduledMessagesStatusOpts{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidUpdateScheduledMessagesStatusOpts(t *testing.T) {
	t.Run("InvalidUpdateScheduledMessagesStatusOpts", func(t *testing.T) {
		test := UpdateScheduledMessagesStatusOpts{}
		err := test.Validate()
		require.Error(t, err)
	})
}
