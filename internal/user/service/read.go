package service

import "admin-panel/internal/user/domain"

type UserReadOut struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Age       int16
}

func toUserReadOut(u *domain.User) UserReadOut {
	return UserReadOut{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
}
