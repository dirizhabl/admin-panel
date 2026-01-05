package service

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
)

type UserCreateIn struct {
	Email    string
	Password string
}

func (u *UserCreateIn) Validate() error {
	if u.Email == "" || u.Password == "" {
		return ErrInvalidInput
	}
	return nil
}

type UserCreateOut struct {
	ID int
}
