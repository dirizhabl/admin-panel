package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"admin-panel/pkg/apperrors"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewHandlerContext(w http.ResponseWriter, r *http.Request) *context {
	return &context{w: w, r: r}
}

func (c *context) JSON(statusCode int, response any) {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(statusCode)
	json.NewEncoder(c.w).Encode(response)
}

func (c *context) BindJson(body Body) error {
	if c.r.Body == nil {
		return apperrors.ValidationError("empty body")
	}
	dec := json.NewDecoder(c.r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(body); err != nil {
		return apperrors.ValidationError("cannot decode body")
	}

	apperr := apperrors.ValidationError("calidation error")
	if errs := body.Validate(); len(errs) > 0 {
		for _, e := range errs {
			apperr.Details = append(apperr.Details, apperrors.Detail{
				Field:   e.Field,
				Message: e.Error(),
			})
		}
		return apperr
	}
	return nil
}

func (c *context) BindQueryParams(body Body) error {
	if err := decoder.Decode(body, c.r.URL.Query()); err != nil {
		return apperrors.ValidationError("cannot decode query params")
	}
	apperr := apperrors.ValidationError("validation error")
	if errs := body.Validate(); len(errs) > 0 {
		for _, e := range errs {
			apperr.Details = append(apperr.Details, apperrors.Detail{
				Field:   e.Field,
				Message: e.Error(),
			})
		}
		return apperr
	}
	return nil
}

func (c *context) Error(err error) {
	var apperr apperrors.AppError

	switch {
	case errors.As(err, &apperr):
		c.JSON(MapCodeToStatusCode()[apperr.Code], apperr)
	}
}

func MapCodeToStatusCode() map[apperrors.Code]int {
	return map[apperrors.Code]int{
		apperrors.BadRequest: 400,
		apperrors.NotFound:   404,
		apperrors.Conflict:   409,
		apperrors.Validation: 422,
	}
}
