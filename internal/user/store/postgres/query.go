package postgres

import (
	"admin-panel/internal/user"

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