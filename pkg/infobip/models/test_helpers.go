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
