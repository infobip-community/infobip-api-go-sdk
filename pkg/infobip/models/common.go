package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate //nolint: gochecknoglobals // thread safe and needed only once, caches validations

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
