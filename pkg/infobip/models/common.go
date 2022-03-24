package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/go-playground/validator/v10"
)

const MinsPerHour = 60

var validate *validator.Validate //nolint: gochecknoglobals // thread safe and needed only once, caches validations

func init() {
	SetupValidation()
}

// SetupValidation configures struct-level validations for all payload models. This method must be
// called in order for validation to work, and it is invoked automatically when models are imported.
func SetupValidation() {
	if validate != nil {
		return
	}
	validate = validator.New()
	setupWhatsAppValidations()
	setupMMSValidations()
}

// Validatable should be implemented by all models which represent request payloads.
// It will be called before a request is made.
type Validatable interface {
	Validate() error
	Marshal() (*bytes.Buffer, error)
}

type MultipartValidatable interface {
	Validatable
	GetMultipartBoundary() string
}

func marshalJSON(t interface{}) (*bytes.Buffer, error) {
	payload, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(payload), nil
}

func escapeQuotes(s string) string {
	quoteEscaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	return quoteEscaper.Replace(s)
}

func writeMultipart(writer *multipart.Writer, fieldName string, content []byte, contentType string) error {
	header := textproto.MIMEHeader{}
	header.Set("Content-Type", contentType)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldName)))
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}

	_, err = part.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func writeMultipartJSON(writer *multipart.Writer, fieldName string, payload interface{}) error {
	content, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return writeMultipart(writer, fieldName, content, "application/json")
}

func writeMultipartText(writer *multipart.Writer, fieldName string, text string) error {
	return writeMultipart(writer, fieldName, []byte(text), "text/plain")
}

func writeMultipartXMLString(writer *multipart.Writer, fieldName string, rawXML string) error {
	return writeMultipart(writer, fieldName, []byte(rawXML), "application/xml")
}

type ResponseDetails struct {
	ErrorResponse ErrorDetails
	HTTPResponse  http.Response
}

type ErrorDetails struct {
	RequestError RequestError `json:"requestError"`
}

type RequestError struct {
	ServiceException ServiceException `json:"serviceException"`
}

type ServiceException struct {
	MessageID        string                 `json:"messageId"`
	Text             string                 `json:"text"`
	ValidationErrors map[string]interface{} `json:"validationErrors"`
}
