package store

import (
	"context"

	"admin-panel/internal/user"
	"admin-panel/internal/user/domain"
)

type UserRepository interface {
	Create(context.Context, *domain.User) error
	FindById(context.Context, int) (*domain.User, error)
	FindByEmail(context.Context, string) (*domain.User, error)
	FindByFilters(context.Context, *user.Filters) ([]*domain.User, error)
	Update(context.Context, *domain.UserUpdate) (*domain.User, error)
}
