package console

import (
	"context"
	"fmt"
	"strings"

	"github.com/elias19r/twitterbot/internal/twitter"
)

func twitterCmdMain(args []string) error {
	if len(args) <= 0 {
		return twitterCmdHelp()
	}
	switch args[0] {
	case "log":
		return twitterCmdLog()
	case "help", "?":
		return twitterCmdHelp()
	}
	return ErrSubcommandNotFound
}

func twitterCmdLog() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(": press ENTER to stop reading log")
	go logStdout(ctx, twitter.LogWriter)

	key := ""
	fmt.Scanln(&key)

	return nil
}

func twitterCmdHelp() error {
	text := `
` + twitterCmd.Regexp.String() + `
Available subcommands:

log       Prints log information from twitter package.
help|?    Shows this help text.
`
	fmt.Println(strings.TrimSpace(text))
	return nil
}
