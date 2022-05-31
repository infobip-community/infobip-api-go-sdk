package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidGetSMSLogsParams(t *testing.T) {
	t.Run("ValidGetSMSLogsParams", func(t *testing.T) {
		test := GetSMSLogsParams{
			From:   "123456789012",
			To:     "123456789012",
			BulkID: []string{"some-bulk-id"},
			Limit:  1,
		}
		err := test.Validate()
		require.NoError(t, err)
	})
}

func TestInvalidGetSMSLogsParams(t *testing.T) {
	t.Run("InvalidGetSMSLogsParams", func(t *testing.T) {
		test := GetSMSLogsParams{
			From:          "123456789012",
			To:            "123456789012",
			BulkID:        []string{"some-bulk-id"},
			Limit:         0,
			GeneralStatus: "some-status",
		}
		err := test.Validate()
		require.Error(t, err)
	})
}
