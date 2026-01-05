package service

type UserUpdateIn struct {
	ID        int
	FirstName *string
	LastName  *string
	Age       *int16
}
