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
	Attachments             []*os.File
	InlineImages            []*os.File
	HTML                    string
	ReplyTo                 string
	DefaultPlaceholders     string
	PreserveRecipients      bool
	TrackingURL             string `validate:"omitempty,url"`
	TrackClicks             bool
	TrackOpens              bool
	Track                   bool
	CallbackData            string
	IntermediateReport      bool
	NotifyURL               string `validate:"omitempty,url"`
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

//nolint:cyclop,funlen,gocognit,gocyclo // Because the EmailMsg has too many fields.
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

	if e.Attachments != nil {
		for _, attachment := range e.Attachments {
			defer attachment.Close()
			partWriter, err = multipartWriter.CreateFormFile("attachment", attachment.Name())
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(partWriter, attachment)
			if err != nil {
				return nil, err
			}

		}
	}

	if e.InlineImages != nil {
		for _, image := range e.InlineImages {
			defer image.Close()
			partWriter, err = multipartWriter.CreateFormFile("inlineImage", image.Name())
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(partWriter, image)
			if err != nil {
				return nil, err
			}
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

type GetEmailDeliveryReportsResponse struct {
	Results []struct {
		BulkID       string `json:"bulkId"`
		MessageID    string `json:"messageId"`
		To           string `json:"to"`
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

type GetEmailDeliveryReportsParams struct {
	BulkID    string
	MessageID string
	Limit     int
}

type GetEmailLogsResponse struct {
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

type GetEmailLogsParams struct {
	MessageID     string
	From          string
	To            string
	BulkID        string
	GeneralStatus string
	SentSince     string
	SentUntil     string
	Limit         int
}

type SentEmailBulksResponse struct {
	ExternalBulkID string `json:"externalBulkId"`
	Bulks          []struct {
		BulkID string `json:"bulkId"`
		SendAt int64  `json:"sendAt"`
	} `json:"bulks"`
}

type GetSentEmailBulksParams struct {
	BulkID string `validate:"required"`
}

func (o *GetSentEmailBulksParams) Validate() error {
	return validate.Struct(o)
}

type SentEmailBulksStatusResponse struct {
	ExternalBulkID string `json:"externalBulkId"`
	Bulks          []struct {
		BulkID string `json:"bulkId"`
		Status string `json:"status"`
	} `json:"bulks"`
}

type GetSentEmailBulksStatusParams struct {
	BulkID string `validate:"required"`
}

func (o *GetSentEmailBulksStatusParams) Validate() error {
	return validate.Struct(o)
}

type RescheduleEmailRequest struct {
	SendAt string `json:"sendAt"`
}

type RescheduleEmailParams struct {
	BulkID string `validate:"required"`
}

func (o *RescheduleEmailParams) Validate() error {
	return validate.Struct(o)
}

type UpdateScheduledEmailStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=PENDING PAUSED PROCESSING CANCELED FINISHED FAILED"`
}

type UpdateScheduledEmailStatusParams struct {
	BulkID string `validate:"required"`
}

func (o *UpdateScheduledEmailStatusParams) Validate() error {
	return validate.Struct(o)
}

func (r *UpdateScheduledEmailStatusRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateScheduledEmailStatusRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type UpdateScheduledStatusResponse struct {
	BulkID string `json:"bulkId"`
	Status string `json:"status"`
}

func (r *RescheduleEmailRequest) Validate() error {
	return validate.Struct(r)
}

func (r *RescheduleEmailRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(r)
}

type RescheduleEmailResponse struct {
	BulkID string `json:"bulkId"`
	SendAt int64  `json:"sendAt"`
}

func (e *EmailMsg) Validate() error {
	return validate.Struct(e)
}

type ValidateEmailAddressesRequest struct {
	To string `json:"to" validate:"required,min=1,max=2147483647"`
}

func (v *ValidateEmailAddressesRequest) Validate() error {
	return validate.Struct(v)
}

func (v *ValidateEmailAddressesRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(v)
}

type ValidateEmailAddressesResponse struct {
	To           string `json:"to"`
	ValidMailbox string `json:"validMailbox"`
	ValidSyntax  bool   `json:"validSyntax"`
	CatchAll     bool   `json:"catchAll"`
	Disposable   bool   `json:"disposable"`
	RoleBased    bool   `json:"roleBased"`
}

type GetEmailDomainsParams struct {
	Size int `validate:"omitempty,min=1,max=20"`
	Page int `validate:"omitempty,min=0"`
}

type GetEmailDomainsResponse struct {
	Paging struct {
		Page         int `json:"page"`
		Size         int `json:"size"`
		TotalPages   int `json:"totalPages"`
		TotalResults int `json:"totalResults"`
	} `json:"paging"`
	Results []EmailDomain `json:"results"`
}

type AddEmailDomainRequest struct {
	DomainName string `json:"domainName" validate:"required"`
}

func (a *AddEmailDomainRequest) Validate() error {
	return validate.Struct(a)
}

func (a *AddEmailDomainRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(a)
}

type EmailDomain struct {
	DomainID   int64  `json:"domainId"`
	DomainName string `json:"domainName"`
	Active     bool   `json:"active"`
	Tracking   struct {
		Clicks      bool `json:"clicks"`
		Opens       bool `json:"opens"`
		Unsubscribe bool `json:"unsubscribe"`
	} `json:"tracking"`
	DNSRecords []struct {
		RecordType    string `json:"recordType"`
		Name          string `json:"name"`
		ExpectedValue string `json:"expectedValue"`
		Verified      bool   `json:"verified"`
	} `json:"dnsRecords"`
	Blocked   bool   `json:"blocked"`
	CreatedAt string `json:"createdAt"`
}

type AddEmailDomainResponse EmailDomain

type GetEmailDomainResponse EmailDomain

type UpdateEmailDomainTrackingRequest struct {
	Opens       bool `json:"open"`
	Clicks      bool `json:"clicks"`
	Unsubscribe bool `json:"unsubscribe"`
}

func (u *UpdateEmailDomainTrackingRequest) Validate() error {
	return validate.Struct(u)
}

func (u *UpdateEmailDomainTrackingRequest) Marshal() (*bytes.Buffer, error) {
	return marshalJSON(u)
}

type UpdateEmailDomainTrackingResponse EmailDomain
