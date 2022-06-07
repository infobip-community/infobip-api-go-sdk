package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidWebRTCApplication(t *testing.T) {
	tests := []struct {
		name     string
		instance WebRTCApplication
	}{
		{
			name: "minimum input",
			instance: WebRTCApplication{
				Name: "some-name",
			},
		},
		{
			name:     "full input",
			instance: GenerateWebRTCApplication(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled WebRTCApplication
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidWebRTCApplication(t *testing.T) {
	tests := []struct {
		name     string
		instance WebRTCApplication
	}{
		{
			name:     "empty input",
			instance: WebRTCApplication{},
		},
		{
			name: "missing name",
			instance: WebRTCApplication{
				Description: "some-description",
			},
		},
		{
			name: "missing ios certificate file name",
			instance: WebRTCApplication{
				Name: "some-name",
				IOS: &WebRTCIOS{
					ApnsCertificateFileContent: "some-content",
					ApnsCertificatePassword:    "some-password",
				},
			},
		},
		{
			name: "missing ios certificate file content",
			instance: WebRTCApplication{
				Name: "some-name",
				IOS: &WebRTCIOS{
					ApnsCertificateFileName: "some-name",
					ApnsCertificatePassword: "some-password",
				},
			},
		},
		{
			name: "missing android fcm server key",
			instance: WebRTCApplication{
				Name: "some-name",
				Android: &WebRTCAndroid{
					FcmServerKey: "",
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
