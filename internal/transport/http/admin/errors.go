package admin

import (
	"errors"
	"strings"
)

var (
	ErrWeakPassword      = errors.New("length of password must be more 6")
	ErrEmptyEmail        = errors.New("length of email must be more 0")
	ErrNotValidEmail     = errors.New("email is not valid")
	ErrEmptyFieldsUpdate = errors.New("no fields to update")
)

type ValidationErrors []error

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for i, err := range v {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func (v ValidationErrors) Errors() []string {
	result := make([]string, len(v))

	for i, err := range v {
		result[i] = err.Error()
	}
	return result
}
