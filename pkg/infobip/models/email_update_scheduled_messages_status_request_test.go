package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledMessagesStatusRequest(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := UpdateScheduledEmailStatusRequest{
			Status: "PENDING",
		}
		err := instance.Validate()
		require.NoError(t, err)

		marshalled, err := instance.Marshal()
		require.NoError(t, err)
		assert.NotEmpty(t, marshalled)

		var unmarshalled UpdateScheduledEmailStatusRequest
		err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
		require.NoError(t, err)
		assert.Equal(t, instance, unmarshalled)
	})
}

func TestInvalidUpdateScheduledMessagesStatusRequest(t *testing.T) {
	t.Run("invalid status", func(t *testing.T) {
		instance := UpdateScheduledEmailStatusRequest{
			Status: "SOMETHING",
		}
		err := instance.Validate()
		require.Error(t, err)
	})
}
