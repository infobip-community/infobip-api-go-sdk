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
		DefaultPlaceholders:     "somePlaceholders",
		PreserveRecipients:      true,
		TrackingURL:             "https://tracking.com",
		TrackClicks:             true,
		TrackOpens:              true,
		Track:                   true,
		CallbackData:            "someData",
		IntermediateReport:      true,
		NotifyURL:               "https://someurl.com",
		NotifyContentType:       "application/json",
		SendAt:                  "2022-01-01T00:00:00Z",
		LandingPagePlaceholders: "somePlaceholders",
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
		Latitude:     1.0,
		Longitude:    1.0,
		Label:        "some-label",
	}
}

func GenerateWebRTCApplication() WebRTCApplication {
	return WebRTCApplication{
		Name:        "some-name",
		Description: "some-description",
		IOS: &WebRTCIOS{
			ApnsCertificateFileName:    "some-file.cert",
			ApnsCertificateFileContent: "some-content",
			ApnsCertificatePassword:    "some-password",
		},
		Android: &WebRTCAndroid{
			FcmServerKey: "some-key",
		},
		AppToApp:           false,
		AppToConversations: false,
		AppToPhone:         false,
	}
}

func GenerateRCSFileMsg() RCSMsg {
	return RCSMsg{
		From:                   "some gopher",
		To:                     "12345678910",
		ValidityPeriod:         1,
		ValidityPeriodTimeUnit: "HOURS",
		Content: &RCSContent{
			Type: "FILE",
			File: &RCSFile{
				URL: "https://some-url.com",
			},
			Thumbnail: &RCSThumbnail{
				URL: "https://some-url.com",
			},
		},
		SMSFailover: &RCSSMSFailover{
			From:                   "some-gopher",
			Text:                   "some-text",
			ValidityPeriod:         1,
			ValidityPeriodTimeUnit: "MINUTES",
		},
		NotifyURL:    "https://some.url",
		CallbackData: "some-callback-data",
		MessageID:    "some-id",
	}
}

func GenerateRCSCardContent() *RCSCardContent {
	return &RCSCardContent{
		Title:       "some-title",
		Description: "some-description",
		Media: &RCSCardContentMedia{
			File: &RCSFile{
				URL: "https://some-url.com",
			},
			Thumbnail: &RCSThumbnail{
				URL: "https://some-url.com",
			},
			Height: "MEDIUM",
		},
		Suggestions: []RCSSuggestion{
			{
				Text:         "some-text",
				PostbackData: "some-postback-data",
				Type:         "REPLY",
			},
		},
	}
}

func GenerateCreateTFAApplicationRequest() CreateTFAApplicationRequest {
	return CreateTFAApplicationRequest{
		ApplicationID: "ABC1234ABC",
		Configuration: &TFAApplicationConfiguration{
			AllowMultiplePINVerifications: true,
			PINAttempts:                   5,
			PINTimeToLive:                 "10m",
			SendPINPerApplicationLimit:    "5000/12h",
			SendPINPerPhoneNumberLimit:    "2/1d",
			VerifyPINLimit:                "2/4s",
		},
		Enabled: true,
		Name:    "some-name",
	}
}

func GenerateUpdateTFAApplicationRequest() UpdateTFAApplicationRequest {
	return UpdateTFAApplicationRequest{
		ApplicationID: "ABC1234ABC",
		Configuration: &TFAApplicationConfiguration{
			AllowMultiplePINVerifications: true,
			PINAttempts:                   5,
			PINTimeToLive:                 "10m",
			SendPINPerApplicationLimit:    "5000/12h",
			SendPINPerPhoneNumberLimit:    "2/1d",
			VerifyPINLimit:                "2/4s",
		},
		Enabled: true,
		Name:    "some-name",
	}
}

func GenerateCreateTFAMessageTemplateRequest() CreateTFAMessageTemplateRequest {
	return CreateTFAMessageTemplateRequest{
		ApplicationID:  "ABC1234",
		Language:       English,
		MessageID:      "ABC1234",
		MessageText:    "Hello {{name}}. Your PIN is {{pin}}",
		PINLength:      4,
		PINPlaceholder: "{{pin}}",
		PINType:        NUMERIC,
		Regional: &SMSRegional{
			IndiaDLT{
				ContentTemplateID: "some-id",
				PrincipalEntityID: "some-id",
			},
		},
		RepeatDTMF: "1#",
		SenderID:   "Infobip 2FA",
		SpeechRate: 1,
	}
}

func GenerateUpdateTFAMessageTemplateRequest() UpdateTFAMessageTemplateRequest {
	return UpdateTFAMessageTemplateRequest{
		ApplicationID:  "ABC1234",
		Language:       English,
		MessageID:      "ABC1234",
		MessageText:    "Hello {{name}}. Your PIN is {{pin}}",
		PINLength:      4,
		PINPlaceholder: "{{pin}}",
		PINType:        NUMERIC,
		Regional: &SMSRegional{
			IndiaDLT{
				ContentTemplateID: "some-id",
				PrincipalEntityID: "some-id",
			},
		},
		RepeatDTMF: "1#",
		SenderID:   "Infobip 2FA",
		SpeechRate: 1,
	}
}

func GenerateSendPINOverSMSRequest() SendPINOverSMSRequest {
	return SendPINOverSMSRequest{
		ApplicationID: "ABC1234",
		MessageID:     "ABC1234",
		From:          "Sender",
		To:            "12345678910",
		Placeholders: map[string]string{
			"name": "some-name",
		},
	}
}
func GenerateSendPINOverVoiceRequest() SendPINOverVoiceRequest {
	return SendPINOverVoiceRequest{
		ApplicationID: "ABC1234",
		MessageID:     "ABC1234",
		From:          "Sender",
		To:            "12345678910",
		Placeholders: map[string]string{
			"name": "some-name",
		},
	}
}

func GenerateResendPINOverSMSRequest() ResendPINOverSMSRequest {
	return ResendPINOverSMSRequest{
		Placeholders: map[string]string{
			"name": "some-name",
		},
	}
}

func GenerateResendPINOverVoiceRequest() ResendPINOverVoiceRequest {
	return ResendPINOverVoiceRequest{
		Placeholders: map[string]string{
			"name": "some-name",
		},
	}
}

func GenerateVerifyPhoneNumberRequest() VerifyPhoneNumberRequest {
	return VerifyPhoneNumberRequest{
		PIN: "1234",
	}
}
