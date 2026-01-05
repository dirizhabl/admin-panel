package store

import (
	"admin-panel/internal/user/domain"
)

type UserRepository interface {
	Create(*domain.User) error
	FindById(int) (*domain.User, error)
	FindByEmail(string) (*domain.User, error)
	Update(*domain.UserUpdate) (*domain.User, error)
}
