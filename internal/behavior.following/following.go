package following

import (
	"context"
	"log"
	"time"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a Behavior with following action.
func New(name string) *behavior.Behavior {
	min := config.FollowingMinInterval
	max := config.FollowingMaxInterval
	delay := config.FollowingBootupDelay

	b := behavior.New(name, min, max, action)
	b.SetBootupDelay(delay)

	return b
}

// action listens to a channel of following (twitter.User). For every following
// user read from the channel, it decides whether the bot should mute/unmute and
// follow/unfollow.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.FollowingRunningTimeout)*time.Second)
	defer cancel()

	following := twitter.GetFollowingC(ctx, "")
	users := []twitter.User{}
	for u := range following {
		users = append(users, u)
		if len(users) >= 100 { // Lookup chunks of 100 users.
			err := lookup(users)
			users = users[:0]
			if err != nil {
				logger.Println(err)
				break
			}
		}
	}
	lookup(users)
}

func lookup(users []twitter.User) error {
	lus, _ := twitter.GetLookup(users...)
	for _, lu := range lus {
		// NOTE: could use context.Context here
		twitter.WillMute(lu)
		err := twitter.WillUnfollow(lu)
		if err == twitter.ErrDailyLimitFollowing {
			return err
		}
	}
	if len(lus) > 0 {
		time.Sleep(time.Duration(config.FollowingDelayBetweenRequests) * time.Second)
	}
	return nil
}
