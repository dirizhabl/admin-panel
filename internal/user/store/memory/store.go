package memory

import (
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"
)

type Store struct {
	user store.UserRepository
}

func New() *Store {
	return &Store{
		user: &UserRepository{
			users: make([]*domain.User, 0),
		},
	}
}

func (s *Store) User() store.UserRepository {
	return s.user
}
