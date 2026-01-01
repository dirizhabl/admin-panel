package memory

import (
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"
)

type UserRepository struct {
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

// mock
func (r *UserRepository) FindById(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (r *UserRepository) Update(upd *model.UserUpdate) (*model.User, error) {
	return &model.User{}, nil
}
