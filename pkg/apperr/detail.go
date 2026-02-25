package apperr

type ErrorDetail struct {
	Err   error
	Field string
}

func (e ErrorDetail) Unwrap() error {
	return e.Err
}

func (e ErrorDetail) Error() string {
	return e.Err.Error()
}
