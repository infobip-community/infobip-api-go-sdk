package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/utils"
	"github.com/stretchr/testify/require"
)

func TestValidTemplateMessage(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		instance WATemplateMsgs
	}{
		{
			name: "minimum input",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: MsgCommon{From: "16175551213", To: "16175551212"},
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "minimum input, empty placeholders",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: MsgCommon{From: "16175551213", To: "16175551212"},
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{}},
							},
							Language: "en_GB",
						},
					},
				},
				BulkID: "100",
			},
		},
		{
			name: "complete input, header TEXT",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "TEXT", Placeholder: "Placeholder header value"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, header DOCUMENT",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:     "DOCUMENT",
									MediaURL: "https://myurl.com/asd.pdf",
									Filename: "asd.pdf",
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, header IMAGE",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "IMAGE", MediaURL: "https://myurl.com/asd.jpg"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, header VIDEO",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "VIDEO", MediaURL: "https://myurl.com/asd.mp4"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, header LOCATION",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:      "LOCATION",
									Latitude:  utils.Float32Ptr(73.5164),
									Longitude: utils.Float32Ptr(56.2502),
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, button QUICK_REPLY",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Buttons: []TemplateMsgButton{
									{Type: "QUICK_REPLY", Parameter: "Some parameter"},
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "complete input, button URL",
			instance: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Buttons: []TemplateMsgButton{
									{
										Type:      "URL",
										Parameter: fmt.Sprintf("over 128 cha%srs", strings.Repeat("a", 128)),
									},
								},
							},
							Language: "en_GB",
						},
					},
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

func TestTemplateConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name     string
		messages WATemplateMsgs
	}{
		{
			name:     "empty messages",
			messages: WATemplateMsgs{},
		},
		{
			name: "missing Content",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
					},
				},
			},
		},
		{
			name: "missing TemplateName",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid TemplateName format",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "Invalid Format",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid TemplateName length",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: strings.Repeat("a", 513),
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing TemplateData",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							Language:     "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing TemplateData body",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{},
							Language:     "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Body placeholder",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", ""}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header Type",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Placeholder: "Text"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Header Type",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "invalid", Placeholder: "Text"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header Placeholder for type TEXT",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "TEXT"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header Filename for type DOCUMENT",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "DOCUMENT", MediaURL: "https://www.myurl.com/1.pdf"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Header Filename for type DOCUMENT",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:     "DOCUMENT",
									MediaURL: "https://www.myurl.com/1.pdf",
									Filename: strings.Repeat("a", 241),
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header MediaURL for type DOCUMENT",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "DOCUMENT", Filename: "asd.pdf"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Header MediaURL for type DOCUMENT",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:     "DOCUMENT",
									MediaURL: fmt.Sprintf("https://%srl.com/asd.pdf", strings.Repeat("a", 2048)),
									Filename: "asd.pdf",
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header MediaURL for type IMAGE",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "IMAGE"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header MediaURL for type VIDEO",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "VIDEO"},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header Latitude for type LOCATION",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "LOCATION", Longitude: utils.Float32Ptr(10.55)},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Header Longitude for type LOCATION",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:   TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{Type: "LOCATION", Latitude: utils.Float32Ptr(10.55)},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Header Latitude for type LOCATION",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:      "LOCATION",
									Latitude:  utils.Float32Ptr(91.5),
									Longitude: utils.Float32Ptr(10.5),
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Header Longitude for type LOCATION",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Placeholder Value1", "Placeholder Value2"}},
								Header: &TemplateMsgHeader{
									Type:      "LOCATION",
									Latitude:  utils.Float32Ptr(10.5),
									Longitude: utils.Float32Ptr(181.5),
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Button Type",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:    TemplateBody{Placeholders: []string{"Value1", "Value2"}},
								Buttons: []TemplateMsgButton{{Type: "invalid", Parameter: "payload"}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "invalid Button Parameter for Type QUICK_REPLY",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body:    TemplateBody{Placeholders: []string{"Value1", "Value2"}},
								Buttons: []TemplateMsgButton{{Type: "QUICK_REPLY", Parameter: strings.Repeat("a", 129)}},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "count over max for QUICK_REPLY Buttons",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
								Buttons: []TemplateMsgButton{
									{Type: "QUICK_REPLY", Parameter: "value1"},
									{Type: "QUICK_REPLY", Parameter: "value2"},
									{Type: "QUICK_REPLY", Parameter: "value3"},
									{Type: "QUICK_REPLY", Parameter: "value4"},
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "count over max for URL Buttons",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
								Buttons: []TemplateMsgButton{
									{Type: "URL", Parameter: "value1"},
									{Type: "URL", Parameter: "value2"},
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "contains both QUICK_REPLY and URL Buttons types",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
								Buttons: []TemplateMsgButton{
									{Type: "QUICK_REPLY", Parameter: "value1"},
									{Type: "URL", Parameter: "value2"},
								},
							},
							Language: "en_GB",
						},
					},
				},
			},
		},
		{
			name: "missing Content Language",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
						},
					},
				},
			},
		},
		{
			name: "missing SMSFailover From",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
							Language: "en_GB",
						},
						SMSFailover: &SMSFailover{Text: "Text"},
					},
				},
			},
		},
		{
			name: "invalid SMSFailover From",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
							Language: "en_GB",
						},
						SMSFailover: &SMSFailover{From: strings.Repeat("1", 25), Text: "Text"},
					},
				},
			},
		},
		{
			name: "missing SMSFailover Text",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
							Language: "en_GB",
						},
						SMSFailover: &SMSFailover{From: "16175551213"},
					},
				},
			},
		},
		{
			name: "invalid SMSFailover Text",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
							Language: "en_GB",
						},
						SMSFailover: &SMSFailover{From: "16175551213", Text: strings.Repeat("a", 4097)},
					},
				},
			},
		},
		{
			name: "invalid BulkID",
			messages: WATemplateMsgs{
				Messages: []TemplateMsg{
					{
						MsgCommon: msgCommon,
						Content: TemplateMsgContent{
							TemplateName: "template_name",
							TemplateData: TemplateData{
								Body: TemplateBody{Placeholders: []string{"Value1", "Value2"}},
							},
							Language: "en_GB",
						},
					},
				},
				BulkID: strings.Repeat("1", 101),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.messages.Validate()
			require.NotNil(t, err)
		})
	}
}
