package service

type UserCreateIn struct {
	Email    string
	Password string
}

type UserCreateOut struct {
	ID int
}
