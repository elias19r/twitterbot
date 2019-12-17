package gc

import (
	"context"
	"log"
	"time"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/garbage"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Garbage Collector (GC) action.
func New(name string) *behavior.Behavior {
	min := config.GCMinInterval
	max := config.GCMaxInterval
	delay := config.GCBootupDelay

	b := behavior.New(name, min, max, action)
	b.SetBootupDelay(delay)

	return b
}

// action reads tweet IDs from garbage package and delete them.
func action(ctx context.Context, logger *log.Logger) {
	cleaned := []int64{}

	tweetIDs := garbage.GetTweetIDs()
loop:
	for _, id := range tweetIDs {
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(5 * time.Second): // Add some delay between deletes.
			err := twitter.DeleteTweet(id)
			if err != nil {
				logger.Println(err)
				continue
			}
			cleaned = append(cleaned, id)
		}
	}
	garbage.RmTweetID(cleaned...)
}
