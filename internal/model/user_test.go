package model_test

import (
	"testing"

	"admin-panel/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		wantErr error
	}{
		{
			name: "valid",
			u: func() *model.User {
				u := model.TestUser()
				u.BeforeCreate("123")
				return u
			},
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser()
				u.BeforeCreate("123")
				u.Email = ""
				return u
			},
			wantErr: model.ErrEmptyEmail,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.u().Validate(), tc.wantErr)
			} else {
				assert.NoError(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser()
	assert.NoError(t, u.BeforeCreate("123"))
	assert.NotEmpty(t, u.HashedPassword)

	u = model.TestUser()
	assert.ErrorIs(t, u.BeforeCreate(""), model.ErrEmptyPassword)
}
