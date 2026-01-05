package domain_test

import (
	"testing"

	"admin-panel/internal/user/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	testCases := []struct {
		name    string
		email   string
		pasword string
		wantErr error
	}{
		{
			name:    "ok",
			email:   "nil@nil.org",
			pasword: "1231231",
		},
		{
			name:    "empty email",
			email:   "",
			pasword: "1231231",
			wantErr: domain.ErrEmptyEmail,
		},
		{
			name:    "weak password",
			email:   "nil@nil.org",
			pasword: "123",
			wantErr: domain.ErrWeakPassword,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := domain.NewUser(tc.email, tc.pasword)
			if tc.wantErr != nil {
				assert.Nil(t, u)
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NotNil(t, u)
				assert.NoError(t, err)
			}
		})
	}
}
