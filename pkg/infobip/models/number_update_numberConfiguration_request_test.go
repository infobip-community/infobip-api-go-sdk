package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidUpdateNumberConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateNumberConfigurationRequest
	}{
		{
			name:     "missing action configuration",
			instance: UpdateNumberConfigurationRequest{},
		},
		{
			name: "missing url in action configuration field",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{},
			},
		},
		{
			name: "bad format url in action configuration field",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{
					URL: "bad format",
				},
			},
		},
		{
			name: "bad format contentYype in action configuration field",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{
					URL:         "http://google.com",
					ContentType: "bad format",
				},
			},
		},
		{
			name: "bad format httpMethod in action configuration field",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{
					URL:        "http://google.com",
					HTTPMethod: "bad format",
				},
			},
		},
		{
			name: "bad format Type in action configuration field",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{
					URL:  "http://google.com",
					Type: "bad format",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}

func TestValidUpdateNumberConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdateNumberConfigurationRequest
	}{
		{
			name: "good format",
			instance: UpdateNumberConfigurationRequest{
				Action: &ActionConfiguration{
					URL: "http://google.com",
				},
				Key: "key",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.NoError(t, err)
		})
	}
}
