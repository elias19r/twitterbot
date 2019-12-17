package crawler

import (
	"errors"
)

// Errors.
var (
	ErrCrawlerMax = errors.New("maximum followers, done")
	ErrNoUsername = errors.New("could not pick a username")
)
