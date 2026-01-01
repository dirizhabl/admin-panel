package postgres

import (
	"admin-panel/internal/store"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db             *sqlx.DB
	userRepository store.UserRepository
}

func New(db *sqlx.DB) store.Store {
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
