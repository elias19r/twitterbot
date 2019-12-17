package cmddm

import (
	"context"
	"log"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/command"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Command DM action.
func New(name string) *behavior.Behavior {
	min := config.CommandMinInterval
	max := config.CommandMaxInterval
	return behavior.New(name, min, max, action)
}

// action listens to a channel of twitter.DM. For every DM received,
// it tries to handle it as a bot command.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dmC := twitter.GetStreamDMC(ctx)
	for dm := range dmC {
		logger.Println(dm)
		go func(dm twitter.DM) {
			err := command.Handle(dm)
			if err != nil {
				logger.Println(err)
			}
			if err == twitter.ErrDailyLimitDM {
				cancel()
			}
		}(dm)
	}
}
