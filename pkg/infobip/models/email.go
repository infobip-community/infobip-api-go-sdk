package models

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type EmailMsg struct {
	From                    string `validate:"required"`
	To                      string `validate:"required"`
	Cc                      string
	Bcc                     string
	Subject                 string `validate:"required"`
	Text                    string
	BulkID                  string
	MessageID               string
	TemplateID              int
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
	LandingPageID           string
	boundary                string
}

type SendEmailResponse struct {
	BulkID   string `json:"bulkId"`
	Messages []struct {
		To           string `json:"to"`
		MessageCount int    `json:"messageCount"`
		MessageID    string `json:"messageId"`
		Status       struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
	} `json:"messages"`
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

	if e.BulkID != "" {
		err = writeMultipartText(multipartWriter, "bulkId", e.BulkID)
		if err != nil {
			return nil, err
		}
	}

	if e.MessageID != "" {
		err = writeMultipartText(multipartWriter, "messageId", e.MessageID)
		if err != nil {
			return nil, err
		}
	}

	if e.TemplateID != 0 {
		err = writeMultipartText(multipartWriter, "templateid", fmt.Sprint(e.TemplateID))
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

	if e.LandingPageID != "" {
		err = writeMultipartText(multipartWriter, "landingPageId", e.LandingPageID)
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

type EmailDeliveryReportsResponse struct {
	Results []struct {
		BulkID    string `json:"bulkId"`
		MessageID string `json:"messageId"`
		To        string `json:"to"`
		// TODO: this is a string, but it should be a time.Time
		SentAt       string `json:"sentAt"`
		DoneAt       string `json:"doneAt"`
		MessageCount int    `json:"messageCount"`
		Price        struct {
			PricePerMessage float64 `json:"pricePerMessage"`
			Currency        string  `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
		} `json:"status"`
		Error struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Permanent   bool   `json:"permanent"`
		} `json:"error"`
		Channel string `json:"channel"`
	} `json:"results"`
}

type GetDeliveryReportsOpts struct {
	BulkID    string
	MessageID string
	Limit     string
}

type EmailLogsResponse struct {
	Results []struct {
		MessageID    string `json:"messageId"`
		To           string `json:"to"`
		From         string `json:"from"`
		Text         string `json:"text"`
		SentAt       string `json:"sentAt"`
		DoneAt       string `json:"doneAt"`
		MessageCount int    `json:"messageCount"`
		Price        struct {
			PricePerMessage float64 `json:"pricePerMessage"`
			Currency        string  `json:"currency"`
		} `json:"price"`
		Status struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
		} `json:"status"`
		BulkID  string `json:"bulkId"`
		Channel string `json:"channel"`
	} `json:"results"`
}

type GetLogsOpts struct {
	MessageID     string
	From          string
	To            string
	BulkID        string
	GeneralStatus string
	SentSince     string
	SentUntil     string
	Limit         string
}

type SentEmailBulksResponse struct {
	ExternalBulkID string `json:"externalBulkId"`
	Bulks          []struct {
		BulkID string `json:"bulkId"`
		SendAt int64  `json:"sendAt"`
	} `json:"bulks"`
}

type GetSentBulksOpts struct {
	BulkID string `validate:"required"`
}

type SentEmailBulksStatusResponse struct {
	ExternalBulkID string `json:"externalBulkId"`
	Bulks          []struct {
		BulkID string `json:"bulkId"`
		Status string `json:"status"`
	} `json:"bulks"`
}

type GetSentBulksStatusOpts struct {
	BulkID string `validate:"required"`
}

type RescheduleMessagesRequest struct {
	SendAt string `json:"sendAt"`
}

type RescheduleMessagesOpts struct {
	BulkID string `validate:"required"`
}

type UpdateScheduledMessagesStatusRequest struct {
	Status string `json:"status"`
}

type UpdateScheduledMessagesStatusOpts struct {
	BulkID string `validate:"required"`
}

func (r *UpdateScheduledMessagesStatusRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateScheduledMessagesStatusRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type UpdateScheduledMessagesStatusResponse struct {
	BulkID string `json:"bulkId"`
	Status string `json:"status"`
}

func (r *RescheduleMessagesRequest) Validate() error {
	return validate.Struct(r)
}

func (r *RescheduleMessagesRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type RescheduleMessagesResponse struct {
	BulkID string `json:"bulkId"`
	SendAt int64  `json:"sendAt"`
}

func (e *EmailMsg) Validate() error {
	return validate.Struct(e)
}

type ValidateAddressesRequest struct {
	To string `json:"to"`
}

func (v *ValidateAddressesRequest) Validate() error {
	return validate.Struct(v)
}

func (v *ValidateAddressesRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(v)
}

type ValidateAddressesResponse struct {
	To           string `json:"to"`
	ValidMailbox string `json:"validMailbox"`
	ValidSyntax  bool   `json:"validSyntax"`
	CatchAll     bool   `json:"catchAll"`
	Disposable   bool   `json:"disposable"`
	RoleBased    bool   `json:"roleBased"`
}
