package request

type Body interface {
	Validate() ValidationErrors
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateUser) Validate() ValidationErrors {
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

type UserUpdate struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int16  `json:"age"`
}

func (u *UserUpdate) Validate() ValidationErrors {
	var errs ValidationErrors
	if u.FirstName == nil && u.LastName == nil && u.Age == nil {
		errs = append(errs, ErrEmptyFieldsUpdate)
	}
	return errs
}
