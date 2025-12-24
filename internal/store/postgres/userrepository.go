package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r *UserRepository) Create(u *model.User) error {
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

func (r *UserRepository) FindById(id int) (*model.User, error) {
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

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
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

func (r *UserRepository) Update(u *model.User) error {
	setParts := []string{}
	args := map[string]any{"id": u.ID}
	if u.FirstName != "" {
		setParts = append(setParts, "first_name = :first_name")
		args["first_name"] = u.FirstName
	}
	if u.LastName != "" {
		setParts = append(setParts, "last_name = :last_name")
		args["last_name"] = u.LastName
	}
	if u.Age != 0 {
		setParts = append(setParts, "age = :age")
		args["age"] = u.Age
	}
	if len(setParts) == 0 {
		return errors.New("no data")
	}

	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE id = :id RETURNING email", 
		strings.Join(setParts, ", "),
	)

	namedQuery, namedArgs, err := sqlx.Named(query, args)
	if err != nil {
		return err
	}

	namedQuery = r.db.Rebind(namedQuery)

	err = r.db.QueryRow(namedQuery, namedArgs...).Scan(&u.Email)
	if err != nil {
		return err
	}
	return nil
}
