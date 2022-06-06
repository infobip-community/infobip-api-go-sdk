package models

func GenerateTestMsgCommon() MsgCommon {
	return MsgCommon{
		From:         "16175551213",
		To:           "16175551212",
		MessageID:    "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
		CallbackData: "some data",
		NotifyURL:    "https://www.google.com",
	}
}

func GenerateEmailMsg() EmailMsg {
	return EmailMsg{
		From:                    "someoneg@selfserviceib.com",
		To:                      "someone@gmail.com",
		Cc:                      "somemail@mail.com",
		Bcc:                     "anothermail@mail.com",
		Subject:                 "Some subject",
		Text:                    "Some text",
		BulkID:                  "esy82u725261jz8e6pi3",
		MessageID:               "somexternalMessageId0",
		TemplateID:              1,
		Attachment:              nil,
		InlineImage:             nil,
		HTML:                    "<body>Some html</body>",
		ReplyTo:                 "reply@infobip.com",
		DefaultPlaceholders:     "someplaceholders",
		PreserveRecipients:      true,
		TrackingURL:             "https://tracking.com",
		TrackClicks:             true,
		TrackOpens:              true,
		Track:                   true,
		CallbackData:            "somedata",
		IntermediateReport:      true,
		NotifyURL:               "https://someurl.com",
		NotifyContentType:       "application/json",
		SendAt:                  "2022-01-01T00:00:00Z",
		LandingPagePlaceholders: "someplaceholders",
		LandingPageID:           "123456",
	}
}

func GenerateSendSMSRequest() SendSMSRequest {
	return SendSMSRequest{
		BulkID:   "some-bulk-id",
		Messages: []SMSMsg{GenerateSMSMsg()},
		SendingSpeedLimit: &SMSSendingSpeedLimit{
			Amount:   1,
			TimeUnit: "MINUTE",
		},
		Tracking: &SMSTracking{
			BaseURL:    "https://tracking.com",
			Track:      "SMS",
			Type:       "ONE_TIME_PIN",
			ProcessKey: "someKey",
		},
	}
}

func GenerateBinarySMSMsg() BinarySMSMsg {
	return BinarySMSMsg{
		From: "Gopher",
		Destinations: []SMSDestination{
			{
				To: "16175551212",
			},
		},
		Binary: &SMSBinary{
			Hex:        "aa bb cc dd ff",
			DataCoding: 0,
			EsmClass:   0,
		},
		IntermediateReport: false,
		NotifyURL:          "https:/some.url",
		NotifyContentType:  "application/json",
		CallbackData:       "some-callback-data",
		ValidityPeriod:     0,
		SendAt:             "10-10-2020T10:10:10Z",
		DeliveryTimeWindow: &SMSDeliveryTimeWindow{
			Days: []string{"MONDAY"},
			From: SMSTime{
				Hour:   1,
				Minute: 0,
			},
			To: SMSTime{
				Hour:   1,
				Minute: 1,
			},
		},
		Regional: &SMSRegional{
			IndiaDLT{
				ContentTemplateID: "some-id",
				PrincipalEntityID: "some-principal-id",
			},
		},
	}
}

func GenerateSendBinarySMSRequest() SendBinarySMSRequest {
	return SendBinarySMSRequest{
		BulkID:   "some-bulk-id",
		Messages: []BinarySMSMsg{GenerateBinarySMSMsg()},
	}
}

func GenerateSMSMsg() SMSMsg {
	return SMSMsg{
		CallbackData: "DLR callback data",
		Destinations: []SMSDestination{{
			MessageID: "some-id",
			To:        "16175551212",
		}},
		Flash:              false,
		From:               "Someone",
		IntermediateReport: false,
		Language: &SMSLanguage{
			LanguageCode: "EN",
		},
		NotifyContentType: "application/json",
		NotifyURL:         "https://someurl.com",
		Text:              "Some content",
		Transliteration:   "CENTRAL_EUROPEAN",
		ValidityPeriod:    1,
		DeliveryTimeWindow: &SMSDeliveryTimeWindow{
			Days: []string{"MONDAY"},
			From: SMSTime{
				Hour:   1,
				Minute: 0,
			},
			To: SMSTime{
				Hour:   1,
				Minute: 1,
			},
		},
		SendAt: "10-10-2020T10:10:10Z",
		Regional: &SMSRegional{
			IndiaDLT{
				ContentTemplateID: "some-id",
				PrincipalEntityID: "some-principal-id",
			},
		},
	}
}

func GeneratePreviewSMSRequest() PreviewSMSRequest {
	return PreviewSMSRequest{
		LanguageCode:    "TR",
		Text:            "Let's see how many characters will remain unused in this message.",
		Transliteration: "TURKISH",
	}
}

func GenerateReplyRCSSuggestion() RCSSuggestion {
	return RCSSuggestion{
		Text:         "some-text",
		PostbackData: "some-postback-data",
		Type:         "REPLY",
	}
}

func GenerateOpenURLRCSSuggestion() RCSSuggestion {
	return RCSSuggestion{
		Text:         "some-text",
		PostbackData: "some-postback-data",
		Type:         "OPEN_URL",
		URL:          "https://some-url.com",
	}
}

func GenerateDialPhoneRCSSuggestion() RCSSuggestion {
	return RCSSuggestion{
		Text:         "some-text",
		PostbackData: "some-postback-data",
		Type:         "DIAL_PHONE",
		PhoneNumber:  "12345678910",
	}
}

func GenerateShowLocationRCSSuggestion() RCSSuggestion {
	return RCSSuggestion{
		Text:         "some-text",
		PostbackData: "some-postback-data",
		Type:         "SHOW_LOCATION",
		Latitude:     82.123,
		Longitude:    123.123,
		Label:        "some-label",
	}
}
