package postgres

import (
	"database/sql"
	"errors"
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"

	"github.com/jackc/pgx/v5/pgconn"
)

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) Create(u *model.User) error {
	err := r.db.QueryRow(
		"INSERT INTO users(email, hashed_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.Hashed_password,
	).Scan(&u.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return store.ErrRecordExists
		}
		return err
	}
	return nil
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.db.QueryRow(
		"SELECT id, email FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.db.QueryRow(
		"SELECT id, email, hashed_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Hashed_password,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
