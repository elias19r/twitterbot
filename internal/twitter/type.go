package twitter

import "fmt"

// User is a subset of anaconda.User type.
type User struct {
	ID             int64
	Username       string
	Name           string
	Description    string
	FollowersCount int
	FollowingCount int
	Location       string
}

// String implements Stringer interface.
func (u User) String() string {
	return u.Username
}

// Tweet is a subset of anaconda.Tweet type.
type Tweet struct {
	ID        int64
	Type      string // "tweet", "retweet", "mention", "tweetreply"
	User      User
	Content   string
	CreatedAt string
}

// String implements Stringer interface.
func (t Tweet) String() string {
	return fmt.Sprintf("%s: (%s) %q", t.Type, t.User, t.Content)
}

// DM is a subset of anaconda.DirectMessage type.
type DM struct {
	From      User
	To        User
	Content   string
	CreatedAt string
}

// String implements Stringer interface.
func (dm DM) String() string {
	return fmt.Sprintf("dm: (%s > %s) %q", dm.From, dm.To, dm.Content)
}

// Follows is a subset of anaconda.EventFollow type
type Follows struct {
	Source    User
	Target    User
	CreatedAt string
}

// String implements Stringer interface.
func (f Follows) String() string {
	return fmt.Sprintf("follows: %s > %s", f.Source, f.Target)
}

// Lookup is a subset of anaconda.Friendship type.
type Lookup struct {
	User               User
	FollowedBy         bool
	Following          bool
	FollowingRequested bool
	None               bool
	Blocking           bool
	Muting             bool
}

// String implements Stringer interface.
func (lu Lookup) String() string {
	return fmt.Sprintf("lookup: %s, %t, %t", lu.User, lu.FollowedBy, lu.Following)
}
