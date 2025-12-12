package user

import (
	"fmt"
	"http-rest-api/internal/service/user"
	"net/http"

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
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OKs")
	h.router.ServeHTTP(w, r)
}
