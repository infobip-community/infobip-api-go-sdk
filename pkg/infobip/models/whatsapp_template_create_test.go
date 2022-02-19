package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidTemplateCreate(t *testing.T) {
	tests := []struct {
		name     string
		instance TemplateCreate
	}{
		{
			name: "minimum input",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "complete input, header TEXT",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT", Text: "Some text"},
					Body:   "body {{1}} content",
					Footer: "Footer text",
					Type:   "TEXT",
				},
			},
		},
		{
			name: "complete input, button PHONE_NUMBER",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT", Text: "Some text"},
					Body:   "body {{1}} content",
					Footer: "Footer text",
					Buttons: []TemplateButton{
						{Type: "PHONE_NUMBER", Text: "Button text", PhoneNumber: "16175551213"},
					},
					Type: "TEXT",
				},
			},
		},
		{
			name: "complete input, button URL",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT", Text: "Some text"},
					Body:   "body {{1}} content",
					Footer: "Footer text",
					Buttons: []TemplateButton{
						{Type: "URL", Text: "Button text", URL: "https://www.google.com"},
					},
					Type: "TEXT",
				},
			},
		},
		{
			name: "complete input, button QUICK_REPLY",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT", Text: "Some text"},
					Body:   "body {{1}} content",
					Footer: "Footer text",
					Buttons: []TemplateButton{
						{Type: "QUICK_REPLY", Text: "Button text"},
					},
					Type: "TEXT",
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

func TestTemplateCreateConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance TemplateCreate
	}{
		{
			name:     "empty payload",
			instance: TemplateCreate{},
		},
		{
			name: "missing Name",
			instance: TemplateCreate{
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "invalid Name",
			instance: TemplateCreate{
				Name:     "INVALID",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "missing Language",
			instance: TemplateCreate{
				Name:     "template_name",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "invalid Language",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "invalid",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "missing Category",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "invalid Category",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "invalid category",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
				},
			},
		},
		{
			name: "missing Structure",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
			},
		},
		{
			name: "invalid Structure Header Type",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "INVALID", Text: "Text"},
					Body:   "body {{1}} content",
					Type:   "TEXT",
				},
			},
		},
		{
			name: "missing Structure Header Text for type TEXT",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT"},
					Body:   "body {{1}} content",
					Type:   "TEXT",
				},
			},
		},
		{
			name: "invalid Structure Header Text for type TEXT",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Header: &TemplateHeader{Format: "TEXT", Text: strings.Repeat("a", 61)},
					Body:   "body {{1}} content",
					Type:   "TEXT",
				},
			},
		},
		{
			name: "missing Structure Body",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Type: "TEXT",
				},
			},
		},
		{
			name: "invalid Structure Footer",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body:   "body {{1}} content",
					Type:   "TEXT",
					Footer: strings.Repeat("a", 61),
				},
			},
		},
		{
			name: "empty Structure Buttons",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body:    "body {{1}} content",
					Type:    "TEXT",
					Buttons: []TemplateButton{},
				},
			},
		},
		{
			name: "invalid Structure Button type",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "INVALID", Text: "Some text"},
					},
				},
			},
		},
		{
			name: "missing Structure Button text",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "PHONE_NUMBER", PhoneNumber: "16175551213"},
					},
				},
			},
		},
		{
			name: "invalid Structure Button text",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "PHONE_NUMBER", PhoneNumber: "16175551213", Text: strings.Repeat("a", 201)},
					},
				},
			},
		},
		{
			name: "missing Structure Button PhoneNumber for type PHONE_NUMBER",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "PHONE_NUMBER", Text: "Some text"},
					},
				},
			},
		},
		{
			name: "missing Structure Button URL for type URL",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "URL", Text: "Some text"},
					},
				},
			},
		},
		{
			name: "missing Structure Button Text for type QUICK_REPLY",
			instance: TemplateCreate{
				Name:     "template_name",
				Language: "en",
				Category: "ACCOUNT_UPDATE",
				Structure: TemplateStructure{
					Body: "body {{1}} content",
					Type: "TEXT",
					Buttons: []TemplateButton{
						{Type: "QUICK_REPLY"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.NotNil(t, err)
		})
	}
}
