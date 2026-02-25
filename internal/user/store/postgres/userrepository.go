package postgres

import (
	"database/sql"
	"errors"

	"admin-panel/internal/user"
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *domain.User) error {
	const op = "UserRepository.Create"
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
		return store.NewOpError(op, err)
	}
	return nil
}

func (r *UserRepository) FindById(id int) (*domain.User, error) {
	const op = "UserRepository.FindById"
	var row userRow
	if err := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, age FROM users WHERE id = $1", id,
	).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, store.NewOpError(op, err)
	}
	return row.toUser(), nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	const op = "UserRepository.FindByEmail"
	var row userRow
	if err := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, age FROM users WHERE email = $1", email,
	).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, store.NewOpError(op, err)
	}
	return row.toUser(), nil
}

func (r *UserRepository) FindByFilters(f *user.Filters) ([]*domain.User, error) {
	const op = "UserRepository.FindByFilters"
	query, args, err := FindByFiltersQuery(f)
	if err != nil {
		return nil, store.NewOpError(op, err)
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, store.NewOpError(op, err)
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		var row userRow
		if err := rows.Scan(
			&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
		); err != nil {
			return nil, store.NewOpError(op, err)
		}
		users = append(users, row.toUser())
	}
	if err := rows.Err(); err != nil {
		return nil, store.NewOpError(op, err)
	}
	return users, nil
}

func (r *UserRepository) Update(dto *domain.UserUpdate) (*domain.User, error) {
	const op = "UserRepository.Update"
	query, args, err := UpdateQuery(dto)
	if err != nil {
		return nil, store.NewOpError(op, err)
	}
	var row userRow
	if err := r.db.QueryRow(query, args...).Scan(
		&row.ID, &row.Email, &row.FirstName, &row.LastName, &row.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, store.NewOpError(op, err)
	}

	return row.toUser(), nil
}
