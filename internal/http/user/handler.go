package user

import (
	"http-rest-api/internal/service/user"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	router  *mux.Router
	service *user.Service
}

func New(logger *logrus.Logger, router *mux.Router, service *user.Service) *Handler {
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
	h.router.HandleFunc("/users/{id}", h.FindUser()).Methods("GET")
	h.router.HandleFunc("/users/{id}", h.UpdateUser()).Methods("PATCH")
}
