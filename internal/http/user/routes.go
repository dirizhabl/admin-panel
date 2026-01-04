package user

import (
	"errors"
	"net/http"
	"strconv"

	req "admin-panel/internal/dto/request"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)

		req := &req.CreateUser{}
		if err := c.BindJson(req); err != nil {
			return
		}

		u, err := h.service.CreateUser(req)

		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(201, u)
	}
}

func (h *Handler) FindUserByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r)
		values := r.URL.Query()
		email := values.Get("email")

		user, err := h.service.FindUserByEmail(email)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, user)
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
		c.JSON(200, u)
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

		req := &req.UserUpdate{}
		if err := c.BindJson(req); err != nil {
			return
		}

		u, err := h.service.UpdateUser(id, req)

		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, u)
	}
}
