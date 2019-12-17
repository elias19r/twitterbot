package command

import (
	"errors"
)

// Errors.
var (
	ErrCommandNotFound = errors.New("command not found")
	ErrInvalidMessage  = errors.New("invalid command message")
	ErrInvalidArgs     = errors.New("invalid message arguments")
)
