package postgres

import "admin-panel/internal/user/domain"

type userRow struct {
	ID        int
	Email     string
	FirstName *string
	LastName  *string
	Age       *int16
}

func (r *userRow) toUser() *domain.User {
	return &domain.User{
		ID:        r.ID,
		Email:     r.Email,
		FirstName: fromPtr(r.FirstName),
		LastName:  fromPtr(r.LastName),
		Age:       fromPtr(r.Age),
	}
}

func fromPtr[T any](arg *T) T {
	var zero T
	if arg == nil {
		return zero
	}
	return *arg
}
