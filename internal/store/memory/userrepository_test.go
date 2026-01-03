package memory_test

import (
	"testing"

	"admin-panel/internal/store"
	"admin-panel/internal/store/memory"
)

func TestUserRepositoryMemory(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		store.UserRepositoryCreate(t, func() (store.Store, func()) {
			return memory.New(), func() {}
		})
	})

	t.Run("FindById", func(t *testing.T) {
		store.UserRepositoryFindById(t, func() (store.Store, func()) {
			return memory.New(), func() {}
		})
	})

	t.Run("FindByEmail", func(t *testing.T) {
		store.UserRepositoryFindByEmail(t, func() (store.Store, func()) {
			return memory.New(), func() {}
		})
	})

	t.Run("Update", func(t *testing.T) {
		store.UserRepositoryUpdate(t, func() (store.Store, func()) {
			return memory.New(), func() {}
		})
	})
}
