package user_test

import (
	"math"
	"testing"

	"admin-panel/internal/dto/request"
	"admin-panel/internal/model"
	"admin-panel/internal/service/user"
	"admin-panel/internal/store"
	"admin-panel/internal/store/memory"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	memoryStoreFabric := func() store.Store {
		return memory.New()
	}
	t.Run("CreateUser", func(t *testing.T) {
		CreateUser(t, memoryStoreFabric)
	})
	t.Run("FindUserById", func(t *testing.T) {
		FindUserById(t, memoryStoreFabric)
	})
	t.Run("FindUserByEmail", func(t *testing.T) {
		FindUserByEmail(t, memoryStoreFabric)
	})
	t.Run("UpdateUser", func(t *testing.T) {
		UpdateUser(t, memoryStoreFabric)
	})
}

func CreateUser(
	t *testing.T,
	newStore func() store.Store,
) {
	testCases := []struct {
		name       string
		prepareDTO func(store.Store) *request.CreateUser
		wantErr    error
	}{
		{
			name: "empty email",
			prepareDTO: func(s store.Store) *request.CreateUser {
				return &request.CreateUser{
					Email:    "",
					Password: "123",
				}
			},
			wantErr: model.ErrEmptyEmail,
		},
		{
			name: "empty password",
			prepareDTO: func(s store.Store) *request.CreateUser {
				return &request.CreateUser{
					Email:    "nil@nil.org",
					Password: "",
				}
			},
			wantErr: model.ErrEmptyPassword,
		},
		{
			name: "user already exists",
			prepareDTO: func(s store.Store) *request.CreateUser {
				u := model.TestUser()
				s.User().Create(u)

				return &request.CreateUser{
					Email:    u.Email,
					Password: "123",
				}
			},
			wantErr: store.ErrRecordExists,
		},
		{
			name: "ok",
			prepareDTO: func(s store.Store) *request.CreateUser {
				return &request.CreateUser{
					Email:    "nil@nil.org",
					Password: "123",
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()
			dto := tc.prepareDTO(store)
			userService := user.New(store)

			user, err := userService.CreateUser(dto)

			if tc.wantErr != nil {
				assert.Nil(t, user)
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NotNil(t, user)
				assert.NoError(t, err)
			}
		})
	}
}

func FindUserById(
	t *testing.T,
	newStore func() store.Store,
) {
	testCases := []struct {
		name      string
		prepareId func(store.Store) int
		wantErr   error
	}{
		{
			name: "user not found",
			prepareId: func(s store.Store) int {
				return math.MaxInt64
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "ok",
			prepareId: func(s store.Store) int {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()
			id := tc.prepareId(store)
			userService := user.New(store)

			user, err := userService.FindUserById(id)

			if tc.wantErr != nil {
				assert.Nil(t, user)
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NotNil(t, user)
				assert.NoError(t, err)
			}
		})
	}
}

func FindUserByEmail(
	t *testing.T,
	newStore func() store.Store,
) {
	testCases := []struct {
		name         string
		prepareEmail func(store.Store) string
		wantErr      error
	}{
		{
			name: "user not found",
			prepareEmail: func(s store.Store) string {
				u := model.TestUser()
				return u.Email
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "ok",
			prepareEmail: func(s store.Store) string {
				u := model.TestUser()
				s.User().Create(u)
				return u.Email
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()
			id := tc.prepareEmail(store)
			userService := user.New(store)

			user, err := userService.FindUserByEmail(id)

			if tc.wantErr != nil {
				assert.Nil(t, user)
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NotNil(t, user)
				assert.NoError(t, err)
			}
		})
	}
}

func toPtr[T any](arg T) *T {
	return &arg
}

func UpdateUser(
	t *testing.T,
	newStore func() store.Store,
) {
	t.Helper()
	testCases := []struct {
		name    string
		prepare func(store.Store) (int, *request.UserUpdate)
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				id := math.MaxInt64
				return id, &request.UserUpdate{
					FirstName: toPtr("john"),
					LastName:  toPtr("doe"),
					Age:       toPtr(int16(18)),
				}
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "no data",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID, &request.UserUpdate{}
			},
			wantErr: model.ErrEmptyFieldsUpdate,
		},
		{
			name: "ok",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID, &request.UserUpdate{
					FirstName: toPtr("john"),
					LastName:  toPtr("doe"),
					Age:       toPtr(int16(18)),
				}
			},
		},
		{
			name: "ok/only first_name",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID, &request.UserUpdate{
					FirstName: toPtr("john"),
				}
			},
		},
		{
			name: "ok/only last_name",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID, &request.UserUpdate{
					LastName: toPtr("doe"),
				}
			},
		},
		{
			name: "ok/only age",
			prepare: func(s store.Store) (int, *request.UserUpdate) {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID, &request.UserUpdate{
					Age: toPtr(int16(18)),
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()
			userService := user.New(store)

			id, dto := tc.prepare(store)
			user, err := userService.UpdateUser(id, dto)

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
