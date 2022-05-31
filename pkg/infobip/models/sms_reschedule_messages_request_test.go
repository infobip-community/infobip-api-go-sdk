package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidRescheduleSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance RescheduleSMSRequest
	}{
		{
			name: "minimum input",
			instance: RescheduleSMSRequest{
				SendAt: "2020-01-01T00:00:00Z",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled RescheduleSMSRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidRescheduleSMSRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance RescheduleSMSRequest
	}{
		{
			name:     "missing sendAt",
			instance: RescheduleSMSRequest{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
