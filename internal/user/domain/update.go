package domain

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
