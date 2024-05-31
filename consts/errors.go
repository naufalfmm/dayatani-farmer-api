package consts

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAuth = errors.New("invalid auth")
)

var (
	ErrEntityNotFoundBuilder = func(ent string) error {
		return fmt.Errorf("%s not found", ent)
	}
)
