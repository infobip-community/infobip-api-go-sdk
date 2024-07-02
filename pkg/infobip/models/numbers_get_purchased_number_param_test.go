package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidGetPurchasedNumberParam(t *testing.T) {
	tests := []struct {
		name     string
		instance GetPurchasedNumberParam
	}{
		{
			name:     "missing NumberKey field",
			instance: GetPurchasedNumberParam{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}

func TestValidGetPurchasedNumberParam(t *testing.T) {
	tests := []struct {
		name     string
		instance GetPurchasedNumberParam
	}{
		{
			name: "valid",
			instance: GetPurchasedNumberParam{
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
