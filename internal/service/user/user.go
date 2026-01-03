package user

import (
	req "admin-panel/internal/dto/request"
	resp "admin-panel/internal/dto/response"
	"admin-panel/internal/model"
	"admin-panel/internal/store"
)

type Service struct {
	store store.Store
}

func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateUser(req *req.CreateUser) (*resp.CreateUser, error) {
	u := &model.User{
		Email: req.Email,
	}
	if err := u.BeforeCreate(req.Password); err != nil {
		return nil, err
	}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.User().Create(u); err != nil {
		return nil, err
	}

	return &resp.CreateUser{ID: u.ID}, nil
}

func (s *Service) FindUserById(id int) (*resp.GetUser, error) {
	u, err := s.store.User().FindById(id)
	if err != nil {
		return nil, err
	}

	return &resp.GetUser{ID: u.ID}, nil
}

func (s *Service) FindUserByEmail(email string) (*resp.GetUser, error) {
	u, err := s.store.User().FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return &resp.GetUser{ID: u.ID}, nil
}

func (s *Service) UpdateUser(id int, req *req.UserUpdate) (*resp.GetUser, error) {
	dto := &model.UserUpdate{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	u, err := s.store.User().Update(dto)
	if err != nil {
		return nil, err
	}

	return &resp.GetUser{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}, nil
}
