package service

type UserUpdateIn struct {
	ID        int
	FirstName *string
	LastName  *string
	Age       *int16
}

func (in *UserUpdateIn) IsEmpty() error {
	if in.FirstName == nil && in.LastName == nil && in.Age == nil {
		return ErrEmptyFieldsUpdate
	}
	return nil
}
