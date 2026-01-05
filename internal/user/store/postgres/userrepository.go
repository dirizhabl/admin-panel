package postgres

import (
	"database/sql"
	"errors"

	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r *UserRepository) Create(u *domain.User) error {
	err := r.db.QueryRow(
		"INSERT INTO users(email, hashed_password) VALUES ($1, $2) RETURNING id, email",
		u.Email,
		u.HashedPassword,
	).Scan(&u.ID, &u.Email)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return store.ErrRecordExists
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id int) (*domain.User, error) {
	var row userRow
	if err := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, age FROM users WHERE id = $1", id,
	).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return row.toUser(), nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var row userRow
	if err := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, age FROM users WHERE email = $1", email,
	).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return row.toUser(), nil
}

func (r *UserRepository) Update(dto *domain.UserUpdate) (*domain.User, error) {
	args := []any{
		dto.FirstName != nil, dto.FirstName,
		dto.LastName != nil, dto.LastName,
		dto.Age != nil, dto.Age,
		dto.ID,
	}
	var row userRow
	const query = `
		UPDATE users
        SET
            first_name = CASE WHEN $1 THEN $2 ELSE first_name END,
            last_name  = CASE WHEN $3 THEN $4 ELSE last_name END,
            age        = CASE WHEN $5 THEN $6 ELSE age END
        WHERE id = $7
        RETURNING id, email, first_name, last_name, age;
	`
	if err := r.db.QueryRow(query, args...).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return row.toUser(), nil
}
