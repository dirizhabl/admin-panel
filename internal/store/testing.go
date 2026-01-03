package store

import (
	"math"
	"testing"

	"admin-panel/internal/model"

	"github.com/stretchr/testify/assert"
)

func toPtr[T any](arg T) *T {
	return &arg
}

func UserRepositoryCreate(
	t *testing.T,
	newStore func() (Store, func()),
) {
	t.Helper()
	testCases := []struct {
		name    string
		prepare func(UserRepository) *model.User
		wantErr error
	}{
		{
			name: "user already exists",
			prepare: func(userRepo UserRepository) *model.User {
				u := model.TestUser()
				userRepo.Create(u)
				return u
			},
			wantErr: ErrRecordExists,
		},
		{
			name: "successful create",
			prepare: func(userRepo UserRepository) *model.User {
				return model.TestUser()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()	

			userRepo := store.User()

			user := tc.prepare(userRepo)
			err := userRepo.Create(user)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func UserRepositoryFindById(
	t *testing.T,
	newStore func() (Store, func()),
) {
	t.Helper()
	testCases := []struct {
		name    string
		prepare func(UserRepository) int
		wantErr error
	}{
		{
			name: "user exists",
			prepare: func(userRepo UserRepository) int {
				u := model.TestUser()
				userRepo.Create(u)
				return u.ID
			},
		},
		{
			name: "user not found",
			prepare: func(userRepo UserRepository) int {
				return math.MaxInt64
			},
			wantErr: ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			userRepo := store.User()

			id := tc.prepare(userRepo)
			user, err := userRepo.FindById(id)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}
		})
	}
}

func UserRepositoryFindByEmail(
	t *testing.T,
	newStore func() (Store, func()),
) {
	t.Helper()
	testCases := []struct {
		name    string
		prepare func(UserRepository) string
		wantErr error
	}{
		{
			name: "user exists",
			prepare: func(userRepo UserRepository) string {
				u := model.TestUser()
				userRepo.Create(u)
				return u.Email
			},
		},
		{
			name: "user not found",
			prepare: func(userRepo UserRepository) string {
				return "nil@nil.org"
			},
			wantErr: ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			userRepo := store.User()

			email := tc.prepare(userRepo)
			user, err := userRepo.FindByEmail(email)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}
		})
	}
}

func UserRepositoryUpdate(
	t *testing.T,
	newStore func() (Store, func()),
) {
	t.Helper()
	testCases := []struct {
		name    string
		prepare func(UserRepository) *model.UserUpdate
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(userRepo UserRepository) *model.UserUpdate {
				id := math.MaxInt64
				return model.TestUserUpdate(id)
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "successful update",
			prepare: func(userRepo UserRepository) *model.UserUpdate {
				u := model.TestUser()
				userRepo.Create(u)
				dto := model.TestUserUpdate(u.ID)
				dto.FirstName = toPtr("john")
				dto.LastName = toPtr("doe")
				dto.Age = toPtr(int16(18))
				return dto
			},
		},
		{
			name: "only first_name",
			prepare: func(userRepo UserRepository) *model.UserUpdate {
				u := model.TestUser()
				userRepo.Create(u)
				dto := model.TestUserUpdate(u.ID)
				dto.FirstName = toPtr("john")
				return dto
			},
		},
		{
			name: "only last_name",
			prepare: func(userRepo UserRepository) *model.UserUpdate {
				u := model.TestUser()
				userRepo.Create(u)
				dto := model.TestUserUpdate(u.ID)
				dto.LastName = toPtr("doe")
				return dto
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			userRepo := store.User()

			dto := tc.prepare(userRepo)
			user, err := userRepo.Update(dto)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				if dto.FirstName != nil {
					assert.Equal(t, *dto.FirstName, user.FirstName)
				}
				if dto.LastName != nil {
					assert.Equal(t, *dto.LastName, user.LastName)
				}
				if dto.Age != nil {
					assert.Equal(t, *dto.Age, user.Age)
				}
			}
		})
	}
}
