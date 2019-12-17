package behavior

import (
	"errors"
)

// Errors.
var (
	ErrNotFound       = errors.New("behavior not found")
	ErrAlreadyStarted = errors.New("behavior is already started")
	ErrNotStarted     = errors.New("behavior is not started")
)
