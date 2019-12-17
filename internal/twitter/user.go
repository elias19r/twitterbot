package twitter

import (
	"math/rand"
	"net/url"
	"strings"
)

// GetRandomUser returns username of a random user from trends.
// TODO: same logic as GetRandomTweet(), should merge both functions.
func GetRandomUser(woeid int64) (string, error) {
	// Get trends for a location WOIED.
	loc, err := api.GetTrendsByPlace(woeid, url.Values{})
	if err != nil {
		return "", err
	}
	// Pick a trend at random from the list of trends.
	if len(loc.Trends) <= 0 {
		return "", ErrNoTrendsFound
	}
	trend := loc.Trends[rand.Intn(len(loc.Trends))]

	// Search for tweets with that trend query.
	v := url.Values{}
	v.Set("include_entities", "false")
	v.Set("lang", "pt")
	v.Set("count", "100")
	sr, err := api.GetSearch(trend.Query, v)
	if err != nil {
		return "", err
	}
	if len(sr.Statuses) <= 0 {
		return "", ErrNoStatusesFound
	}
	// Pick a tweet at random from search response.
	tweet := sr.Statuses[rand.Intn(len(sr.Statuses))]

	return tweet.User.ScreenName, nil
}

// GetLookupOne looks up one user.
func GetLookupOne(user User) (Lookup, error) {
	lus, err := GetLookup(user)
	if len(lus) > 0 {
		return lus[0], err
	}
	return Lookup{}, err
}

// GetLookup returns a slice of Lookup from user.
func GetLookup(users ...User) ([]Lookup, error) {
	usernames := []string{}
	for _, u := range users {
		usernames = append(usernames, u.Username)
	}
	v := url.Values{}
	v.Set("screen_name", strings.Join(usernames, ","))
	friendships, err := api.GetFriendshipsLookup(v)
	if err != nil {
		return nil, err
	}
	lookups := []Lookup{}
	for _, fs := range friendships {
		lu := Lookup{}
		// Set User.
		for _, u := range users {
			if fs.Id == u.ID {
				lu.User = u
				break
			}
		}
		// Set connections.
		for _, conn := range fs.Connections {
			if conn == "followed_by" {
				lu.FollowedBy = true
			}
			if conn == "following" {
				lu.Following = true
			}
			if conn == "following_requested" {
				lu.FollowingRequested = true
			}
			if conn == "none" {
				lu.None = true
			}
			if conn == "blocking" {
				lu.Blocking = true
			}
			if conn == "muting" {
				lu.Muting = true
			}
		}
		lookups = append(lookups, lu)
	}
	return lookups, nil
}
