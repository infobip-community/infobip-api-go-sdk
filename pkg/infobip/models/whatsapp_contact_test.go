package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidContactMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance WAContactMsg
	}{
		{
			name: "minimum input",
			instance: WAContactMsg{
				MsgCommon: MsgCommon{
					From: "16175551213",
					To:   "16175551212",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "complete input",
			instance: WAContactMsg{
				MsgCommon: GenerateTestMsgCommon(),
				Content: ContactContent{
					Contacts: []Contact{
						{
							Addresses: []ContactAddress{
								{
									Street:      "Istarska",
									City:        "Vodnjan",
									Zip:         "52215",
									Country:     "Croatia",
									CountryCode: "HR",
									Type:        "WORK",
								},
								{
									Street:      "Istarska",
									City:        "Vodnjan",
									Zip:         "52215",
									Country:     "Croatia",
									CountryCode: "HR",
									Type:        "HOME",
								},
							},
							Birthday: "2010-12-20",
							Emails: []ContactEmail{
								{Email: "John.Smith@example.com", Type: "WORK"},
								{Email: "John.Smith.priv@example.com", Type: "HOME"},
							},
							Name: ContactName{
								FirstName:     "John",
								LastName:      "Smith",
								MiddleName:    "B",
								NamePrefix:    "Mr.",
								FormattedName: "Mr. John Smith",
							},
							Org: ContactOrg{Company: "Company Name", Department: "Department", Title: "Director"},
							Phones: []ContactPhone{
								{Phone: "+441134960019", Type: "HOME", WaID: "441134960019"},
								{Phone: "+441134960000", Type: "WORK", WaID: "441134960000"},
							},
							Urls: []ContactURL{
								{URL: "https://example.com/John.Smith", Type: "WORK"},
								{URL: "https://example.com/home/John.Smith", Type: "HOME"},
							},
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

func TestContactMessageConstraints(t *testing.T) {
	msgCommon := GenerateTestMsgCommon()
	tests := []struct {
		name    string
		content ContactContent
	}{
		{
			name:    "empty Content field",
			content: ContactContent{},
		},
		{
			name: "missing Contact FirstName",
			content: ContactContent{
				Contacts: []Contact{{Name: ContactName{FormattedName: "Mr. John Smith"}}},
			},
		},
		{
			name: "missing Contact FormattedName",
			content: ContactContent{
				Contacts: []Contact{{Name: ContactName{FirstName: "John"}}},
			},
		},
		{
			name: "invalid Contact Address Type",
			content: ContactContent{
				Contacts: []Contact{
					{
						Addresses: []ContactAddress{{Type: "Invalid"}},
						Name:      ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
					},
				},
			},
		},
		{
			name: "invalid Contact Birthday",
			content: ContactContent{
				Contacts: []Contact{
					{
						Birthday: "2020-22-12",
						Name:     ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
					},
				},
			},
		},
		{
			name: "invalid Contact Email",
			content: ContactContent{
				Contacts: []Contact{
					{
						Emails: []ContactEmail{{Email: "invalid"}},
						Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
					},
				},
			},
		},
		{
			name: "invalid Contact Email type",
			content: ContactContent{
				Contacts: []Contact{
					{
						Emails: []ContactEmail{{Email: "email@domain.com", Type: "invalid"}},
						Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
					},
				},
			},
		},
		{
			name: "invalid Contact Phone Type",
			content: ContactContent{
				Contacts: []Contact{
					{
						Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						Phones: []ContactPhone{{Type: "invalid"}},
					},
				},
			},
		},
		{
			name: "invalid Contact URL Type",
			content: ContactContent{
				Contacts: []Contact{
					{
						Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						Urls: []ContactURL{{Type: "Invalid"}},
					},
				},
			},
		},
		{
			name: "invalid Contact URL",
			content: ContactContent{
				Contacts: []Contact{
					{
						Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						Urls: []ContactURL{{URL: "asd"}},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := WAContactMsg{
				MsgCommon: msgCommon,
				Content:   tc.content,
			}
			err := msg.Validate()
			require.NotNil(t, err)
		})
	}
}
