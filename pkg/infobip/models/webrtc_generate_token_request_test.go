package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidGenerateWebRTCTokenRequestTest(t *testing.T) {
	tests := []struct {
		name     string
		instance GenerateWebRTCTokenRequest
	}{
		{
			name: "minimum input",
			instance: GenerateWebRTCTokenRequest{
				Identity: "some-identity",
				Capabilities: &WebRTCTokenCapabilities{
					Recording: "ALWAYS",
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

			var unmarshalled GenerateWebRTCTokenRequest
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidGenerateWebRTCTokenRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance GenerateWebRTCTokenRequest
	}{
		{
			name:     "empty input",
			instance: GenerateWebRTCTokenRequest{},
		},
		{
			name: "invalid identity",
			instance: GenerateWebRTCTokenRequest{
				Identity: "a",
			},
		},
		{
			name: "invalid capabilities recording",
			instance: GenerateWebRTCTokenRequest{
				Identity: "some-identity",
				Capabilities: &WebRTCTokenCapabilities{
					Recording: "invalid",
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
