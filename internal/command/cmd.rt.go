package command

import (
	"net/http"
	"regexp"

	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

var rtRegexp = regexp.MustCompile(`(?i)(^|[[:space:]]+)/(rt|(re|quote)tweet|quote|retuite)([[:space:]]+|$)`)
var rtInfo = `
/RT LINK

- Eu retuito o tuite indicado pelo link que você me mandar
- Por exemplo:
/RT https:` + "\u2063" + `//twitter.com/evaristocosta/status/951070839868452864
`

var tcoLinkRegexp = regexp.MustCompile(`https://t.co/[a-zA-Z0-9]{10}`)
var tweetLinkRegexp = regexp.MustCompile(`https://twitter.com/[a-zA-Z0-9_]{1,15}/status/[0-9]{1,20}`)

// rtRun tries to retweet (quoted tweet actually) a tweet pointed by a link.
func rtRun(msg Message) error {
	// Try to match and decode a t.co link.
	if !tcoLinkRegexp.MatchString(msg.Args) {
		return rtFailed(msg)
	}
	r, err := http.Get(tcoLinkRegexp.FindString(msg.Args))
	if err != nil {
		return rtFailed(msg)
	}
	url := r.Request.URL.String()

	// Try to match and extract a status ID.
	if !tweetLinkRegexp.MatchString(url) {
		return rtFailed(msg)
	}
	link := tweetLinkRegexp.FindString(url)

	// Send a quoted tweet.
	_, err = twitter.PostTweetCC(msg.User, link)
	if err != nil {
		return rtFailed(msg)
	}
	return nil
}

func rtFailed(msg Message) error {
	switch msg.Type {
	case "tweet":
		_, err := twitter.PostTweetAt(msg.User, config.RTLinkFailed)
		return err
	case "tweetreply":
		_, err := twitter.PostTweetReply(msg.User, config.RTLinkFailed, msg.ReplyID)
		return err
	case "dm":
		return twitter.SendDM(msg.User, config.RTLinkFailed)
	}
	return ErrInvalidMessage
}
