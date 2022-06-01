package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidSendSMSOverQueryParamsParams(t *testing.T) {
	t.Run("ValidSendSMSOverQueryParamsParams", func(t *testing.T) {
		test := SendSMSOverQueryParamsParams{
			Username: "some-username",
			Password: "some-password",
			To:       []string{"123456789012"},
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidSendSMSOverQueryParamsParams(t *testing.T) {
	tests := []struct {
		name     string
		instance SendSMSOverQueryParamsParams
	}{
		{
			name:     "empty",
			instance: SendSMSOverQueryParamsParams{},
		},
		{
			name: "no username",
			instance: SendSMSOverQueryParamsParams{
				Password: "some-password",
				To:       []string{"123456789012"},
			},
		},
		{
			name: "no password",
			instance: SendSMSOverQueryParamsParams{
				Username: "some-username",
				To:       []string{"123456789012"},
			},
		},
		{
			name: "no to",
			instance: SendSMSOverQueryParamsParams{
				Username: "some-username",
				Password: "some-password",
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
