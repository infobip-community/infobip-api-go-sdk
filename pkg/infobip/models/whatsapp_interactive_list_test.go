package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidInteractiveListMessage(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name     string
		instance InteractiveListMessage
	}{
		{
			name: "minimum input",
			instance: InteractiveListMessage{
				MessageCommon: MessageCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: InteractiveListContent{
					Body: InteractiveListBody{Text: "Some text"},
					Action: InteractiveListAction{
						Title:    "Choose one",
						Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
					},
				},
			},
		},
		{
			name: "complete input",
			instance: InteractiveListMessage{
				MessageCommon: msgCommon,
				Content: InteractiveListContent{
					Body: InteractiveListBody{Text: "Some text"},
					Action: InteractiveListAction{
						Title: "some title",
						Sections: []Section{
							{Title: "Title 1", Rows: []SectionRow{{ID: "1", Title: "Row1 Title", Description: "desc"}}},
							{Title: "Title 2", Rows: []SectionRow{{ID: "2", Title: "Row2 Title", Description: "desc"}}},
						},
					},
					Header: &InteractiveListHeader{
						Type: "TEXT",
						Text: "Header text",
					},
					Footer: &InteractiveListFooter{Text: "Footer text"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Nil(t, err)
		})
	}
}

func TestTextInteractiveListConstraints(t *testing.T) {
	msgCommon := GenerateTestMessageCommon()
	tests := []struct {
		name    string
		content InteractiveListContent
	}{
		{
			name:    "empty Content field",
			content: InteractiveListContent{},
		},
		{
			name: "missing Body",
			content: InteractiveListContent{
				Action: InteractiveListAction{
					Title:    "Choose one",
					Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
				},
			},
		},
		{
			name: "missing Body Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{},
				Action: InteractiveListAction{
					Title:    "Choose one",
					Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
				},
			},
		},
		{
			name: "invalid Body Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{Text: strings.Repeat("a", 1025)},
				Action: InteractiveListAction{
					Title:    "Choose one",
					Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
				},
			},
		},
		{
			name: "missing Action",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
			},
		},
		{
			name: "missing Action Title",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
				},
			},
		},
		{
			name: "invalid Action Title",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title:    strings.Repeat("a", 21),
					Sections: []Section{{Rows: []SectionRow{{ID: "1", Title: "row title"}}}},
				},
			},
		},
		{
			name: "missing Action Sections",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
				},
			},
		},
		{
			name: "count over max for Action Sections",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{
						{Rows: []SectionRow{{ID: "1", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "2", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "3", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "4", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "5", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "6", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "7", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "8", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "9", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "10", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "11", Title: "row title"}}},
					},
				},
			},
		},
		{
			name: "invalid Action Section Title",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Title: strings.Repeat("a", 25),
						Rows:  []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
			},
		},
		{
			name: "missing Action Section Title when there are multiple sections",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{
						{Title: "section title", Rows: []SectionRow{{ID: "1", Title: "row title"}}},
						{Rows: []SectionRow{{ID: "1", Title: "row title"}}},
					},
				},
			},
		},
		{
			name: "missing Action Section Rows",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title:    "Choose one",
					Sections: []Section{{}},
				},
			},
		},
		{
			name: "count over max for Action Section Rows",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{
							{ID: "1", Title: "row title"},
							{ID: "2", Title: "row title"},
							{ID: "3", Title: "row title"},
							{ID: "4", Title: "row title"},
							{ID: "5", Title: "row title"},
							{ID: "6", Title: "row title"},
							{ID: "7", Title: "row title"},
							{ID: "8", Title: "row title"},
							{ID: "9", Title: "row title"},
							{ID: "10", Title: "row title"},
							{ID: "11", Title: "row title"},
						},
					}},
				},
			},
		},
		{
			name: "missing Action Section Row ID",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{Title: "row title"}},
					}},
				},
			},
		},
		{
			name: "invalid Action Section Row ID",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: strings.Repeat("a", 201), Title: "row title"}},
					}},
				},
			},
		},
		{
			name: "duplicate row ID for Action Section Rows, single section",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{
							{ID: "1", Title: "row title"},
							{ID: "2", Title: "row title"},
							{ID: "1", Title: "row title"},
						}},
					},
				},
			},
		},
		{
			name: "duplicate row ID for Action Section Rows, multiple sections",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{
						{Title: "First title", Rows: []SectionRow{{ID: "1", Title: "row title"}}},
						{Title: "Second title", Rows: []SectionRow{{ID: "1", Title: "row title"}}},
					},
				},
			},
		},
		{
			name: "missing Action Section Row Title",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1"}},
					}},
				},
			},
		},
		{
			name: "invalid Action Section Row Title",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{Title: strings.Repeat("a", 25), ID: "1"}},
					}},
				},
			},
		},
		{
			name: "invalid Action Section Row Description",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title", Description: strings.Repeat("a", 73)}},
					}},
				},
			},
		},
		{
			name: "missing Header Type",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Header: &InteractiveListHeader{Text: "some test"},
			},
		},
		{
			name: "invalid Header Type",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Header: &InteractiveListHeader{Type: "invalid", Text: "some test"},
			},
		},
		{
			name: "missing Header Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Header: &InteractiveListHeader{Type: "TEXT"},
			},
		},
		{
			name: "invalid Header Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Header: &InteractiveListHeader{Type: "TEXT", Text: strings.Repeat("a", 61)},
			},
		},
		{
			name: "missing Footer Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Footer: &InteractiveListFooter{},
			},
		},
		{
			name: "invalid Footer Text",
			content: InteractiveListContent{
				Body: InteractiveListBody{"Some text"},
				Action: InteractiveListAction{
					Title: "Choose one",
					Sections: []Section{{
						Rows: []SectionRow{{ID: "1", Title: "row title"}},
					}},
				},
				Footer: &InteractiveListFooter{Text: strings.Repeat("a", 61)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := InteractiveListMessage{
				MessageCommon: msgCommon,
				Content:       tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
