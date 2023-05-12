package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidVerifyPhoneNumberRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance VerifyPhoneNumberRequest
	}{
		{
			name: "minimum input",
			instance: VerifyPhoneNumberRequest{
				PIN: "123456",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled VerifyPhoneNumberRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidVerifyPhoneNumberRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance VerifyPhoneNumberRequest
	}{
		{
			name:     "empty",
			instance: VerifyPhoneNumberRequest{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
