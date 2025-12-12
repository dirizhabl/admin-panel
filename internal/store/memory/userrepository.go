package memory

import (
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"
)

type userRepository struct {
	users map[string]*model.User
}

func (r *userRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

// mock
func (r *userRepository) FindById(id int) (*model.User, error) {
	return &model.User{}, nil
}