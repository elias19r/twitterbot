package follower

import (
	"context"
	"log"
	"time"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/garbage"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Follower action.
func New(name string) *behavior.Behavior {
	min := config.FollowerMinInterval
	max := config.FollowerMaxInterval
	delay := config.FollowerBootupDelay

	b := behavior.New(name, min, max, action)
	b.SetBootupDelay(delay)

	return b
}

// action listens to a channel of followers (twitter.User). For every follower
// read from the channel, it decides whether the bot should mute/unmute and
// follow back.
func action(ctx context.Context, logger *log.Logger) {
	id1, _ := twitter.PostTweet(twitter.Bot, config.FollowingStartEN())
	id2, _ := twitter.PostTweet(twitter.Bot, config.FollowingStartPT())
	garbage.AddTweetID(id1, id2)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.FollowerRunningTimeout)*time.Second)
	defer cancel()

	followers := twitter.GetFollowersC(ctx, "")
	users := []twitter.User{}
	for u := range followers {
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
		// NOTE: could use context.Context here.
		twitter.WillMute(lu)
		err := twitter.WillFollow(lu)
		if err == twitter.ErrDailyLimitFollowing {
			return err
		}
	}
	if len(lus) > 0 {
		time.Sleep(time.Duration(config.FollowerDelayBetweenRequests) * time.Second)
	}
	return nil
}
