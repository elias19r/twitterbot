package console

import (
	"context"
	"fmt"
	"strings"

	"github.com/elias19r/twitterbot/internal/command"
	"github.com/elias19r/twitterbot/internal/twitter"
)

func commandCmdMain(args []string) error {
	if len(args) <= 0 {
		return commandCmdHelp()
	}
	switch args[0] {
	case "list", "ls":
		return commandCmdList()
	case "run":
		return commandCmdRun(args[1:])
	case "log":
		return commandCmdLog()
	case "help", "?":
		return commandCmdHelp()
	}
	return ErrSubcommandNotFound
}

func commandCmdList() error {
	fmt.Println(strings.Join(command.List(), "\n"))
	return nil
}

func commandCmdRun(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	msg := command.Message{
		Type: "dm",
		User: twitter.Bot,
		Args: strings.Join(args, " "),
	}
	return command.FindCommand(msg)
}

func commandCmdLog() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(": press ENTER to stop reading log")
	go logStdout(ctx, command.LogWriter)

	key := ""
	fmt.Scanln(&key)

	return nil
}

func commandCmdHelp() error {
	text := `
` + commandCmd.Regexp.String() + `
Available subcommands:

list|ls     Lists all commands in command package.
run ARGS    Runs a command of command package with ARGS.
log         Prints log information from command package.
help|?      Shows this help text.
`
	fmt.Println(strings.TrimSpace(text))
	return nil
}
