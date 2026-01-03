package main

import (
	"log"
	"net/http"

	"admin-panel/internal/config"
	"admin-panel/internal/http/middleware"
	userHTTP "admin-panel/internal/http/user"
	userService "admin-panel/internal/service/user"
	"admin-panel/internal/store/postgres"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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

	userHTTP.New(logger, router, userService)

	server := http.Server{
		Addr:    config.HTTPServer.BindAddr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func newOpenDB(config *config.Store) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
