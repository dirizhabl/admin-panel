package admin

import "errors"

var (
	ErrWeakPassword      = errors.New("length of password must be more 6")
	ErrEmptyEmail        = errors.New("length of email must be more 0")
	ErrNotValidEmail     = errors.New("email is not valid")
	ErrMinAge            = errors.New("min age must be more 0")
	ErrMaxAge            = errors.New("max age must be less 255")
	ErrMinMaxAge         = errors.New("max age must be more min age")
	ErrNegativeAge       = errors.New("age must be more 0")
)
