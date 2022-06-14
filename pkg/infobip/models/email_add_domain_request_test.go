package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidAddDomainRequest(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := AddEmailDomainRequest{
			DomainName: "somedomain.com",
		}
		err := instance.Validate()
		require.NoError(t, err)

		marshalled, err := instance.Marshal()
		require.NoError(t, err)
		assert.NotEmpty(t, marshalled)

		var unmarshalled AddEmailDomainRequest
		err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
		require.NoError(t, err)
		assert.Equal(t, instance, unmarshalled)
	})
}

func TestInvalidAddDomainRequest(t *testing.T) {
	t.Run("empty request", func(t *testing.T) {
		instance := AddEmailDomainRequest{}
		err := instance.Validate()
		require.Error(t, err)
	})
}
