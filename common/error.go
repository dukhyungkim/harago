package common

import (
	"fmt"
)

var (
	ErrDuplicateCommand = func(name string) error { return fmt.Errorf("duplicate command name: %s", name) }
)
