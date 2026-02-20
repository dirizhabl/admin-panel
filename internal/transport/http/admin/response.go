package admin

import "admin-panel/internal/user/service"

type ErrorResponse struct {
	Error string `json:"error"`
}

func toErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}

type ErrorsResponse struct {
	Errors []string `json:"errors"`
}

func toErrorsResponse(errs []string) ErrorsResponse {
	return ErrorsResponse{
		Errors: errs,
	}
}

type UserCreateResponse struct {
	ID int `json:"id"`
}

type UserReadResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int16  `json:"age,omitempty"`
}

func toUserReadResponse(u *service.UserReadOut) UserReadResponse {
	return UserReadResponse{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
}
