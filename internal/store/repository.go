package store

import "admin-panel/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindById(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	Update(*model.UserUpdate) (*model.User, error)
}
