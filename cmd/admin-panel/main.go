package main

import (
	"database/sql"
	"log"
	"net/http"

	"admin-panel/internal/config"
	adminHTTP "admin-panel/internal/transport/http/admin"
	"admin-panel/internal/transport/http/middleware"
	userService "admin-panel/internal/user/service"
	"admin-panel/internal/user/store/postgres"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.MustLoad()

	logger := logrus.New()

	db, err := newOpenDB(&config.Store)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.Use(
		middleware.Logger(logger),
		middleware.Recover(),
	)

	store := postgres.New(db)
	userService := userService.New(store)

	adminHTTP.New(logger, router, userService)

	server := http.Server{
		Addr:    config.HTTPServer.BindAddr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func newOpenDB(config *config.Store) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
