package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSentBulksParams(t *testing.T) {
	t.Run("ValidGetSentBulksParams", func(t *testing.T) {
		test := GetSentEmailBulksParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetSentBulksParams(t *testing.T) {
	t.Run("InvalidGetSentBulksParams", func(t *testing.T) {
		test := GetSentEmailBulksParams{}
		err := test.Validate()
		require.Error(t, err)
	})
}
