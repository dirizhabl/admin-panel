package command_test

import (
	"testing"

	"admin-panel/internal/user/service/command"

	"github.com/stretchr/testify/assert"
)

func TestUserCreate_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *command.UserCreate
		wantErr error
	}{
		{
			name: "ok",
			u: func() *command.UserCreate {
				u := command.TestUserCreate()
				return u
			},
		},
		{
			name: "empty email",
			u: func() *command.UserCreate {
				u := command.TestUserCreate()
				u.Email = ""
				return u
			},
			wantErr: command.ErrInvalidInput,
		},
		{
			name: "empty password",
			u: func() *command.UserCreate {
				u := command.TestUserCreate()
				u.Password = ""
				return u
			},
			wantErr: command.ErrInvalidInput,
		},
		{
			name: "empty email and password",
			u: func() *command.UserCreate {
				u := command.TestUserCreate()
				u.Email = ""
				u.Password = ""
				return u
			},
			wantErr: command.ErrInvalidInput,
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
