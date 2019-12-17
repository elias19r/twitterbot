package behavior

import (
	"bytes"
	"context"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Behavior represents an action that is performed over time.
type Behavior struct {
	LogWriter io.ReadWriter // LogWriter can be used to read log from Behavior.
	logger    *log.Logger

	name string

	// Minimum and maximum intervals in seconds between executions.
	minInterval int
	maxInterval int
	bootupDelay int // Delay in seconds before the first execution.

	// Function that will run.
	action func(context.Context, *log.Logger)

	// Variables to control its execution.
	mu      sync.Mutex // mu guards started, idle, cancel, ctx.
	started bool
	idle    bool
	cancel  context.CancelFunc
	ctx     context.Context
}

// New returns a new *Behavior.
func New(name string, min, max int, action func(context.Context, *log.Logger)) *Behavior {
	b := &Behavior{
		name:        name,
		minInterval: min,
		maxInterval: max,
		action:      action,
		bootupDelay: 1, // Minimum boot up delay: 1 second.
	}

	b.LogWriter = new(bytes.Buffer)
	b.logger = log.New(b.LogWriter, b.name+": ", 0)
	b.idle = true

	return b
}

// SetBootupDelay sets the bootup delay in seconds.
func (b *Behavior) SetBootupDelay(s int) {
	b.bootupDelay = s
}

// Name returns name.
func (b *Behavior) Name() string {
	return b.name
}

// Started returns started.
func (b *Behavior) Started() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.started
}

// Idle returns idle.
func (b *Behavior) Idle() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.idle
}

func (b *Behavior) setIdle(v bool) {
	b.mu.Lock()
	b.idle = v
	b.mu.Unlock()
}

// Status returns a status string.
func (b *Behavior) Status() string {
	status := b.name + ": "

	b.mu.Lock()
	if b.started {
		status += "started ("
		if b.idle {
			status += "idle"
		} else {
			status += "running"
		}
		status += ")"
	} else {
		status += "stopped"
	}
	b.mu.Unlock()

	return status
}

// Start starts performing behavior's action function.
func (b *Behavior) Start(skipBootupDelay bool) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.started {
		return ErrAlreadyStarted
	}
	b.ctx, b.cancel = context.WithCancel(context.Background())
	b.started = true
	b.logger.Println("started")

	go func(b *Behavior, skipBootupDelay bool) {
		defer b.cancel()

		d := b.bootupDelay
		if skipBootupDelay {
			d = 1
		}
		timer := time.NewTimer(time.Duration(d) * time.Second)
		for {
			select {
			case <-b.ctx.Done():
				return
			case <-timer.C:
				b.setIdle(false)
				b.logger.Println("running")
				b.action(b.ctx, b.logger)
				b.logger.Println("idle")
				b.setIdle(true)
			}

			d := rand.Intn(b.maxInterval-b.minInterval) + b.minInterval
			timer.Stop()
			timer = time.NewTimer(time.Duration(d) * time.Second)
		}
	}(b, skipBootupDelay)

	return nil
}

// Stop stops behavior's action function.
func (b *Behavior) Stop() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.started {
		b.cancel()
		b.started = false
		b.logger.Println("stopped")
		return nil
	}
	return ErrNotStarted
}
