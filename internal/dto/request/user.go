package request

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int16  `json:"age"`
}
