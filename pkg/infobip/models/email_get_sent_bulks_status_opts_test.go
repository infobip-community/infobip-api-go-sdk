package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSentBulksStatusOpts(t *testing.T) {
	t.Run("ValidGetSentBulksStatusOpts", func(t *testing.T) {
		test := GetSentBulksStatusOpts{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetSentBulksStatusOpts(t *testing.T) {
	t.Run("InvalidGetSentBulksStatusOpts", func(t *testing.T) {
		test := GetSentBulksStatusOpts{}
		err := test.Validate()
		require.Error(t, err)
	})
}
