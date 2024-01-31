package handlers

import (
	"errors"
	"fmt"
)

var (
	ErrNilEvent   = errors.New("received nil event")
	ErrBadRequest = errors.New("bad request")
)

func badRequest(paramName string) error {
	return fmt.Errorf("%w: excepted param %s", ErrBadRequest, paramName)
}
