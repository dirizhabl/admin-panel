package postgres_test

import (
	"testing"

	"admin-panel/internal/store"
	"admin-panel/internal/store/postgres"
)

func TestUserRepositoryPostgres(t *testing.T) {
	db, teardown := postgres.TestDB(t, databaseURL)
	t.Cleanup(func() {
		db.Close()
	})

	t.Run("Create", func(t *testing.T) {
		store.UserRepositoryCreate(t, func() (store.Store, func()) {
			return postgres.New(db), func() {
				teardown("users")
			}
		})
	})

	t.Run("FindById", func(t *testing.T) {
		store.UserRepositoryFindById(t, func() (store.Store, func()) {
			return postgres.New(db), func() {
				teardown("users")
			}
		})
	})

	t.Run("FindByEmail", func(t *testing.T) {
		store.UserRepositoryFindByEmail(t, func() (store.Store, func()) {
			return postgres.New(db), func() {
				teardown("users")
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		store.UserRepositoryUpdate(t, func() (store.Store, func()) {
			return postgres.New(db), func() {
				teardown("users")
			}
		})
	})
}
