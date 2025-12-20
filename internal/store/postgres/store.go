package postgres

import (
	"database/sql"
	"http-rest-api/internal/store"

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
