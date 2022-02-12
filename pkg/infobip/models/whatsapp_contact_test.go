package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidContactMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance ContactMessage
	}{
		{
			name: "minimum input",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
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
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
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
			require.Nil(t, err)
		})
	}
}

func TestContactMessageConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance ContactMessage
	}{
		{
			name: "missing From field",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "",
					To:           "16175551213",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "missing To field",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "missing Content field",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
			},
		},
		{
			name: "From too long",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "1617555121333333333333333",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "To too long",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551212",
					To:           "1617555121333333333333333",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "MessageID too long",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    strings.Repeat("a", 51),
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "CallbackData too long",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: strings.Repeat("a", 4001),
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "NotifyURL too long",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    fmt.Sprintf("https://www.google%s.com", strings.Repeat("a", 4097)),
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "NotifyURL not an url",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "if only this was an url...",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "missing Content FirstName",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FormattedName: "Mr. John Smith"}}},
				},
			},
		},
		{
			name: "missing Content FormattedName",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{{Name: ContactName{FirstName: "John"}}},
				},
			},
		},
		{
			name: "invalid Content Address Type",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Addresses: []ContactAddress{{Type: "Invalid"}},
							Name:      ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						},
					},
				},
			},
		},
		{
			name: "invalid Content Birthday",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Birthday: "2020-22-12",
							Name:     ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						},
					},
				},
			},
		},
		{
			name: "invalid Content Email",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Emails: []ContactEmail{{Email: "invalid"}},
							Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						},
					},
				},
			},
		},
		{
			name: "invalid Content Email type",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Emails: []ContactEmail{{Email: "email@domain.com", Type: "invalid"}},
							Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
						},
					},
				},
			},
		},
		{
			name: "invalid Content Phone Type",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Name:   ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
							Phones: []ContactPhone{{Type: "invalid"}},
						},
					},
				},
			},
		},
		{
			name: "invalid Content URL Type",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
							Urls: []ContactURL{{Type: "Invalid"}},
						},
					},
				},
			},
		},
		{
			name: "invalid Content URL",
			instance: ContactMessage{
				MessageCommon: MessageCommon{
					From:         "16175551213",
					To:           "16175551212",
					MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
					CallbackData: "some data",
					NotifyURL:    "https://www.google.com",
				},
				Content: ContactContent{
					Contacts: []Contact{
						{
							Name: ContactName{FirstName: "John", FormattedName: "Mr. John Smith"},
							Urls: []ContactURL{{URL: "asd"}},
						},
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
