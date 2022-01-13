package common

import (
	"fmt"
)

var (
	ErrDuplicateCommand = func(name string) error { return fmt.Errorf("duplicate command name: %s", name) }

	ErrHarborResponse = func(err error) error { return fmt.Errorf("HarborError!: %v", err) }
)
