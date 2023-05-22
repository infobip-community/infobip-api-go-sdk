package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidCreateTFAApplicationRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance CreateTFAApplicationRequest
	}{
		{
			name: "minimum input",
			instance: CreateTFAApplicationRequest{
				Name: "some application",
			},
		},
		{
			name:     "full input",
			instance: GenerateCreateTFAApplicationRequest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			require.NoError(t, err)
			assert.NotEmpty(t, marshalled)

			var unmarshalled CreateTFAApplicationRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidCreateTFAApplicationRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance CreateTFAApplicationRequest
	}{
		{
			name:     "empty",
			instance: CreateTFAApplicationRequest{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
