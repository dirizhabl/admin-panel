package main

import (
	"database/sql"
	"http-rest-api/internal/config"
	"http-rest-api/internal/http/middleware"
	userHTTP "http-rest-api/internal/http/user"
	userService "http-rest-api/internal/service/user"

	"http-rest-api/internal/store/postgres"
	"log"
	"net/http"

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
	wrappedRouter := middleware.Recover()(router)
	wrappedRouter = middleware.Logger(logger)(wrappedRouter)

	store := postgres.New(db)
	userService := userService.New(store)

	userHTTP.New(logger, router, userService)

	server := http.Server{
		Addr:    config.HTTPServer.BindAddr,
		Handler: wrappedRouter,
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
