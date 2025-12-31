package postgres

type userRow struct {
	ID        int
	Email     string
	FirstName *string
	LastName  *string
	Age       *int16
}
