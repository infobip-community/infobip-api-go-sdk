package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSentBulksStatusParams(t *testing.T) {
	t.Run("ValidGetSentBulksStatusParams", func(t *testing.T) {
		test := GetSentEmailBulksStatusParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetSentBulksStatusParams(t *testing.T) {
	t.Run("InvalidGetSentBulksStatusParams", func(t *testing.T) {
		test := GetSentEmailBulksStatusParams{}
		err := test.Validate()
		require.Error(t, err)
	})
}
