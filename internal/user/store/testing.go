package store

import (
	"context"
	"math"
	"testing"

	"admin-panel/internal/user/domain"

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
		prepare func(context.Context, Store) *domain.User
		wantErr error
	}{
		{
			name: "user already exists",
			prepare: func(ctx context.Context, s Store) *domain.User {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return u
			},
			wantErr: ErrRecordExists,
		},
		{
			name: "ok",
			prepare: func(ctx context.Context, s Store) *domain.User {
				return domain.TestUser()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			ctx := context.Background()
			user := tc.prepare(ctx, store)

			err := store.User().Create(ctx, user)

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
		prepare func(context.Context, Store) int
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(ctx context.Context, s Store) int {
				return math.MaxInt64
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(ctx context.Context, s Store) int {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return u.ID
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			ctx := context.Background()
			id := tc.prepare(ctx, store)

			user, err := store.User().FindById(ctx, id)

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
		prepare func(context.Context, Store) string
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(ctx context.Context, s Store) string {
				return "nil@nil.org"
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(ctx context.Context, s Store) string {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return u.Email
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			ctx := context.Background()
			email := tc.prepare(ctx, store)

			user, err := store.User().FindByEmail(ctx, email)

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
		prepare func(context.Context, Store) *domain.UserUpdate
		wantErr error
	}{
		{
			name: "empty input/user not exists",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				id := math.MaxInt64
				return &domain.UserUpdate{
					ID: id,
				}
			},
			wantErr: ErrNothingToUpdate,
		},
		{
			name: "empty input/user exists",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &domain.UserUpdate{
					ID: u.ID,
				}
			},
			wantErr: ErrNothingToUpdate,
		},
		{
			name: "user not found",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				id := math.MaxInt64
				return &domain.UserUpdate{
					ID: id,
					FirstName: toPtr("john"),
				}
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &domain.UserUpdate{
					ID:        u.ID,
					FirstName: toPtr("john"),
					LastName:  toPtr("doe"),
					Age:       toPtr(int16(18)),
				}
			},
		},
		{
			name: "ok/only first_name",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &domain.UserUpdate{
					ID:        u.ID,
					FirstName: toPtr("john"),
				}
			},
		},
		{
			name: "ok/only last_name",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &domain.UserUpdate{
					ID:       u.ID,
					LastName: toPtr("doe"),
				}
			},
		},
		{
			name: "ok/only age",
			prepare: func(ctx context.Context, s Store) *domain.UserUpdate {
				u := domain.TestUser()
				s.User().Create(ctx, u)
				return &domain.UserUpdate{
					ID:  u.ID,
					Age: toPtr(int16(18)),
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			ctx := context.Background()
			dto := tc.prepare(ctx, store)

			user, err := store.User().Update(ctx, dto)

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
