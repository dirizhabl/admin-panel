package domain

type User struct {
	ID             int
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	Age            int16
}

func NewUser(email, password string) (*User, error) {
	if email == "" {
		return nil, ErrEmptyEmail
	}
	if len(password) < 7 {
		return nil, ErrWeakPassword
	}

	hash, err := hashString(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:          email,
		HashedPassword: hash,
	}, nil
}
