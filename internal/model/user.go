package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	Age            int16
}

type UserUpdate struct {
	ID        int
	FirstName *string
	LastName  *string
	Age       *int16
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrEmptyEmail
	}
	if u.HashedPassword == "" {
		return ErrEmptyPassword
	}
	return nil
}

func (u *User) BeforeCreate(password string) error {
	hash, err := hashString(password)
	if err != nil {
		return err
	}
	u.HashedPassword = hash
	return nil
}

func hashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
