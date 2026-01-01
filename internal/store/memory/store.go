package memory

import (
	"admin-panel/internal/model"
	"admin-panel/internal/store"
)

type Store struct {
	user store.UserRepository
}

func New() *Store {
	return &Store{
		user: &UserRepository{
			users: make(map[string]*model.User),
		},
	}
}

func (s *Store) User() store.UserRepository {
	return s.user
}
