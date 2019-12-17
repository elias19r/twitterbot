package twitter

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/pkg/text"
)

// PostTweetAt adds a ".@user" to content and then tweets.
func PostTweetAt(user User, content string) (int64, error) {
	content = ".@" + user.Username + " " + content
	return PostTweet(user, content)
}

// PostTweetCC adds a "cc @user" to content and then tweets.
func PostTweetCC(user User, content string) (int64, error) {
	content = "cc @" + user.Username + " " + content
	return PostTweet(user, content)
}

// PostTweet posts tweet and then sends tweet link to user via DM.
func PostTweet(user User, content string) (int64, error) {
	if !dailyLimitTweet.dec() {
		return 0, ErrDailyLimitTweet
	}

	content = text.Truncate(content, 280)
	tweet, err := api.PostTweet(content, url.Values{})
	if err != nil {
		SendDM(user, config.TweetFailed)
		if err, ok := err.(*anaconda.ApiError); ok {
			if twitterError, ok := err.Decoded.First().(anaconda.TwitterError); ok {
				switch twitterError.Code {
				case 64, 89, 185, 205, 226, 326:
					logger.Println(err)
					dailyLimitTweet.drain()
					return 0, ErrDailyLimitTweet
				}
			}
		}
		return 0, err
	}
	link := fmt.Sprintf("https://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
	logger.Printf("sent tweet: %q [%s]\n", content, link)

	// Send tweet link via DM to user.
	SendDM(user, link)
	SendDM(user, config.GetEmojiOK()) // NOTE: comment this out to improve DM rate limit.

	return tweet.Id, nil
}

// PostTweetReply sends a tweet replying to a tweet specified by replyID.
func PostTweetReply(user User, content string, replyID int64) (int64, error) {
	if !dailyLimitTweet.dec() {
		return 0, ErrDailyLimitTweet
	}

	content = "@" + user.Username + " " + content
	content = text.Truncate(content, 280)

	v := url.Values{}
	v.Set("in_reply_to_status_id", strconv.FormatInt(replyID, 10))
	tweet, err := api.PostTweet(content, v)
	if err != nil {
		SendDM(user, config.TweetReplyFailed)
		if err, ok := err.(*anaconda.ApiError); ok {
			if twitterError, ok := err.Decoded.First().(anaconda.TwitterError); ok {
				switch twitterError.Code {
				case 64, 89, 185, 205, 226, 326:
					logger.Println(err)
					dailyLimitTweet.drain()
					return 0, ErrDailyLimitTweet
				}
			}
		}
		return 0, err
	}
	link := fmt.Sprintf("https://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
	logger.Printf("sent tweetreply: %q [%s]\n", content, link)

	return tweet.Id, nil
}

// DeleteTweet deletes a tweet specified by id.
func DeleteTweet(id int64) error {
	if id <= 0 {
		return ErrInvalidTweetID
	}
	_, err := api.DeleteTweet(id, true)
	if err != nil {
		return err
	}
	logger.Println("deleted tweet", id)
	return nil
}

// FavTweet likes/favorite a specified tweet.
func FavTweet(id int64) error {
	if id <= 0 {
		return ErrInvalidTweetID
	}
	tweet, err := api.Favorite(id)
	if err != nil {
		return err
	}
	logger.Printf("liked tweet: https://twitter.com/%s/status/%s\n", tweet.User.ScreenName, tweet.IdStr)
	return nil
}

// GetRandomTweet returns ID of a random tweet from trends.
func GetRandomTweet(woeid int64) (int64, error) {
	// Get trends for a location WOIED.
	loc, err := api.GetTrendsByPlace(woeid, nil)
	if err != nil {
		return 0, err
	}
	// Pick a trend at random from the list of trends.
	if len(loc.Trends) <= 0 {
		return 0, ErrNoTrendsFound
	}
	// Ensure top 10 trends.
	if len(loc.Trends) > 10 {
		loc.Trends = loc.Trends[:10]
	}
	trend := loc.Trends[rand.Intn(len(loc.Trends))]

	// Search for tweets with that trend query.
	v := url.Values{}
	v.Set("count", "50")
	v.Set("lang", "pt")
	v.Set("include_entities", "false")
	sr, err := api.GetSearch(trend.Query, v)
	if err != nil {
		return 0, err
	}
	// Pick a tweet at random from search response.
	if len(sr.Statuses) <= 0 {
		return 0, ErrNoStatusesFound
	}
	tweet := sr.Statuses[rand.Intn(len(sr.Statuses))]

	return tweet.Id, nil
}
