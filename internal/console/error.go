package console

import (
	"errors"
)

// Errors.
var (
	ErrCommandNotFound    = errors.New("command not found")
	ErrSubcommandNotFound = errors.New("subcommand not found")
	ErrInvalidArgs        = errors.New("invalid arguments")
)
