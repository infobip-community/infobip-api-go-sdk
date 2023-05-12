package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidUpdateTFAApplicationRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateTFAApplicationRequest
	}{
		{
			name: "minimum input",
			instance: UpdateTFAApplicationRequest{
				Name: "some-name",
			},
		},
		{
			name:     "full input",
			instance: GenerateUpdateTFAApplicationRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled UpdateTFAApplicationRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidUpdateTFARequestRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateTFAApplicationRequest
	}{
		{
			name:     "empty",
			instance: UpdateTFAApplicationRequest{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
