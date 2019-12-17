package twitter

import (
	"container/list"
	"context"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// handleTweet handles a Tweet message and broadcast it to listener of tweets.
func handleTweet(msg anaconda.Tweet) {
	// Ignore bot's own tweets.
	if msg.User.ScreenName == Bot.Username {
		return
	}
	hasMention := func() bool {
		for _, mention := range msg.Entities.User_mentions {
			if mention.Screen_name == Bot.Username {
				return true
			}
		}
		return false
	}

	tweet := Tweet{}
	tweet.ID = msg.Id
	tweet.Type = "tweet"

	switch {
	case msg.RetweetedStatus != nil:
		tweet.Type = "retweet"
	case hasMention():
		tweet.Type = "mention"
		fallthrough
	case msg.InReplyToScreenName == Bot.Username:
		tweet.Type = "tweetreply"
	}

	tweet.User = User{
		ID:             msg.User.Id,
		Username:       msg.User.ScreenName,
		Name:           msg.User.Name,
		Description:    msg.User.Description,
		FollowersCount: msg.User.FollowersCount,
		FollowingCount: msg.User.FriendsCount,
	}
	tweet.Content = msg.Text
	tweet.CreatedAt = msg.CreatedAt
	broadcast(tweet)
}

// handleDM handles a DM message and broadcast it to listener of DMs.
func handleDM(msg anaconda.DirectMessage) {
	// Ignore bot's own DM.
	if msg.Sender.ScreenName == Bot.Username {
		return
	}
	dm := DM{}
	dm.From = User{
		ID:             msg.Sender.Id,
		Username:       msg.Sender.ScreenName,
		Name:           msg.Sender.Name,
		Description:    msg.Sender.Description,
		FollowersCount: msg.Sender.FollowersCount,
		FollowingCount: msg.Sender.FriendsCount,
	}
	dm.To = User{
		ID:             msg.Recipient.Id,
		Username:       msg.Recipient.ScreenName,
		Name:           msg.Recipient.Name,
		Description:    msg.Recipient.Description,
		FollowersCount: msg.Recipient.FollowersCount,
		FollowingCount: msg.Recipient.FriendsCount,
	}
	dm.Content = msg.Text
	dm.CreatedAt = msg.CreatedAt
	broadcast(dm)
}

// handleFollow handles a Follows message and broadcast it to listener of Follows.
func handleFollows(msg anaconda.EventFollow) {
	// Ignore bot's own follows.
	if msg.Source.ScreenName == Bot.Username {
		return
	}
	flws := Follows{}
	flws.Source = User{
		ID:             msg.Source.Id,
		Username:       msg.Source.ScreenName,
		Name:           msg.Source.Name,
		Description:    msg.Source.Description,
		FollowersCount: msg.Source.FollowersCount,
		FollowingCount: msg.Source.FriendsCount,
	}
	flws.Target = User{
		ID:             msg.Target.Id,
		Username:       msg.Target.ScreenName,
		Name:           msg.Target.Name,
		Description:    msg.Target.Description,
		FollowersCount: msg.Target.FollowersCount,
		FollowingCount: msg.Target.FriendsCount,
	}
	flws.CreatedAt = msg.CreatedAt
	broadcast(flws)
}

func broadcast(data interface{}) {
	var l *list.List

	switch data.(type) {
	case Tweet:
		l = stream.tweetListeners
	case DM:
		l = stream.dmListeners
	case Follows:
		l = stream.followsListeners
	}

	stream.mu.Lock()
	for e := l.Front(); e != nil; e = e.Next() {
		lstnr := e.Value.(listener)
		ctx, cancel := context.WithTimeout(lstnr.ctx, 1*time.Minute)
		go func(ctx context.Context, c interface{}) {
			defer cancel()
			switch c := c.(type) {
			case chan Tweet:
				select {
				case <-ctx.Done():
				case c <- data.(Tweet):
				}
			case chan DM:
				select {
				case <-ctx.Done():
				case c <- data.(DM):
				}
			case chan Follows:
				select {
				case <-ctx.Done():
				case c <- data.(Follows):
				}
			}
		}(ctx, lstnr.c)
	}
	stream.mu.Unlock()
}
