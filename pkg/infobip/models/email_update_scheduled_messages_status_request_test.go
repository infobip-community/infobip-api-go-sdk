package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledMessagesStatusRequest(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := UpdateScheduledEmailMessagesStatusRequest{
			Status: "PENDING",
		}
		err := instance.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidUpdateScheduledMessagesStatusRequest(t *testing.T) {
	t.Run("invalid status", func(t *testing.T) {
		instance := UpdateScheduledEmailMessagesStatusRequest{
			Status: "SOMETHING",
		}
		err := instance.Validate()
		require.Error(t, err)
	})
}
