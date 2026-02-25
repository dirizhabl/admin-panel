package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"admin-panel/internal/config"
	adminHTTP "admin-panel/internal/transport/http/admin"
	"admin-panel/internal/transport/http/middleware"
	userService "admin-panel/internal/user/service"
	"admin-panel/internal/user/store/postgres"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.MustLoad()

	logger := logrus.New()

	dsn := &config.Store.DatabaseURL

	db, err := newOpenDB(*dsn)
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

func newOpenDB(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10
	config.MaxConns = 2
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
