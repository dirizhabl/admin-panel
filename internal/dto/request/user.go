package request

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int16  `json:"age"`
}
