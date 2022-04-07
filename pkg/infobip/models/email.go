package models

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type EmailMsg struct {
	From                    string `validate:"required"`
	To                      string `validate:"required"`
	Cc                      string
	Bcc                     string
	Subject                 string `validate:"required"`
	Text                    string
	BulkId                  string
	MessageId               string
	TemplateId              int
	Attachment              *os.File
	InlineImage             *os.File
	HTML                    string
	ReplyTo                 string
	DefaultPlaceholders     string
	PreserveRecipients      bool
	TrackingURL             string
	TrackClicks             bool
	TrackOpens              bool
	Track                   bool
	CallbackData            string
	IntermediateReport      bool
	NotifyURL               string
	NotifyContentType       string
	SendAt                  string
	LandingPagePlaceholders string
	LandingPageId           string
	boundary                string
}

type SendEmailResponse struct {
	Messages []struct {
		To           string `json:"to"`
		MessageCount int    `json:"messageCount"`
		MessageId    string `json:"messageId"`
		Status       struct {
			GroupId     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
	} `json:"messages"`
}

type EmailMsgResponseStatus struct {
	GroupId     string `json:"groupId"`
	GroupName   string `json:"groupName"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *EmailMsg) Marshal() (*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	multipartWriter.Boundary()
	var partWriter io.Writer
	var err error

	if e.From != "" {
		err = writeMultipartText(multipartWriter, "from", e.From)
		if err != nil {
			return nil, err
		}
	}

	if e.To != "" {
		err = writeMultipartText(multipartWriter, "to", e.To)
		if err != nil {
			return nil, err
		}
	}

	if e.Cc != "" {
		err = writeMultipartText(multipartWriter, "cc", e.Cc)
		if err != nil {
			return nil, err
		}
	}

	if e.Bcc != "" {
		err = writeMultipartText(multipartWriter, "bcc", e.Bcc)
		if err != nil {
			return nil, err
		}
	}

	if e.Subject != "" {
		err = writeMultipartText(multipartWriter, "subject", e.Subject)
		if err != nil {
			return nil, err
		}
	}

	if e.Text != "" {
		err = writeMultipartText(multipartWriter, "text", e.Text)
		if err != nil {
			return nil, err
		}
	}

	if e.BulkId != "" {
		err = writeMultipartText(multipartWriter, "bulkId", e.BulkId)
		if err != nil {
			return nil, err
		}
	}

	if e.MessageId != "" {
		err = writeMultipartText(multipartWriter, "messageId", e.MessageId)
		if err != nil {
			return nil, err
		}
	}

	if e.TemplateId != 0 {
		err = writeMultipartText(multipartWriter, "templateid", fmt.Sprint(e.TemplateId))
		if err != nil {
			return nil, err
		}
	}

	if e.Attachment != nil {
		defer e.Attachment.Close()
		partWriter, err = multipartWriter.CreateFormFile("attachment", e.Attachment.Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(partWriter, e.Attachment)
		if err != nil {
			return nil, err
		}
	}

	if e.InlineImage != nil {
		defer e.InlineImage.Close()
		partWriter, err = multipartWriter.CreateFormFile("inlineImage", e.InlineImage.Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(partWriter, e.InlineImage)
		if err != nil {
			return nil, err
		}
	}

	if e.HTML != "" {
		err = writeMultipartText(multipartWriter, "HTML", e.HTML)
		if err != nil {
			return nil, err
		}
	}

	if e.ReplyTo != "" {
		err = writeMultipartText(multipartWriter, "replyto", e.ReplyTo)
		if err != nil {
			return nil, err
		}
	}

	if e.DefaultPlaceholders != "" {
		err = writeMultipartText(multipartWriter, "defaultplaceholders", e.DefaultPlaceholders)
		if err != nil {
			return nil, err
		}
	}

	if e.PreserveRecipients {
		err = writeMultipartText(multipartWriter, "preserverecipients", "true")
		if err != nil {
			return nil, err
		}
	}

	if e.TrackingURL != "" {
		err = writeMultipartText(multipartWriter, "trackingUrl", e.TrackingURL)
		if err != nil {
			return nil, err
		}
	}

	if e.TrackClicks {
		err = writeMultipartText(multipartWriter, "trackclicks", "true")
		if err != nil {
			return nil, err
		}
	}

	if e.TrackOpens {
		err = writeMultipartText(multipartWriter, "trackopens", "true")
		if err != nil {
			return nil, err
		}
	}

	if e.Track {
		err = writeMultipartText(multipartWriter, "track", "true")
		if err != nil {
			return nil, err
		}
	}

	if e.CallbackData != "" {
		err = writeMultipartText(multipartWriter, "callbackData", e.CallbackData)
		if err != nil {
			return nil, err
		}
	}

	if e.IntermediateReport {
		err = writeMultipartText(multipartWriter, "intermediateReport", "true")
		if err != nil {
			return nil, err
		}
	}

	if e.NotifyURL != "" {
		err = writeMultipartText(multipartWriter, "notifyUrl", e.NotifyURL)
		if err != nil {
			return nil, err
		}
	}

	if e.NotifyContentType != "" {
		err = writeMultipartText(multipartWriter, "notifyContentType", e.NotifyContentType)
		if err != nil {
			return nil, err
		}
	}

	if e.SendAt != "" {
		err = writeMultipartText(multipartWriter, "sendAt", e.SendAt)
		if err != nil {
			return nil, err
		}
	}

	if e.LandingPagePlaceholders != "" {
		err = writeMultipartText(multipartWriter, "landingPagePlaceholders", e.LandingPagePlaceholders)
		if err != nil {
			return nil, err
		}
	}

	if e.LandingPageId != "" {
		err = writeMultipartText(multipartWriter, "landingPageId", e.LandingPageId)
		if err != nil {
			return nil, err
		}
	}

	multipartWriter.Close()
	e.boundary = multipartWriter.Boundary()
	return &buf, nil
}

func (e *EmailMsg) GetMultipartBoundary() string {
	return e.boundary
}

type EmailDeliveryReportsResult struct {
	Results []struct {
		BulkId       string    `json:"bulkId"`
		MessageId    string    `json:"messageId"`
		To           string    `json:"to"`
		SentAt       time.Time `json:"sentAt"`
		DoneAt       time.Time `json:"doneAt"`
		MessageCount int       `json:"messageCount"`
		Price        struct {
			PricePerMessage int    `json:"pricePerMessage"`
			Currency        string `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupId     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
		} `json:"status"`
		Error struct {
			GroupId     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Permanent   bool   `json:"permanent"`
		} `json:"error"`
	} `json:"results"`
}

type EmailLogsResult struct {
	Results []struct {
		MessageId    string    `json:"messageId"`
		To           string    `json:"to"`
		From         string    `json:"from"`
		Text         string    `json:"text"`
		SentAt       time.Time `json:"sentAt"`
		DoneAt       time.Time `json:"doneAt"`
		MessageCount int       `json:"messageCount"`
		Price        struct {
			PricePerMessage int    `json:"pricePerMessage"`
			Currency        string `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupId     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
		} `json:"status"`
		BulkId string `json:"bulkId"`
	} `json:"results"`
}

type SentBulksResult struct {
	ExternalBulkId string `json:"externalBulkId"`
	Bulks          []struct {
		BulkId string    `json:"bulkId"`
		SendAt time.Time `json:"sendAt"`
	} `json:"bulks"`
}

func (e *EmailMsg) Validate() error {
	return validate.Struct(e)
}
