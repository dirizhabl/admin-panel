package response

type UserCreate struct {
	ID int `json:"id"`
}

type UserRead struct {
	ID        int    `json:"id"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int16  `json:"age,omitempty"`
}
