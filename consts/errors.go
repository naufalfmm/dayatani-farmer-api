package consts

import (
	"fmt"
)

var (
	ErrEntityNotFoundBuilder = func(ent string) error {
		return fmt.Errorf("%s not found", ent)
	}
)
