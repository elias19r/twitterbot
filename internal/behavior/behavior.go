package behavior

import (
	"sync"
)

// bhvrs is a list of all available behaviors.
var (
	mu    sync.Mutex
	bhvrs []*Behavior
)

// Add adds one or more behaviors to the list.
func Add(bs ...*Behavior) {
	mu.Lock()
	bhvrs = append(bhvrs, bs...)
	mu.Unlock()
}

// Get returns behavior called name.
func Get(name string) (*Behavior, error) {
	for _, b := range bhvrs {
		if b.Name() == name {
			return b, nil
		}
	}
	return nil, ErrNotFound
}

// Start starts behavior called name.
func Start(name string) error {
	for _, b := range bhvrs {
		if b.Name() == name {
			return b.Start(false)
		}
	}
	return ErrNotFound
}

// Stop stops a behavior called name.
func Stop(name string) error {
	for _, b := range bhvrs {
		if b.Name() == name {
			return b.Stop()
		}
	}
	return ErrNotFound
}

// StartAll starts all available behaviors.
func StartAll() {
	for _, b := range bhvrs {
		b.Start(false)
	}
}

// StopAll stops all available behaviors.
func StopAll() {
	for _, b := range bhvrs {
		b.Stop()
	}
}

// List returns a list with names of available behaviors.
func List() []string {
	names := []string{}
	for _, b := range bhvrs {
		names = append(names, b.Name())
	}
	return names
}

// Info returns a list with all behavior status.
func Info() []string {
	info := []string{}
	for _, b := range bhvrs {
		info = append(info, b.Status())
	}
	return info
}
