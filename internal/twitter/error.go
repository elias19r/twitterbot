package twitter

import (
	"errors"
)

// Errors.
var (
	ErrNoTrendsFound       = errors.New("no trends found")
	ErrNoStatusesFound     = errors.New("no statuses found")
	ErrNoFollowersFound    = errors.New("no followers found")
	ErrNoTimelineFound     = errors.New("no timeline found")
	ErrNoUnFollowPerformed = errors.New("no (un)follow performed")
	ErrNoUnMutePerformed   = errors.New("no (un)mute performed")

	ErrCannotDMYourself = errors.New("cannot send a DM to yourself")

	ErrInvalidTweetID = errors.New("invalid tweet ID")

	ErrDailyLimitFollowing = errors.New("following daily limit reached")
	ErrDailyLimitDM        = errors.New("DM daily limit reached")
	ErrDailyLimitTweet     = errors.New("tweet daily limit reached")
)
