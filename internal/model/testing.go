package model

func TestUser() *User {
	return &User{
		Email: "user@example.com",
	}
}

func TestUserUpdate(id int) *UserUpdate {
	return &UserUpdate{
		ID: id,
	}
}
