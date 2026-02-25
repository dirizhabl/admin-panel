package service_test

import (
	"context"
	"math"
	"testing"

	"admin-panel/internal/user/domain"
	"admin-panel/internal/user/service"
	"admin-panel/internal/user/store"
	"admin-panel/internal/user/store/memory"

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
		prepareDTO func(context.Context, store.Store) *service.UserCreateIn
		wantErr    error
	}{
		{
			name: "empty email",
			prepareDTO: func(ctx context.Context, s store.Store) *service.UserCreateIn {
				return &service.UserCreateIn{
					Password: "123",
				}
			},
			wantErr: domain.ErrEmptyEmail,
		},
		{
			name: "weak password",
			prepareDTO: func(ctx context.Context, s store.Store) *service.UserCreateIn {
				return &service.UserCreateIn{
					Email:    "nil@nil.org",
					Password: "123",
				}
			},
			wantErr: domain.ErrWeakPassword,
		},
		{
			name: "user already exists",
			prepareDTO: func(ctx context.Context, s store.Store) *service.UserCreateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)

				return &service.UserCreateIn{
					Email:    u.Email,
					Password: "1231231",
				}
			},
			wantErr: store.ErrRecordExists,
		},
		{
			name: "ok",
			prepareDTO: func(ctx context.Context, s store.Store) *service.UserCreateIn {
				return &service.UserCreateIn{
					Email:    "nil@nil.org",
					Password: "1231231",
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()

			ctx := context.Background()
			dto := tc.prepareDTO(ctx, store)
			userService := service.New(store)

			user, err := userService.CreateUser(ctx, dto)

			if tc.wantErr != nil {
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
		prepareId func(context.Context, store.Store) int
		wantErr   error
	}{
		{
			name: "user not found",
			prepareId: func(ctx context.Context, s store.Store) int {
				return math.MaxInt64
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "ok",
			prepareId: func(ctx context.Context, s store.Store) int {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return u.ID
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()

			ctx := context.Background()
			id := tc.prepareId(ctx, store)
			userService := service.New(store)

			user, err := userService.FindUserById(ctx, id)

			if tc.wantErr != nil {
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
		prepareEmail func(context.Context, store.Store) string
		wantErr      error
	}{
		{
			name: "user not found",
			prepareEmail: func(ctx context.Context, s store.Store) string {
				u := domain.TestUser()
				return u.Email
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "ok",
			prepareEmail: func(ctx context.Context, s store.Store) string {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return u.Email
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()

			ctx := context.Background()
			id := tc.prepareEmail(ctx, store)
			userService := service.New(store)

			user, err := userService.FindUserByEmail(ctx, id)

			if tc.wantErr != nil {
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
		prepare func(context.Context, store.Store) *service.UserUpdateIn
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				id := math.MaxInt64
				return &service.UserUpdateIn{
					ID:        id,
					FirstName: toPtr("john"),
					LastName:  toPtr("doe"),
					Age:       toPtr(int16(18)),
				}
			},
			wantErr: store.ErrRecordNotFound,
		},
		{
			name: "no data",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &service.UserUpdateIn{}
			},
			wantErr: service.ErrEmptyFieldsUpdate,
		},
		{
			name: "ok",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &service.UserUpdateIn{
					ID:        u.ID,
					FirstName: toPtr("john"),
					LastName:  toPtr("doe"),
					Age:       toPtr(int16(18)),
				}
			},
		},
		{
			name: "ok/only first_name",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &service.UserUpdateIn{
					ID:        u.ID,
					FirstName: toPtr("john"),
				}
			},
		},
		{
			name: "ok/only last_name",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &service.UserUpdateIn{
					ID:       u.ID,
					LastName: toPtr("doe"),
				}
			},
		},
		{
			name: "ok/only age",
			prepare: func(ctx context.Context, s store.Store) *service.UserUpdateIn {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &service.UserUpdateIn{
					ID:  u.ID,
					Age: toPtr(int16(18)),
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := newStore()
			userService := service.New(store)

			ctx := context.Background()
			dto := tc.prepare(ctx, store)
			user, err := userService.UpdateUser(ctx, dto)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
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
