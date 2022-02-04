package common

import (
	"fmt"
)

var (
	ErrDuplicateCommand = func(name string) error { return fmt.Errorf("duplicate command name: %s", name) }

	ErrConnEtcd = func(err error) error { return fmt.Errorf("failed to connect etcd; %w", err) }

	ErrHarborResponse = func(err error) error { return fmt.Errorf("HarborError!: %v", err) }
)
