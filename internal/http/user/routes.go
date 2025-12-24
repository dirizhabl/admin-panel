package user

import (
	"encoding/json"
	"errors"
	req "http-rest-api/internal/dto/request"
	resp "http-rest-api/internal/dto/response"
	"http-rest-api/internal/model"
	"http-rest-api/internal/store"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &req.CreateUser{}
		if err := h.BindJson(w, r, req); err != nil {
			return
		}

		u, err := h.service.CreateUser(req)

		if err != nil {
			h.Error(w, err)
			return
		}
		h.JSON(201, w, u)
	}
}

func (h *Handler) FindUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := mux.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			h.Error(w, errors.New("cannot convert to int"))
			return
		}

		u, err := h.service.FindUserById(id)
		if err != nil {
			h.Error(w, err)
			return
		}
		h.JSON(200, w, u)
	}
}

func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := mux.Vars(r)
		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			h.Error(w, errors.New("cannot convert to int"))
			return
		}

		req := &req.UpdateUser{}
		req.ID = id
		if err := h.BindJson(w, r, req); err != nil {
			return 
		}

		u, err := h.service.UpdateUser(req)
		if err != nil {
			h.Error(w, err)
			return
		}
		h.JSON(200, w, u)
	}
}

func (h *Handler) JSON(statusCode int, w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) BindJson(w http.ResponseWriter, r *http.Request, in any) error {
	if r.Body == nil {
		h.JSON(400, w, resp.Err{Err: "empty body"})
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(in); err != nil {
		h.JSON(400, w, resp.Err{Err: err.Error()})
		return err
	}
	return nil
}

func (h *Handler) Error(w http.ResponseWriter, err error) {
	var errs model.ValidationErrors

	switch {
	case errors.As(err, &errs):
		h.JSON(422, w, resp.Errs{Errs: errs.Errors()})
	case errors.Is(err, store.ErrRecordExists):
		h.JSON(409, w, resp.Err{Err: err.Error()})
	case errors.Is(err, store.ErrRecordNotFound):
		h.JSON(404, w, resp.Err{Err: err.Error()})
	default:
		h.JSON(500, w, resp.Err{Err: err.Error()})
	}
}
