package user

import "errors"

var ErrMinMaxAge = errors.New("max age must be more min age")

type Filters struct {
	Email     *string
	FirstName *string
	LastName  *string
	MinAge    *int16
	MaxAge    *int16
}

func (f *Filters) Validate() error {
	if f.MinAge != nil && f.MaxAge != nil {
		if *f.MinAge > *f.MaxAge {
			return ErrMinMaxAge
		}
	}
	return nil
}
