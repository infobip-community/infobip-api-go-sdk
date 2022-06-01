package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidUpdateScheduledSMSStatusRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateScheduledSMSStatusRequest
	}{
		{
			name: "minimum input",
			instance: UpdateScheduledSMSStatusRequest{
				Status: "CANCELED",
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

			var unmarshalled UpdateScheduledSMSStatusRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidUpdateScheduledSMSStasusRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateScheduledSMSStatusRequest
	}{
		{
			name:     "missing status",
			instance: UpdateScheduledSMSStatusRequest{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
