package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"admin-panel/internal/user/store"
)

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
		c.JSON(400, toErrorResponse("empty body"))
	}
	dec := json.NewDecoder(c.r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(body); err != nil {
		c.JSON(400, toErrorResponse(err.Error()))
		return err
	}
	if errs := body.Validate(); len(errs) > 0 {
		c.JSON(422, toErrorsResponse(errs.Errors()))
		return errs
	}
	return nil
}

func (c *context) Error(err error) {
	switch {
	case errors.Is(err, store.ErrRecordExists):
		c.JSON(409, toErrorResponse(err.Error()))
	case errors.Is(err, store.ErrRecordNotFound):
		c.JSON(404, toErrorResponse(err.Error()))
	default:
		c.JSON(500, toErrorResponse(err.Error()))
	}
}
