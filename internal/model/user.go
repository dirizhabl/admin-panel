package model

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int    `json:"id"`
	Email           string `json:"-"`
	Password        string `json:"-"`
	Hashed_password string `json:"-"`
}

var (
	ErrEmptyPassword = errors.New("length of password must be more 0")
	ErrEmptyEmail    = errors.New("length of email must be more 0")
	ErrNotValidEmail = errors.New("email must contain \"@\"")
)

func (u *User) Validate() error {
	var errs []error
	if u.Password == "" && u.Hashed_password == "" {
		errs = append(errs, ErrEmptyPassword)
	}
	if u.Email == "" {
		errs = append(errs, ErrEmptyEmail)
	}
	if !strings.ContainsRune(u.Email, '@') {
		errs = append(errs, ErrNotValidEmail)
	}
	if len(errs) > 0 {
		return ValidationErrors(errs)
	}
	return nil
}

func (u *User) BeforeCreate() error {
	if u.Password != "" {
		hash, err := hashString(u.Password)
		if err != nil {
			return err
		}
		u.Hashed_password = hash
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func hashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
