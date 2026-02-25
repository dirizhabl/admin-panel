package store

import (
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrRecordExists    = errors.New("record already exists")
	ErrNothingToUpdate = errors.New("nothing to update")
)

type OpError struct {
	Err error
	Op  string
}

func NewOpError(op string, err error) *OpError {
	return &OpError{
		Err: err,
		Op:  op,
	}
}

func (e *OpError) Error() string {
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *OpError) Unwrap() error {
	return e.Err
}
