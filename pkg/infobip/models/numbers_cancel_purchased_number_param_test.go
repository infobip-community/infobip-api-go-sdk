package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidCancelPurchasedNumberParam(t *testing.T) {
	tests := []struct {
		name     string
		instance CancelPurchasedNumberParam
	}{
		{
			name:     "missing NumberKey field",
			instance: CancelPurchasedNumberParam{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}

func TestValidCancelPurchasedNumberParam(t *testing.T) {
	tests := []struct {
		name     string
		instance CancelPurchasedNumberParam
	}{
		{
			name: "valid",
			instance: CancelPurchasedNumberParam{
				NumberKey: "f4sd5f4s5d4fd4",
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
