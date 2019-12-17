package cmdtweet

import (
	"context"
	"log"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/command"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Command Tweet action.
func New(name string) *behavior.Behavior {
	min := config.CommandMinInterval
	max := config.CommandMaxInterval
	return behavior.New(name, min, max, action)
}

// action listens to a channel of twitter.Tweet. For every tweet received,
// it tries to handle it as a bot command.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tweetC := twitter.GetStreamTweetC(ctx)
	for tweet := range tweetC {
		go func(tweet twitter.Tweet) {
			err := command.Handle(tweet)
			if err == twitter.ErrDailyLimitTweet {
				logger.Println(err)
				cancel()
			}
		}(tweet)
	}
}
