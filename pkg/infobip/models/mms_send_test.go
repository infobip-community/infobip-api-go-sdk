package models

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/utils"
	"github.com/stretchr/testify/require"
)

func TestValidMMSMsg(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	require.NoError(t, err)
	tests := []struct {
		name     string
		instance MMSMsg
	}{
		{
			name: "minimum input",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212"},
			},
		},
		{
			name: "complete input",
			instance: MMSMsg{
				Head: MMSHead{
					From:                  "16175551213",
					To:                    "16175551212",
					ID:                    "26",
					Subject:               "This is a sample message",
					ValidityPeriodMinutes: 10,
					CallbackData:          "data",
					NotifyURL:             "https://www.google.com",
					SendAt:                "2006-01-02T15:04:05.123Z",
					IntermediateReport:    utils.BoolPtr(true),
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{"MONDAY", "TUESDAY"},
						From: &MMSTime{Minute: 1, Hour: 1},
						To:   &MMSTime{Minute: 1, Hour: 2},
					},
				},
				Text:  "Some text",
				Media: tmpfile,
				ExternallyHostedMedia: []ExternallyHostedMedia{
					{
						ContentType: "image/jpeg",
						ContentID:   "1",
						ContentURL:  "https://myurl.com/asd.jpg",
					},
				},
				SMIL: "<smil></smil>",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err = tc.instance.Validate()
			require.NoError(t, err)
		})
	}
}

func TestMMSMsgConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance MMSMsg
	}{
		{
			name:     "missing Head",
			instance: MMSMsg{},
		},
		{
			name: "missing Head From",
			instance: MMSMsg{
				Head: MMSHead{To: "16175551213"},
			},
		},
		{
			name: "missing Head To",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213"},
			},
		},
		{
			name: "invalid Head callbackData length",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212", CallbackData: strings.Repeat("a", 201)},
			},
		},
		{
			name: "invalid Head notifyURL",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212", NotifyURL: "asd"},
			},
		},
		{
			name: "invalid sendAt format",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212", SendAt: "Thu, 01 Sep 2016 10:11:12.123456 -0500"},
			},
		},
		{
			name: "missing Head DeliveryTimeWindow Days",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						From: &MMSTime{Hour: 1, Minute: 1},
						To:   &MMSTime{Hour: 2, Minute: 1},
					},
				},
			},
		},
		{
			name: "invalid Head DeliveryTimeWindow Days",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{"INVALID"},
					},
				},
			},
		},
		{
			name: "invalid Head DeliveryTimeWindow Days",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{},
					},
				},
			},
		},
		{
			name: "missing Head DeliveryTimeWindow To with From provided",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{"MONDAY"},
						From: &MMSTime{Hour: 1, Minute: 1},
					},
				},
			},
		},
		{
			name: "missing Head DeliveryTimeWindow From with To provided",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{"MONDAY"},
						To:   &MMSTime{Hour: 1, Minute: 1},
					},
				},
			},
		},
		{
			name: "Head DeliveryTimeWindow From and To less than an hour apart",
			instance: MMSMsg{
				Head: MMSHead{
					From: "16175551213",
					To:   "16175551212",
					DeliveryTimeWindow: &DeliveryTimeWindow{
						Days: []string{"MONDAY"},
						From: &MMSTime{Hour: 1, Minute: 1},
						To:   &MMSTime{Hour: 2, Minute: 0},
					},
				},
			},
		},
		{
			name: "missing  ExternallyHostedMedia ContentType",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212"},
				ExternallyHostedMedia: []ExternallyHostedMedia{
					{ContentID: "1", ContentURL: "https://myurl.com/asd.jpg"},
				},
			},
		},
		{
			name: "missing  ExternallyHostedMedia ContentID",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212"},
				ExternallyHostedMedia: []ExternallyHostedMedia{
					{ContentType: "image/jpeg", ContentURL: "https://myurl.com/asd.jpg"},
				},
			},
		},
		{
			name: "missing  ExternallyHostedMedia ContentURL",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212"},
				ExternallyHostedMedia: []ExternallyHostedMedia{
					{ContentType: "image/jpeg", ContentID: "1"},
				},
			},
		},
		{
			name: "missing  ExternallyHostedMedia ContentURL",
			instance: MMSMsg{
				Head: MMSHead{From: "16175551213", To: "16175551212"},
				ExternallyHostedMedia: []ExternallyHostedMedia{
					{ContentType: "image/jpeg", ContentID: "1", ContentURL: "asd"},
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
