package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidRescheduleEmailRequest(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := RescheduleEmailRequest{
			SendAt: "2020-01-01T00:00:00Z",
		}
		err := instance.Validate()
		require.NoError(t, err)

		marshalled, err := instance.Marshal()
		require.NoError(t, err)
		assert.NotEmpty(t, marshalled)

		var unmarshalled RescheduleEmailRequest
		err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
		require.NoError(t, err)
		assert.Equal(t, instance, unmarshalled)
	})
}
