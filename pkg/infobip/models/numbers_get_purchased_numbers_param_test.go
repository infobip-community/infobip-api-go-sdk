package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetPurchasedNumbersParam(t *testing.T) {
	tests := []struct {
		name     string
		instance ListPurchasedNumbersParam
	}{
		{
			name:     "valid",
			instance: ListPurchasedNumbersParam{},
		},
		{
			name: "another valid format",
			instance: ListPurchasedNumbersParam{
				Limit: 1,
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

func TestIncalidGetPurchasedNumbersParam(t *testing.T) {
	tests := []struct {
		name     string
		instance ListPurchasedNumbersParam
	}{
		{
			name:     "low limit",
			instance: ListPurchasedNumbersParam{Limit: -1},
		},
		{
			name:     "high limit",
			instance: ListPurchasedNumbersParam{Limit: 2000},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.instance.Validate()
			require.Error(t, err)
		})
	}
}
