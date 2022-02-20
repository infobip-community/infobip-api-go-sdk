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
