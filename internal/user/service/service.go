package service

import (
	"admin-panel/internal/user/domain"
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

func (s *Service) CreateUser(in *UserCreateIn) (UserCreateOut, error) {
	if err := in.Validate(); err != nil {
		return UserCreateOut{}, err
	}
	u, err := domain.NewUser(in.Email, in.Password)
	if err != nil {
		return UserCreateOut{}, err
	}
	if err := s.store.User().Create(u); err != nil {
		return UserCreateOut{}, err
	}

	return UserCreateOut{ID: u.ID}, nil
}

func (s *Service) FindUserById(id int) (UserReadOut, error) {
	u, err := s.store.User().FindById(id)
	if err != nil {
		return UserReadOut{}, err
	}

	return toUserReadOut(u), nil
}

func (s *Service) FindUserByEmail(email string) (UserReadOut, error) {
	u, err := s.store.User().FindByEmail(email)
	if err != nil {
		return UserReadOut{}, err
	}

	return toUserReadOut(u), nil
}

func (s *Service) UpdateUser(in *UserUpdateIn) (UserReadOut, error) {
	dto := &domain.UserUpdate{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       in.Age,
	}
	if err := dto.Validate(); err != nil {
		return UserReadOut{}, err
	}
	u, err := s.store.User().Update(dto)
	if err != nil {
		return UserReadOut{}, err
	}

	return toUserReadOut(u), nil
}
