package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidValidateEmailAddresses(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		instance := ValidateEmailAddressesRequest{
			To: "some@email.com",
		}
		err := instance.Validate()
		require.NoError(t, err)

		marshalled, err := instance.Marshal()
		require.NoError(t, err)
		assert.NotEmpty(t, marshalled)

		var unmarshalled ValidateEmailAddressesRequest
		err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
		require.NoError(t, err)
		assert.Equal(t, instance, unmarshalled)
	})
}

func TestInvalidValidateEmailAddresses(t *testing.T) {
	t.Run("empty to", func(t *testing.T) {
		instance := ValidateEmailAddressesRequest{
			To: "",
		}
		err := instance.Validate()
		require.Error(t, err)
	})
}
