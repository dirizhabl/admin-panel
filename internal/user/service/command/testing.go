package command

func TestUserCreate() *UserCreate {
	return &UserCreate{
		Email:    "nil@nil.org",
		Password: "123",
	}
}
