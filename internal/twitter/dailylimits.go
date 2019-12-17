package twitter

import (
	"sync"
	"time"

	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/garbage"
)

// Daily limit counters.
var (
	dailyLimitFollowing dailyLimitCounter
	dailyLimitDM        dailyLimitCounter
	dailyLimitTweet     dailyLimitCounter
)

type dailyLimitCounter struct {
	name    string
	message string
	max     int

	mu     sync.Mutex // mutex guards zeroed, count, end
	zeroed bool
	count  int
	end    time.Time
}

func init() {
	now := time.Now()

	dailyLimitFollowing.name = "followinglimit"
	dailyLimitFollowing.message = config.DailyLimitFollowing()
	dailyLimitFollowing.max = 1000
	dailyLimitFollowing.count = dailyLimitFollowing.max
	dailyLimitFollowing.end = now.Add(24 * time.Hour)

	dailyLimitDM.name = "dmlimit"
	dailyLimitDM.message = config.DailyLimitDM()
	dailyLimitDM.max = 1000 - 10
	dailyLimitDM.count = dailyLimitDM.max
	dailyLimitDM.end = now.Add(24 * time.Hour)

	dailyLimitTweet.name = "tweetlimit"
	dailyLimitTweet.message = config.DailyLimitTweet()
	dailyLimitTweet.max = 2400 - 10
	dailyLimitTweet.count = dailyLimitTweet.max
	dailyLimitTweet.end = now.Add(24 * time.Hour)
}

func (dlc *dailyLimitCounter) dec() bool {
	dlc.checkTime()

	dlc.mu.Lock()
	defer dlc.mu.Unlock()

	if dlc.count < 1 {
		if !dlc.zeroed {
			dlc.zeroed = true
			dlc.count = 0
			dlc.end = time.Now().Add(24 * time.Hour)
			id, _ := PostTweet(Bot, dlc.message) // Post a notice tweet.
			garbage.AddTweetID(id)
		}
		return false
	}
	dlc.count--
	return true
}

func (dlc *dailyLimitCounter) checkTime() {
	dlc.mu.Lock()
	now := time.Now()
	if now.Sub(dlc.end) >= time.Duration(0) {
		dlc.zeroed = false
		dlc.count = dlc.max
		dlc.end = now.Add(24 * time.Hour)
	}
	dlc.mu.Unlock()
}

func (dlc *dailyLimitCounter) drain() {
	dlc.mu.Lock()

	if !dlc.zeroed {
		dlc.zeroed = true
		dlc.count = 0
		dlc.end = time.Now().Add(24 * time.Hour)
		id, _ := PostTweet(Bot, dlc.message) // Post a notice tweet.
		garbage.AddTweetID(id)
	}

	dlc.mu.Unlock()
}
