package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidRescheduleMessagesRequest(t *testing.T) {
	t.Run("ValidRescheduleMessagesRequest", func(t *testing.T) {
		test := RescheduleMessagesRequest{
			SendAt: "2016-01-01T00:00:00Z",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidRescheduleMessagesRequest(t *testing.T) {
	t.Run("InvalidRescheduleMessagesRequest", func(t *testing.T) {
		test := RescheduleMessagesRequest{}
		err := test.Validate()
		require.Error(t, err)
	})
}
