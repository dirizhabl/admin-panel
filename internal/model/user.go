package model

type User struct {
	ID             int
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	Age            int16
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrEmptyEmail
	}
	return nil
}

func (u *User) BeforeCreate(password string) error {
	if password == "" {
		return ErrEmptyPassword
	}
	hash, err := hashString(password)
	if err != nil {
		return err
	}
	u.HashedPassword = hash
	return nil
}

type UserUpdate struct {
	ID        int
	FirstName *string
	LastName  *string
	Age       *int16
}

func (u *UserUpdate) Validate() error {
	if u.FirstName == nil && u.LastName == nil && u.Age == nil {
		return ErrEmptyFieldsUpdate
	}
	return nil
}
