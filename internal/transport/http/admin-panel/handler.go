package user

import (
	"admin-panel/internal/user/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	router  *mux.Router
	service *service.Service
}

func New(logger *logrus.Logger, router *mux.Router, service *service.Service) *Handler {
	h := &Handler{
		logger:  logger,
		router:  router,
		service: service,
	}
	h.registerRoutes()
	return h
}

func (h *Handler) registerRoutes() {
	h.router.HandleFunc("/users", h.CreateUser()).Methods("POST")
	h.router.HandleFunc("/users", h.FindUserByEmail()).Methods("GET")
	h.router.HandleFunc("/users/{id}", h.FindUserById()).Methods("GET")
	h.router.HandleFunc("/users/{id}", h.UpdateUser()).Methods("PATCH")
}
