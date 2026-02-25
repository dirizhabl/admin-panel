package apperr

type Code string

const (
	BadRequest Code = "bad_request"
	Validation Code = "validation_error"
	Conflict   Code = "conflict"
	NotFound   Code = "not_found"
	Internal   Code = "internal"
)

func NewConflict(err error, message string) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Code:    Conflict,
		Message: message,
	}
}

func NewNotFound(err error, message string) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Code:    NotFound,
		Message: message,
	}
}

func NewValidation(err error, message string) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Code:    Validation,
		Message: message,
	}
}

func NewValidationWithDetails(errs []ErrorDetail) *ErrorResponse {
	details := make([]DetailResponse, len(errs))

	for i, e := range errs {
		details[i] = DetailResponse{
			Field:   e.Field,
			Message: e.Error(),
		}
	}

	return &ErrorResponse{
		Code:    Validation,
		Message: "validation error",
		Details: details,
	}
}

func NewInternal(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Code:    Internal,
		Message: "internal server error",
	}
}
