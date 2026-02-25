package postgres

import (
	"admin-panel/internal/user/store"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	db             *sql.DB
	userRepository store.UserRepository
}

func New(db *sql.DB) store.Store {
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
