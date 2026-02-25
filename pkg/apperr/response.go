package apperr

type DetailResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Err     error            `json:"-"`
	Code    Code             `json:"code"`
	Message string           `json:"message"`
	Details []DetailResponse `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) Unwrap() error {
	return e.Err
}
