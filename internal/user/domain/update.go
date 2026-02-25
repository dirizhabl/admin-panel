package domain

import "admin-panel/pkg/apperr"

type UserUpdate struct {
	ID        int
	FirstName *string
	LastName  *string
	Age       *int16
}

func (u *UserUpdate) Validate() error {
	if u.Age != nil && *u.Age < 0 {
		return apperr.ErrorDetail{Err: ErrNegativeAge, Field: "age"}
	}
	return nil
}
