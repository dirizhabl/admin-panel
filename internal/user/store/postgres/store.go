package postgres

import (
	"admin-panel/internal/user/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db             *pgxpool.Pool
	userRepository store.UserRepository
}

func New(db *pgxpool.Pool) store.Store {
	return &Store{
		db: db,
		userRepository: &UserRepository{
			db: db,
		},
	}
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}
