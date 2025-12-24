package user

import (
	req "http-rest-api/internal/dto/request"
	resp "http-rest-api/internal/dto/response"
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"
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
		Email:    req.Email,
		Password: req.Password,
	}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := s.store.User().Create(u); err != nil {
		return nil, err
	}

	u.Sanitize()
	return &resp.CreateUser{ID: u.ID}, nil
}

func (s *Service) FindUserById(id int) (*resp.GetUser, error) {
	u, err := s.store.User().FindById(id)
	if err != nil {
		return nil, err
	}

	u.Sanitize()

	return &resp.GetUser{ID: u.ID}, nil
}

func (s *Service) UpdateUser(req *req.UpdateUser) (*resp.GetUser, error) {
	u := &model.User{
		ID: req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}
	err := s.store.User().Update(u)
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
