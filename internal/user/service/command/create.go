package command

type UserCreate struct {
	Email    string
	Password string
}

func (u *UserCreate) Validate() error {
	if u.Email == "" || u.Password == "" {
		return ErrInvalidInput
	}
	return nil
}
