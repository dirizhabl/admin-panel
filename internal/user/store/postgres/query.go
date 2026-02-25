package postgres

import (
	"admin-panel/internal/user"
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"

	sq "github.com/Masterminds/squirrel"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func FindByFiltersQuery(f *user.Filters) (string, []any, error) {
	builder := psql.Select("id", "email", "first_name", "last_name", "age").
		From("users")

	if f.Email != nil {
		builder = builder.Where(sq.Eq{"email": *f.Email})
	}
	if f.FirstName != nil {
		builder = builder.Where(sq.Eq{"first_name": *f.FirstName})
	}
	if f.LastName != nil {
		builder = builder.Where(sq.Eq{"last_name": *f.LastName})
	}
	if f.MinAge != nil {
		builder = builder.Where(sq.GtOrEq{"age": *f.MinAge})
	}
	if f.MaxAge != nil {
		builder = builder.Where(sq.LtOrEq{"age": *f.MaxAge})
	}
	return builder.OrderBy("id").ToSql()
}

func UpdateQuery(dto *domain.UserUpdate) (string, []any, error) {
	builder := psql.Update("users")

	if !(dto.FirstName != nil || dto.LastName != nil || dto.Age != nil) {
		return "", nil, store.ErrNothingToUpdate
	}
	if dto.FirstName != nil {
		builder = builder.Set("first_name", *dto.FirstName)
	}
	if dto.LastName != nil {
		builder = builder.Set("last_name", *dto.LastName)
	}
	if dto.Age != nil {
		builder = builder.Set("age", *dto.Age)
	}
	return builder.
		Where(sq.Eq{"id": dto.ID}).
		Suffix("RETURNING id, email, first_name, last_name, age;").
		ToSql()
}
