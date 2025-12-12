package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrRecordExists   = errors.New("record already exists")
)
