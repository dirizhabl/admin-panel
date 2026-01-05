package user

import (
	"encoding/json"
	"errors"
	"net/http"

	req "admin-panel/internal/transport/http/admin-panel/request"
	resp "admin-panel/internal/transport/http/admin-panel/response"
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

func (c *context) BindJson(body req.Body) error {
	if c.r.Body == nil {
		c.JSON(400, resp.Err{Err: "empty body"})
	}
	dec := json.NewDecoder(c.r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(body); err != nil {
		c.JSON(400, resp.Err{Err: err.Error()})
		return err
	}
	if errs := body.Validate(); len(errs) > 0 {
		c.JSON(422, resp.Errs{Errs: errs.Errors()})
		return errs
	}
	return nil
}

func (c *context) Error(err error) {
	switch {
	case errors.Is(err, store.ErrRecordExists):
		c.JSON(409, resp.Err{Err: err.Error()})
	case errors.Is(err, store.ErrRecordNotFound):
		c.JSON(404, resp.Err{Err: err.Error()})
	default:
		c.JSON(500, resp.Err{Err: err.Error()})
	}
}
