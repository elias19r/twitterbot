package like

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Likes action.
func New(name string) *behavior.Behavior {
	min := config.LikeMinInterval
	max := config.LikeMaxInterval
	delay := config.LikeBootupDelay

	b := behavior.New(name, min, max, action)
	b.SetBootupDelay(delay)

	return b
}

// actions likes random tweets over time.
func action(ctx context.Context, logger *log.Logger) {
	n := rand.Intn(config.LikeMaxN-config.LikeMinN) + config.LikeMinN

	timer := time.NewTimer(1 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			id, _ := twitter.GetRandomTweet(config.LikeWOEID)
			twitter.FavTweet(id)

			d := rand.Intn(config.LikeMaxDelay-config.LikeMinDelay) + config.LikeMinDelay
			timer.Stop()
			timer = time.NewTimer(time.Duration(d) * time.Second)
		}
	}
}
