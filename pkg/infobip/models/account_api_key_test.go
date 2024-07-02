package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidApiKey(t *testing.T) {
	tests := []struct {
		name     string
		instance APIKey
	}{
		{
			name: "just with name field",
			instance: APIKey{
				Name: "testing",
			},
		},
		{
			name: "with name field and permissions",
			instance: APIKey{
				Name:        "testing",
				Permissions: []string{"PUBLIC_API", "WEB_SDK"},
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

func TestInvalidApiKey(t *testing.T) {
	tests := []struct {
		name     string
		instance APIKey
	}{
		{
			name:     "missing name field",
			instance: APIKey{},
		},
		{
			name: "missing applicationId field",
			instance: APIKey{
				Name: "testing",
				Platform: []Platform{{
					Key: "",
				}},
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
