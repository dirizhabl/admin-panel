package service

import (
	"admin-panel/internal/user"
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/store"
	"admin-panel/pkg/apperr"
	"errors"
)

type Service struct {
	store store.Store
}

func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateUser(in *UserCreateIn) (UserCreateOut, error) {
	u, err := domain.NewUser(in.Email, in.Password)
	if err != nil {
		return UserCreateOut{}, apperr.NewValidation(err, err.Error())
	}
	if err := s.store.User().Create(u); err != nil {
		if errors.Is(err, store.ErrRecordExists) {
			return UserCreateOut{}, apperr.NewConflict(err, err.Error())
		}
		return UserCreateOut{}, apperr.NewInternal(err)
	}

	return UserCreateOut{ID: u.ID}, nil
}

func (s *Service) FindUserById(id int) (UserReadOut, error) {
	u, err := s.store.User().FindById(id)
	if err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return UserReadOut{}, apperr.NewNotFound(err, err.Error())
		}
		return UserReadOut{}, apperr.NewInternal(err)
	}

	return toUserReadOut(u), nil
}

func (s *Service) FindUserByEmail(email string) (UserReadOut, error) {
	u, err := s.store.User().FindByEmail(email)
	if err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return UserReadOut{}, apperr.NewNotFound(err, err.Error())
		}
		return UserReadOut{}, apperr.NewInternal(err)
	}

	return toUserReadOut(u), nil
}

func (s *Service) FindUsersByFilters(f *user.Filters) ([]UserReadOut, error) {
	if err := f.Validate(); err != nil {
		return []UserReadOut{}, apperr.NewValidation(err, err.Error())
	}
	users, err := s.store.User().FindByFilters(f)
	if err != nil {
		return []UserReadOut{}, apperr.NewInternal(err)
	}
	out := make([]UserReadOut, len(users))
	for i, u := range users {
		out[i] = toUserReadOut(u)
	}

	return out, nil
}

func (s *Service) UpdateUser(in *UserUpdateIn) (UserReadOut, error) {
	if err := in.IsEmpty(); err != nil {
		return UserReadOut{}, apperr.NewValidation(err, err.Error())
	}
	dto := &domain.UserUpdate{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       in.Age,
	}
	if err := dto.Validate(); err != nil {
		return UserReadOut{}, apperr.NewValidation(err, err.Error())
	}
	u, err := s.store.User().Update(dto)
	if err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return UserReadOut{}, apperr.NewNotFound(err, err.Error())
		}
		return UserReadOut{}, apperr.NewInternal(err)
	}

	return toUserReadOut(u), nil
}
