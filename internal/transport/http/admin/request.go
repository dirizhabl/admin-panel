package admin

import "admin-panel/pkg/apperr"

type Body interface {
	Validate() []apperr.ErrorDetail
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) Validate() []apperr.ErrorDetail {
	var errs = make([]apperr.ErrorDetail, 0)
	if len(u.Password) < 7 {
		errs = append(errs, apperr.ErrorDetail{Err: ErrWeakPassword, Field: "password"})
	}
	if u.Email == "" {
		errs = append(errs, apperr.ErrorDetail{Err: ErrEmptyEmail, Field: "email"})
	}
	if !IsValidEmail(u.Email) {
		errs = append(errs, apperr.ErrorDetail{Err: ErrNotValidEmail, Field: "email"})
	}
	return errs
}

type UserUpdateRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int16  `json:"age"`
}

func (u *UserUpdateRequest) Validate() []apperr.ErrorDetail {
	var errs = make([]apperr.ErrorDetail, 0)
	if u.Age != nil && *u.Age < 0 {
		errs = append(errs, apperr.ErrorDetail{Err: ErrNegativeAge, Field: "age"})
	}
	return errs
}

type UserFiltersQuery struct {
	Email     *string `schema:"email"`
	FirstName *string `schema:"first_name"`
	LastName  *string `schema:"last_name"`
	MinAge    *int16  `schema:"min_age"`
	MaxAge    *int16  `schema:"max_age"`
}

func (u *UserFiltersQuery) Validate() []apperr.ErrorDetail {
	var errs = make([]apperr.ErrorDetail, 0)

	if u.MinAge != nil && *u.MinAge < 0 {
		errs = append(errs, apperr.ErrorDetail{Err: ErrMinAge, Field: "min_age"})
	}
	if u.MaxAge != nil && *u.MaxAge > 255 {
		errs = append(errs, apperr.ErrorDetail{Err: ErrMaxAge, Field: "max_age"})
	}
	if u.MinAge != nil && u.MaxAge != nil && *u.MinAge > *u.MaxAge {
		errs = append(errs, apperr.ErrorDetail{Err: ErrMaxAge})
	}

	return errs
}
