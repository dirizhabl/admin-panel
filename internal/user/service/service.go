package service

import (
	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/service/command"
	"admin-panel/internal/user/service/out"
	"admin-panel/internal/user/store"
)

type Service struct {
	store store.Store
}

func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateUser(in *command.UserCreate) (out.UserCreate, error) {
	if err := in.Validate(); err != nil {
		return out.UserCreate{}, err
	}
	u, err := domain.NewUser(in.Email, in.Password)
	if err != nil {
		return out.UserCreate{}, err
	}
	if err := s.store.User().Create(u); err != nil {
		return out.UserCreate{}, err
	}

	return out.UserCreate{ID: u.ID}, nil
}

func (s *Service) FindUserById(id int) (out.UserRead, error) {
	u, err := s.store.User().FindById(id)
	if err != nil {
		return out.UserRead{}, err
	}

	return out.UserRead{ID: u.ID}, nil
}

func (s *Service) FindUserByEmail(email string) (out.UserRead, error) {
	u, err := s.store.User().FindByEmail(email)
	if err != nil {
		return out.UserRead{}, err
	}

	return out.UserRead{ID: u.ID}, nil
}

func (s *Service) UpdateUser(in *command.UserUpdate) (out.UserRead, error) {
	dto := &domain.UserUpdate{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       in.Age,
	}
	if err := dto.Validate(); err != nil {
		return out.UserRead{}, err
	}
	u, err := s.store.User().Update(dto)
	if err != nil {
		return out.UserRead{}, err
	}

	return out.UserRead{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}, nil
}
