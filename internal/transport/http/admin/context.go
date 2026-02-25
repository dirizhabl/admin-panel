package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"admin-panel/pkg/apperr"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

var (
	codeMapper = map[apperr.Code]int{
		apperr.BadRequest: 400,
		apperr.NotFound:   404,
		apperr.Conflict:   409,
		apperr.Validation: 422,
		apperr.Internal:   500,
	}
	decoder = schema.NewDecoder()
)

type context struct {
	w http.ResponseWriter
	r *http.Request
	l *logrus.Logger
}

func NewHandlerContext(w http.ResponseWriter, r *http.Request, l *logrus.Logger) *context {
	return &context{w: w, r: r, l: l}
}

func (c *context) JSON(statusCode int, response any) {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(statusCode)
	json.NewEncoder(c.w).Encode(response)
}

func (c *context) BindJson(body Body) error {
	if c.r.Body == nil {
		return apperr.NewValidation(nil, "empty body")
	}
	dec := json.NewDecoder(c.r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(body); err != nil {
		return apperr.NewValidation(err, "cannot decode body")
	}
	if errs := body.Validate(); len(errs) > 0 {
		return apperr.NewValidationWithDetails(errs)
	}
	return nil
}

func (c *context) BindQueryParams(body Body) error {
	if err := decoder.Decode(body, c.r.URL.Query()); err != nil {
		return apperr.NewValidation(err, "cannot decode query params")
	}
	if errs := body.Validate(); len(errs) > 0 {
		return apperr.NewValidationWithDetails(errs)
	}
	return nil
}

func (c *context) Error(err error) {
	var appErr *apperr.ErrorResponse

	switch {
	case errors.As(err, &appErr):
		if appErr.Code == apperr.Internal {
			c.l.Error(appErr.Err.Error())
		}
		c.JSON(codeMapper[appErr.Code], appErr)
	default:
		c.JSON(500, err)
	}
}
