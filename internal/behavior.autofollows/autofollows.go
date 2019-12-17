package autofollows

import (
	"context"
	"log"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with AutoFollows action.
func New(name string) *behavior.Behavior {
	min := config.AutoFollowsMinInterval
	max := config.AutoFollowsMaxInterval
	return behavior.New(name, min, max, action)
}

// action listens to a channel of twitter.Follows. For every follows received,
// it decides whether to mute/unmute and to follow back the user.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	follows := twitter.GetStreamFollowsC(ctx)
	for flws := range follows {
		logger.Println(flws)
		user := flws.Source

		lu, err := twitter.GetLookupOne(user)
		if err != nil {
			logger.Println(err)
			continue
		}

		twitter.WillMute(lu)
		err = twitter.WillFollow(lu)
		if err != nil {
			logger.Println(err)
		}
		if err == twitter.ErrDailyLimitFollowing {
			break
		}
	}
}
