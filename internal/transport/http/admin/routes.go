package admin

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"admin-panel/internal/user"
	"admin-panel/internal/user/service"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r, h.logger)

		req := &UserCreateRequest{}
		if err := c.BindJson(req); err != nil {
			c.Error(err)
			return
		}

		in := &service.UserCreateIn{
			Email:    req.Email,
			Password: req.Password,
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		u, err := h.service.CreateUser(ctx, in)
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
		c := NewHandlerContext(w, r, h.logger)

		var query UserFiltersQuery

		if err := c.BindQueryParams(&query); err != nil {
			c.Error(err)
			return
		}

		f := &user.Filters{
			Email:     query.Email,
			FirstName: query.FirstName,
			LastName:  query.LastName,
			MinAge:    query.MinAge,
			MaxAge:    query.MaxAge,
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		users, err := h.service.FindUsersByFilters(ctx, f)
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
		c := NewHandlerContext(w, r, h.logger)
		pathVars := mux.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			c.Error(errors.New("cannot convert to int"))
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		u, err := h.service.FindUserById(ctx, id)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(200, toUserReadResponse(&u))
	}
}

func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewHandlerContext(w, r, h.logger)
		pathVars := mux.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			c.Error(errors.New("cannot convert to int"))
			return
		}

		req := &UserUpdateRequest{}
		if err := c.BindJson(req); err != nil {
			c.Error(err)
			return
		}

		in := &service.UserUpdateIn{
			ID:        id,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Age:       req.Age,
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		u, err := h.service.UpdateUser(ctx, in)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(200, toUserReadResponse(&u))
	}
}
