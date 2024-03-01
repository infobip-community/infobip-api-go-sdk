package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidUpdatePurchasedNumberRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdatePurchasedNumberRequest
	}{
		{
			name:     "missing applicationID field",
			instance: UpdatePurchasedNumberRequest{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}

func TestValidUpdatePurchasedNumberRequest(t *testing.T) {
	tests := []struct {
		name     string
		instance UpdatePurchasedNumberRequest
	}{
		{
			name: "valid",
			instance: UpdatePurchasedNumberRequest{
				ApplicationID: "f4sd5f4s5d4fd4",
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
