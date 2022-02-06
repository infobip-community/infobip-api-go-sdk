package models

import "github.com/go-playground/validator/v10"

var validate *validator.Validate //nolint: gochecknoglobals // thread safe and needed only once, caches validations

type Validatable interface {
	Validate() error
}
