package command

import (
	"regexp"
	"strings"

	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

var tweetRegexp = regexp.MustCompile(`(?i)(^|[[:space:]]+)/(tweet|tuite)([[:space:]]+|$)`)
var tweetInfo = `
/tuite TEXTO

- Eu tuito o texto para você e te mando o link via DM
- Por exemplo:
/tuite Pq o petróleo foi ao terapeuta? Pq ele tava no fundo do poço kkkk
`

// tweetRun posts a tweet with content (args) passed to it in message.
func tweetRun(msg Message) error {
	msg.Args = strings.TrimSpace(msg.Args)
	if msg.Args == "" {
		return ErrInvalidArgs
	}
	_, err := twitter.PostTweetCC(msg.User, msg.Args)
	if err != nil {
		return tweetFailed(msg)
	}
	return nil
}

func tweetFailed(msg Message) error {
	switch msg.Type {
	case "tweet":
		_, err := twitter.PostTweetAt(msg.User, config.TweetFailed)
		return err
	case "tweetreply":
		_, err := twitter.PostTweetReply(msg.User, config.TweetFailed, msg.ReplyID)
		return err
	case "dm":
		return twitter.SendDM(msg.User, config.TweetFailed)
	}
	return ErrInvalidMessage
}
