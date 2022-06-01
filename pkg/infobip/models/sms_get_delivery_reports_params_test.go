package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSMSDeliveryReportsParams(t *testing.T) {
	t.Run("ValidGetSMSDeliveryReportsParams", func(t *testing.T) {
		test := GetSMSDeliveryReportsParams{
			BulkID: "some-bulk-id",
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}
