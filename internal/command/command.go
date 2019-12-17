package command

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/elias19r/twitterbot/internal/twitter"
)

// LogWriter can be used to read log from this package.
var (
	LogWriter = new(bytes.Buffer)
	logger    = log.New(LogWriter, "command: ", 0)
)

// cmds is a list of all available commands.
var cmds []Command

// Command represents a bot command.
type Command struct {
	Name   string
	Regexp *regexp.Regexp
	Info   string
	Run    func(Message) error
}

// Message represents a command message.
type Message struct {
	Type    string // "tweet", "retweet", "mention", "tweetreply" or "dm"
	User    twitter.User
	Command string
	Args    string
	ReplyID int64
}

// String implements Stringer interface.
func (m Message) String() string {
	return fmt.Sprintf("%s/%s: (%s) %q", m.Type, m.Command, m.User, m.Args)
}

// Handle handles a tweet or a dm with a command.
func Handle(data interface{}) error {
	msg := Message{}

	switch data := data.(type) {
	case twitter.Tweet:
		msg.Type = data.Type
		msg.User = data.User
		msg.Args = data.Content
		msg.ReplyID = data.ID
	case twitter.DM:
		msg.Type = "dm"
		msg.User = data.From
		msg.Args = data.Content
	default:
		return ErrInvalidMessage
	}

	return FindCommand(msg)
}

// List returns a list with available commands.
func List() []string {
	names := []string{}
	for _, cmd := range cmds {
		names = append(names, cmd.Name)
	}
	return names
}

// FindCommand iterates through command (cmds) list trying to match and
// run a command.
func FindCommand(msg Message) error {
	switch msg.Type {
	case "dm", "mention", "tweetreply":
		for _, cmd := range cmds {
			if !cmd.Regexp.MatchString(msg.Args) {
				continue
			}
			msg.Command = cmd.Name
			msg.Args = cmd.Regexp.ReplaceAllString(msg.Args, " ")
			msg.Args = strings.Replace(msg.Args, "@"+twitter.Bot.Username, " ", -1)
			logger.Println(msg)
			return cmd.Run(msg)
		}
		return ErrCommandNotFound
	default:
		return nil
	}
}
