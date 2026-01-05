package domain

import "errors"

var (
	ErrWeakPassword      = errors.New("length of password must be more 6")
	ErrEmptyEmail        = errors.New("length of email must be more 0")
	ErrEmptyFieldsUpdate = errors.New("no fields to update")
)
