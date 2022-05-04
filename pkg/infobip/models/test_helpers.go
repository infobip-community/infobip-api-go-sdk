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
		SendingSpeedLimit: &SendingSpeedLimit{
			Amount:   1,
			TimeUnit: "MINUTE",
		},
		Tracking: &Tracking{
			BaseURL:    "https://tracking.com",
			Track:      "SMS",
			Type:       "ONE_TIME_PIN",
			ProcessKey: "someKey",
		},
	}
}

func GenerateSMSMsg() SMSMsg {
	return SMSMsg{
		CallbackData: "DLR callback data",
		Destinations: []Destination{{
			MessageID: "some-id",
			To:        "16175551212",
		}},
		Flash:              false,
		From:               "Someone",
		IntermediateReport: false,
		Language: Language{
			LanguageCode: "EN",
		},
		NotifyContentType: "application/json",
		NotifyURL:         "https://someurl.com",
		Text:              "Some content",
		Transliteration:   "CENTRAL_EUROPEAN",
		ValidityPeriod:    1,
		SMSDeliveryTimeWindow: &SMSDeliveryTimeWindow{
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
	}
}
