package model_test

import (
	"http-rest-api/internal/model"
	"testing"

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
				return model.TestUser()
			},
			wantErr: nil,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = ""
				return u
			},
			wantErr: model.ErrEmptyEmail,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				return u
			},
			wantErr: model.ErrEmptyPassword,
		},
		{
			name: "with hashed password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				u.Hashed_password = "hashed_password"
				return u
			},
			wantErr: nil,
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
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.Hashed_password)
}
