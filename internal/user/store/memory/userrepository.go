package memory

import (
	"context"

	"admin-panel/internal/user"
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"
)

type UserRepository struct {
	users []*domain.User
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	for _, realUser := range r.users {
		if u.Email == realUser.Email {
			return store.ErrRecordExists
		}
	}
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, realUser := range r.users {
		if email == realUser.Email {
			return realUser, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindByFilters(ctx context.Context, f *user.Filters) ([]*domain.User, error) {
	return []*domain.User{}, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int) (*domain.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Update(ctx context.Context, dto *domain.UserUpdate) (*domain.User, error) {
	for _, realUser := range r.users {
		if dto.ID == realUser.ID {
			if dto.FirstName != nil {
				realUser.FirstName = *dto.FirstName
			}
			if dto.LastName != nil {
				realUser.LastName = *dto.LastName
			}
			if dto.Age != nil {
				realUser.Age = *dto.Age
			}

			return realUser, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
