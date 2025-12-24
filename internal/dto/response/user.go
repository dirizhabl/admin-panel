package response

type CreateUser struct {
	ID int `json:"id"`
}

type GetUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int16  `json:"age,omitempty"`
}
