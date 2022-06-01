package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidBinarySMSMsg(t *testing.T) {
	tests := []struct {
		name     string
		instance BinarySMSMsg
	}{
		{
			name: "minimum input",
			instance: BinarySMSMsg{
				Destinations: []SMSDestination{
					{To: "1212345678"},
				},
			},
		},
		{
			name:     "full input",
			instance: GenerateBinarySMSMsg(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)

			marshalled, err := tc.instance.Marshal()
			assert.NotEmpty(t, marshalled)
			require.NoError(t, err)

			var unmarshalled BinarySMSMsg
			err = json.Unmarshal(marshalled.Bytes(), &unmarshalled)
			require.NoError(t, err)
			assert.Equal(t, tc.instance, unmarshalled)
		})
	}
}

func TestInvalidBinarySMSMsg(t *testing.T) {
	tests := []struct {
		name     string
		instance SMSMsg
	}{
		{
			name:     "empty input",
			instance: SMSMsg{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
