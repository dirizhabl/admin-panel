package model

import "errors"

var (
	ErrEmptyPassword = errors.New("length of password must be more 0")
	ErrEmptyEmail    = errors.New("length of email must be more 0")
)
