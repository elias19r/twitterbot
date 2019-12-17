package twitter

import (
	"context"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/elias19r/twitterbot/internal/config"
)

// Follow follows a specified user.
func Follow(user User) error {
	if !dailyLimitFollowing.dec() {
		return ErrDailyLimitFollowing
	}
	_, err := api.FollowUser(user.Username)
	if err != nil {
		if err, ok := err.(*anaconda.ApiError); ok {
			if twitterError, ok := err.Decoded.First().(anaconda.TwitterError); ok {
				switch twitterError.Code {
				case 64, 89, 161, 226, 326:
					logger.Println(err)
					dailyLimitFollowing.drain()
					return ErrDailyLimitFollowing
				}
			}
		}
		return err
	}
	logger.Println("followed", user)
	return nil
}

// Unfollow unfollows a specified user name.
func Unfollow(user User) error {
	if !dailyLimitFollowing.dec() {
		return ErrDailyLimitFollowing
	}
	_, err := api.UnfollowUser(user.Username)
	if err != nil {
		if err, ok := err.(*anaconda.ApiError); ok {
			if twitterError, ok := err.Decoded.First().(anaconda.TwitterError); ok {
				switch twitterError.Code {
				case 64, 89, 161, 226, 326:
					logger.Println(err)
					dailyLimitFollowing.drain()
					return ErrDailyLimitFollowing
				}
			}
		}
		return err
	}
	logger.Println("unfollowed", user)
	return nil
}

func getUsersPage(cursor string, username string, get func(url.Values) (anaconda.UserCursor, error)) (page []User, nextCursor string) {
	v := url.Values{}
	v.Set("include_user_entities", "false")
	v.Set("skip_status", "true")
	v.Set("count", "200")
	v.Set("cursor", cursor)
	if username != "" {
		v.Set("screen_name", username)
	}
	c, err := get(v)
	if err != nil {
		logger.Println(err)
		return nil, ""
	}
	page = []User{}
	for _, u := range c.Users {
		page = append(page, User{
			ID:             u.Id,
			Username:       u.ScreenName,
			Name:           u.Name,
			Description:    u.Description,
			FollowersCount: u.FollowersCount,
			FollowingCount: u.FriendsCount,
			Location:       u.Location,
		})
	}
	nextCursor = c.Next_cursor_str
	logger.Println("len(page)", len(page))
	return
}

// GetFollowingPage returns a slice of User from one page of following.
func GetFollowingPage(cursor string, username string) ([]User, string) {
	return getUsersPage(cursor, username, api.GetFriendsList)
}

// GetFollowersPage returns a slice of User from one page of followers.
func GetFollowersPage(cursor string, username string) ([]User, string) {
	return getUsersPage(cursor, username, api.GetFollowersList)
}

func getUsersC(ctx context.Context, username string, getPage func(string, string) ([]User, string)) <-chan User {
	users := make(chan User, 200)

	go func(ctx context.Context, users chan<- User) {
		var page []User
		var cursor = "-1"
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				if cursor == "0" {
					break loop
				}
				page, cursor = getPage(cursor, username)
				if page == nil || cursor == "" {
					break loop
				}
				for _, user := range page {
					users <- user
				}
			}
		}
		close(users)
	}(ctx, users)

	return users
}

// GetFollowingC returns a buffered channel of twitter.User that I follow.
func GetFollowingC(ctx context.Context, username string) <-chan User {
	return getUsersC(ctx, username, GetFollowingPage)
}

// GetFollowersC returns a buffered channel of twitter.User that follow me.
func GetFollowersC(ctx context.Context, username string) <-chan User {
	return getUsersC(ctx, username, GetFollowersPage)
}

// WillUnfollow decides whether to unfollow a user.
func WillUnfollow(lu Lookup) error {
	if lu.Following && !lu.FollowedBy {
		return Unfollow(lu.User)
	}
	return ErrNoUnFollowPerformed
}

// FollowOrUnfollow decides whether to (un)follow a user.
func FollowOrUnfollow(lu Lookup) error {
	err := WillUnfollow(lu)
	if err == ErrNoUnFollowPerformed {
		return WillFollow(lu)
	}
	return err
}

// WillFollow decides whether to follow a user.
func WillFollow(lu Lookup) error {
	if !lu.Following && !lu.FollowingRequested {
		if lu.FollowedBy || followerCandidate(lu.User) {
			return Follow(lu.User)
		}
	}
	return ErrNoUnFollowPerformed
}

// followerCandidate checks following and followers and decide if user is
// likely to follow me back in the future.
func followerCandidate(user User) bool {
	a := user.FollowingCount
	b := user.FollowersCount

	if a >= config.FollowCandidateMin && b >= config.FollowCandidateMin && float64(a)/float64(b) >= config.FollowCandidateRatio {
		return true
	}
	return false
}
