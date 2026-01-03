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
		prepare func(Store) *model.User
		wantErr error
	}{
		{
			name: "user already exists",
			prepare: func(s Store) *model.User {
				u := model.TestUser()
				s.User().Create(u)
				return u
			},
			wantErr: ErrRecordExists,
		},
		{
			name: "ok",
			prepare: func(s Store) *model.User {
				return model.TestUser()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			user := tc.prepare(store)
			err := store.User().Create(user)

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
		prepare func(Store) int
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(s Store) int {
				return math.MaxInt64
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(s Store) int {
				u := model.TestUser()
				s.User().Create(u)
				return u.ID
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			id := tc.prepare(store)
			user, err := store.User().FindById(id)

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
		prepare func(Store) string
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(s Store) string {
				return "nil@nil.org"
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(s Store) string {
				u := model.TestUser()
				s.User().Create(u)
				return u.Email
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			email := tc.prepare(store)
			user, err := store.User().FindByEmail(email)

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
		prepare func(Store) *model.UserUpdate
		wantErr error
	}{
		{
			name: "user not found",
			prepare: func(s Store) *model.UserUpdate {
				id := math.MaxInt64
				return &model.UserUpdate{
					ID: id,
				}
			},
			wantErr: ErrRecordNotFound,
		},
		{
			name: "ok",
			prepare: func(s Store) *model.UserUpdate {
				u := model.TestUser()
				s.User().Create(u)
				return &model.UserUpdate{
					ID: u.ID,
					FirstName: toPtr("john"),
					LastName : toPtr("doe"),
					Age : toPtr(int16(18)),
				}
			},
		},
		{
			name: "ok/only first_name",
			prepare: func(s Store) *model.UserUpdate {
				u := model.TestUser()
				s.User().Create(u)
				return &model.UserUpdate{
					ID: u.ID,
					FirstName: toPtr("john"),
				}
			},
		},
		{
			name: "ok/only last_name",
			prepare: func(s Store) *model.UserUpdate {
				u := model.TestUser()
				s.User().Create(u)
				return &model.UserUpdate{
					ID: u.ID,
					LastName: toPtr("doe"),
				}
			},
		},
		{
			name: "ok/only age",
			prepare: func(s Store) *model.UserUpdate {
				u := model.TestUser()
				s.User().Create(u)
				return &model.UserUpdate{
					ID: u.ID,
					Age: toPtr(int16(18)),
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store, teardown := newStore()
			defer teardown()

			dto := tc.prepare(store)
			user, err := store.User().Update(dto)

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
