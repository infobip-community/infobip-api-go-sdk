package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidGetAvailableNumbersParams(t *testing.T) {
	tests := []struct {
		name     string
		instance GetAvailableNumbersParams
	}{
		{
			name: "limit value more to 1000",
			instance: GetAvailableNumbersParams{
				Limit: 50000,
			},
		},
		{
			name: "page value more to 1000",
			instance: GetAvailableNumbersParams{
				Page: 50000,
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
