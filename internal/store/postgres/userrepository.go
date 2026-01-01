package postgres

import (
	"database/sql"
	"errors"
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"

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
		u.HashedPassword,
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
		&u.HashedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Update(dto *model.UserUpdate) (*model.User, error) {
	args := []any{
		dto.FirstName != nil, dto.FirstName,
		dto.LastName != nil, dto.LastName,
		dto.Age != nil, dto.Age,
		dto.ID,
	}
	var userRow userRow
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
		&userRow.ID, &userRow.Email, &userRow.FirstName, &userRow.LastName, &userRow.Age,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	user := &model.User{
		ID:    userRow.ID,
		Email: userRow.Email,
	}

	if userRow.FirstName != nil {
		user.FirstName = *userRow.FirstName
	}

	if userRow.LastName != nil {
		user.LastName = *userRow.LastName
	}

	if userRow.Age != nil {
		user.Age = *userRow.Age
	}

	return user, nil
}
