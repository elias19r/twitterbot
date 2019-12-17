package autolikes

import (
	"context"
	"log"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with AutoLikes action.
func New(name string) *behavior.Behavior {
	min := config.AutoLikesMinInterval
	max := config.AutoLikesMaxInterval
	return behavior.New(name, min, max, action)
}

// action listens to a channel of twitter.Tweet. For every tweet received,
// if it is of type "mention" or "tweetreply", action likes the tweet and then
// decides whether to follow the user.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tweetC := twitter.GetStreamTweetC(ctx)
	for tweet := range tweetC {
		if tweet.Type != "mention" && tweet.Type != "tweetreply" {
			continue
		}
		logger.Println(tweet)
		twitter.FavTweet(tweet.ID)

		lu, err := twitter.GetLookupOne(tweet.User)
		if err != nil {
			logger.Println(err)
			continue
		}

		err = twitter.WillFollow(lu)
		if err != nil {
			logger.Println(err)
		}
	}
}
