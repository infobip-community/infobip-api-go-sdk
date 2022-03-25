package models

type EmailMsg struct {
	// TODO: Implement builder pattern.
	From                    string `json:"from" validate:"required"`
	To                      string `json:"to" validate:"required"`
	Cc                      string `json:"cc"`
	Bcc                     string `json:"bcc"`
	Subject                 string `json:"subject" validate:"required"`
	Text                    string `json:"text"`
	BulkId                  string `json:"bulkId"`
	MessageId               string `json:"messageId"`
	TemplateId              int    `json:"templateid"`
	Attachment              string `json:"attachment"`
	InlineImage             string `json:"inlineImage"`
	HTML                    string `json:"HTML"`
	ReplyTo                 string `json:"replyto"`
	DefaultPlaceholders     string `json:"defaultplaceholders"`
	PreserveRecipients      bool   `json:"preserverecipients"`
	TrackingURL             string `json:"trackingUrl"`
	TrackClicks             bool   `json:"trackclicks"`
	TrackOpens              bool   `json:"trackopens"`
	Track                   bool   `json:"track"`
	CallbackData            string `json:"callbackData"`
	IntermediateReport      bool   `json:"intermediateReport"`
	NotifyURL               string `json:"notifyUrl"`
	NotifyContentType       string `json:"notifyContentType"`
	SendAt                  string `json:"sendAt"`
	LandingPagePlaceholders string `json:"landingPagePlaceholders"`
	LandingPageId           string `json:"landingPageId"`
}

type EmailResponse struct {
	Messages []EmailMsg `json:"messages"`
}

func (e *EmailMsg) Validate() error {
	return validate.Struct(e)
}
