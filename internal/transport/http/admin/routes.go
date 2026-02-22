package admin

import (
	"errors"
	"net/http"
	"strconv"

	"admin-panel/internal/user"
	"admin-panel/internal/user/service"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)

		req := &UserCreateRequest{}
		if err := c.BindJson(req); err != nil {
			return
		}

		in := &service.UserCreateIn{
			Email:    req.Email,
			Password: req.Password,
		}

		u, err := h.service.CreateUser(in)
		if err != nil {
			c.Error(err)
			return
		}

		response := &UserCreateResponse{
			ID: u.ID,
		}
		c.JSON(201, response)
	}
}


func (h *Handler) FindUsersByParams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)

		var query UserFiltersQuery

		if err := c.BindQueryParams(&query); err != nil {
			return
		}

		f := &user.Filters{
			Email:     query.Email,
			FirstName: query.FirstName,
			LastName:  query.LastName,
			MinAge:    query.MinAge,
			MaxAge:    query.MaxAge,
		}

		users, err := h.service.FindUsersByFilters(f)
		if err != nil {
			c.Error(err)
			return
		}

		usersResponse := make([]UserReadResponse, len(users))
		for i, u := range users {
			usersResponse[i] = toUserReadResponse(&u)
		}

		c.JSON(200, usersResponse)
	}
}

func (h *Handler) FindUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)
		pathVars := mux.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			c.Error(errors.New("cannot convert to int"))
			return
		}

		u, err := h.service.FindUserById(id)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(200, toUserReadResponse(&u))
	}
}

func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)
		pathVars := mux.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			c.Error(errors.New("cannot convert to int"))
			return
		}

		req := &UserUpdateRequest{}
		if err := c.BindJson(req); err != nil {
			return
		}

		in := &service.UserUpdateIn{
			ID:        id,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Age:       req.Age,
		}

		u, err := h.service.UpdateUser(in)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(200, toUserReadResponse(&u))
	}
}
