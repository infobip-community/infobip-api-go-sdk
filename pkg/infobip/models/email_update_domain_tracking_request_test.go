package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidUpdateDomainTrackingRequest(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := UpdateEmailDomainTrackingRequest{
			Opens:       false,
			Clicks:      false,
			Unsubscribe: false,
		}
		err := instance.Validate()
		require.NoError(t, err)

		marshalled, err := instance.Marshal()
		require.NoError(t, err)
		assert.NotEmpty(t, marshalled)

		var unmarshalled UpdateEmailDomainTrackingRequest
		err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
		require.NoError(t, err)
		assert.Equal(t, instance, unmarshalled)
	})
}
