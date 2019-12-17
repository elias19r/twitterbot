package command

import (
	"regexp"

	"github.com/elias19r/twitterbot/internal/twitter"
)

var helpRegexp = regexp.MustCompile(`(?i)(^|[[:space:]]+)/(help|ajudar?|ajude)([[:space:]]+|$)`)
var helpInfo = `
/ajuda

- Eu mostro esse texto de ajuda
`

var helpLink = "https://t.co/2A3RK2TGZW"
var helpTweet = "Você pode conferir sobre mim aqui: " + helpLink
var helpDM = "Eu sou um bot simples que aceita os seguintes comandos via DM (mensagem), resposta ou se me marcar num tuite:\n"

// helpRun sends a help message with info about bot's commands.
func helpRun(msg Message) error {
	switch msg.Type {
	case "tweet", "mention":
		_, err := twitter.PostTweetAt(msg.User, helpTweet)
		return err
	case "tweetreply":
		_, err := twitter.PostTweetReply(msg.User, helpTweet, msg.ReplyID)
		return err
	case "dm":
		return twitter.SendDM(msg.User, helpDM)
	}
	return ErrInvalidMessage
}
