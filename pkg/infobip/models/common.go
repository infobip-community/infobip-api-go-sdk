package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

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
}

// Validatable should be implemented by all models which represent request payloads.
// It will be called before a request is made.
type Validatable interface {
	Validate() error
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
