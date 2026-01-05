package admin

type Body interface {
	Validate() ValidationErrors
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) Validate() ValidationErrors {
	var errs ValidationErrors
	if len(u.Password) < 7 {
		errs = append(errs, ErrWeakPassword)
	}
	if u.Email == "" {
		errs = append(errs, ErrEmptyEmail)
	}
	if !IsValidEmail(u.Email) {
		errs = append(errs, ErrNotValidEmail)
	}
	return errs
}

type UserUpdateRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int16  `json:"age"`
}

func (u *UserUpdateRequest) Validate() ValidationErrors {
	var errs ValidationErrors
	if u.FirstName == nil && u.LastName == nil && u.Age == nil {
		errs = append(errs, ErrEmptyFieldsUpdate)
	}
	return errs
}
