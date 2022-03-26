package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidInteractiveMultiproductMessage(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		instance InteractiveMultiproductMsg
	}{
		{
			name: "minimum input",
			instance: InteractiveMultiproductMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: InteractiveMultiproductContent{
					Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Header"},
					Body:   InteractiveMultiproductBody{Text: "Some Text"},
					Action: InteractiveMultiproductAction{
						CatalogID: "1",
						Sections: []InteractiveMultiproductSection{
							{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
						},
					},
				},
			},
		},
		{
			name: "complete input",
			instance: InteractiveMultiproductMsg{
				MsgCommon: msgCommon,
				Content: InteractiveMultiproductContent{
					Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Header"},
					Body:   InteractiveMultiproductBody{Text: "Some Text"},
					Action: InteractiveMultiproductAction{
						CatalogID: "1",
						Sections: []InteractiveMultiproductSection{
							{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
							{Title: "Title 2", ProductRetailerIDs: []string{"1", "2"}},
						},
					},
					Footer: &InteractiveMultiproductFooter{Text: "Footer text"},
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

func TestTextInteractiveMultiproductConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content InteractiveMultiproductContent
	}{
		{
			name:    "empty Content field",
			content: InteractiveMultiproductContent{},
		},
		{
			name: "missing Header",
			content: InteractiveMultiproductContent{
				Body: InteractiveMultiproductBody{Text: "Some Text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "missing Header Type",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some Text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "invalid Header Type",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "invalid", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some Text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "missing Header Text",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT"},
				Body:   InteractiveMultiproductBody{Text: "Some Text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "invalid Header Text length",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: strings.Repeat("a", 61)},
				Body:   InteractiveMultiproductBody{Text: "Some Text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "missing Body",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "invalid Body Text length",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: strings.Repeat("a", 1025)},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "missing Action",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
			},
		},
		{
			name: "missing Action CatalogID",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					Sections: []InteractiveMultiproductSection{
						{Title: "Title", ProductRetailerIDs: []string{"1", "2"}},
					},
				},
			},
		},
		{
			name: "missing Action Sections",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
				},
			},
		},
		{
			name: "count over max for Action Sections",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Title1", ProductRetailerIDs: []string{"1"}},
						{Title: "Title2", ProductRetailerIDs: []string{"2"}},
						{Title: "Title3", ProductRetailerIDs: []string{"3"}},
						{Title: "Title4", ProductRetailerIDs: []string{"4"}},
						{Title: "Title5", ProductRetailerIDs: []string{"5"}},
						{Title: "Title6", ProductRetailerIDs: []string{"6"}},
						{Title: "Title7", ProductRetailerIDs: []string{"7"}},
						{Title: "Title8", ProductRetailerIDs: []string{"8"}},
						{Title: "Title9", ProductRetailerIDs: []string{"9"}},
						{Title: "Title10", ProductRetailerIDs: []string{"10"}},
						{Title: "Title11", ProductRetailerIDs: []string{"11"}},
					},
				},
			},
		},
		{
			name: "invalid Action Section title length",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: strings.Repeat("a", 25), ProductRetailerIDs: []string{"1"}},
					},
				},
			},
		},
		{
			name: "missing Action Section Title when there are multiple sections",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Section 1 title", ProductRetailerIDs: []string{"1"}},
						{ProductRetailerIDs: []string{"1"}},
					},
				},
			},
		},
		{
			name: "missing Action Section productRetailerIDs",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Section 1 title"},
					},
				},
			},
		},
		{
			name: "missing Action Section productRetailerIDs",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Section 1 title", ProductRetailerIDs: []string{}},
					},
				},
			},
		},
		{
			name: "missing Footer text",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Section 1 title", ProductRetailerIDs: []string{"1"}},
					},
				},
				Footer: &InteractiveMultiproductFooter{},
			},
		},
		{
			name: "invalid Footer text length",
			content: InteractiveMultiproductContent{
				Header: InteractiveMultiproductHeader{Type: "TEXT", Text: "Some text"},
				Body:   InteractiveMultiproductBody{Text: "Some text"},
				Action: InteractiveMultiproductAction{
					CatalogID: "1",
					Sections: []InteractiveMultiproductSection{
						{Title: "Section 1 title", ProductRetailerIDs: []string{"1"}},
					},
				},
				Footer: &InteractiveMultiproductFooter{Text: strings.Repeat("a", 61)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := InteractiveMultiproductMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
