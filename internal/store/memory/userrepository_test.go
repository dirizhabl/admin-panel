package memory_test

import (
	"admin-panel/internal/model"
	"admin-panel/internal/store"
	"admin-panel/internal/store/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := memory.New()
	u := model.TestUser()
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := memory.New()
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.ErrorIs(t, err, store.ErrRecordNotFound)

	u := model.TestUser()
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
