package admin

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) FindUserByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)
		values := r.URL.Query()
		email := values.Get("email")

		u, err := h.service.FindUserByEmail(email)
		if err != nil {
			c.Error(err)
			return
		}

		response := &UserReadResponse{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
		}
		c.JSON(200, response)
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

		response := &UserReadResponse{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
		}
		c.JSON(200, response)
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

		response := &UserReadResponse{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
		}
		c.JSON(200, response)
	}
}
