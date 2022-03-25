package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidInteractiveProductMessage(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		instance InteractiveProductMsg
	}{
		{
			name: "minimum input",
			instance: InteractiveProductMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: InteractiveProductContent{
					Action: InteractiveProductAction{
						CatalogID:         "1",
						ProductRetailerID: "2",
					},
				},
			},
		},
		{
			name: "complete input",
			instance: InteractiveProductMsg{
				MsgCommon: msgCommon,
				Content: InteractiveProductContent{
					Action: InteractiveProductAction{
						CatalogID:         "1",
						ProductRetailerID: "2",
					},
					Body:   &InteractiveProductBody{Text: "Product text"},
					Footer: &InteractiveProductFooter{Text: "Footer text"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NoError(t, err)
		})
	}
}

func TestTextInteractiveProductConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content InteractiveProductContent
	}{
		{
			name:    "empty Content field",
			content: InteractiveProductContent{},
		},
		{
			name: "missing Action",
			content: InteractiveProductContent{
				Body: &InteractiveProductBody{Text: "Some text"},
			},
		},
		{
			name: "missing Action CatalogID",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					ProductRetailerID: "2",
				},
			},
		},
		{
			name: "missing Action ProductRetailerID",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					CatalogID: "1",
				},
			},
		},
		{
			name: "missing Body Text",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					CatalogID:         "1",
					ProductRetailerID: "2",
				},
				Body: &InteractiveProductBody{},
			},
		},
		{
			name: "invalid Body Text length",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					CatalogID:         "1",
					ProductRetailerID: "2",
				},
				Body: &InteractiveProductBody{Text: strings.Repeat("a", 1025)},
			},
		},
		{
			name: "missing Footer Text",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					CatalogID:         "1",
					ProductRetailerID: "2",
				},
				Footer: &InteractiveProductFooter{},
			},
		},
		{
			name: "invalid Footer Text length",
			content: InteractiveProductContent{
				Action: InteractiveProductAction{
					CatalogID:         "1",
					ProductRetailerID: "2",
				},
				Footer: &InteractiveProductFooter{Text: strings.Repeat("a", 61)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := InteractiveProductMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
