package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidSendRCSBulkRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendRCSBulkRequest
	}{
		{
			name: "minimum input",
			instance: SendRCSBulkRequest{
				Messages: []RCSMsg{
					GenerateRCSFileMsg(),
				},
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

			var unmarshalled SendRCSBulkRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidSendRCSBulkRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance SendRCSBulkRequest
	}{
		{
			name: "empty to",
			instance: SendRCSBulkRequest{
				Messages: []RCSMsg{
					{},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
