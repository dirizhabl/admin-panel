package memory

import (
	"admin-panel/internal/model"
	"admin-panel/internal/store"
)

type UserRepository struct {
	users []*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	for _, realUser := range r.users {
		if u.Email == realUser.Email {
			return store.ErrRecordExists
		}
	}
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, realUser := range r.users {
		if email == realUser.Email {
			return realUser, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Update(dto *model.UserUpdate) (*model.User, error) {
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
