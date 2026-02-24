package apperrors

type Code string

const (
	BadRequest Code = "bad_request"
	Validation Code = "validation_error"
	Conflict   Code = "conflict"
	NotFound   Code = "not_found"
	Internal   Code = "internal"
)

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

type AppError struct {
	Err     error
	Code    Code     `json:"code"`
	Message string   `json:"message"`
	Path    string   `json:"-"`
	Details []Detail `json:"details,omitempty"`
}

type Detail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Unwrap() error {
	return e.Err
}

func New(err error, code Code, message string) AppError {
	return AppError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

func ValidationError(message string) AppError {
	return AppError{
		Code:    Validation,
		Message: message,
	}
}
