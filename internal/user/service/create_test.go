package service_test

import (
	"testing"

	"admin-panel/internal/user/service"

	"github.com/stretchr/testify/assert"
)

func TestUserCreateIn_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *service.UserCreateIn
		wantErr error
	}{
		{
			name: "ok",
			u: func() *service.UserCreateIn {
				return &service.UserCreateIn{
					Email:    "nil@nil.org",
					Password: "123",
				}
			},
		},
		{
			name: "empty email",
			u: func() *service.UserCreateIn {
				return &service.UserCreateIn{
					Password: "123",
				}
			},
			wantErr: service.ErrInvalidInput,
		},
		{
			name: "empty password",
			u: func() *service.UserCreateIn {
				return &service.UserCreateIn{
					Email: "nil@nil.org",
				}
			},
			wantErr: service.ErrInvalidInput,
		},
		{
			name: "empty email and password",
			u: func() *service.UserCreateIn {
				return &service.UserCreateIn{}
			},
			wantErr: service.ErrInvalidInput,
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
