package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSentBulksOpts(t *testing.T) {
	t.Run("ValidGetSentBulksOpts", func(t *testing.T) {
		test := GetSentBulksOpts{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetDelieryReportsOpts(t *testing.T) {
	t.Run("InvalidGetSentBulksOpts", func(t *testing.T) {
		test := GetSentBulksOpts{}
		err := test.Validate()
		require.Error(t, err)
	})
}
