package console

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/elias19r/twitterbot/internal/garbage"
)

func garbageCmdMain(args []string) error {
	if len(args) <= 0 {
		return garbageCmdHelp()
	}
	switch args[0] {
	case "list", "ls":
		return garbageCmdList()
	case "add":
		return garbageCmdAdd(args[1:])
	case "remove", "rm":
		return garbageCmdRemove(args[1:])
	case "log":
		return garbageCmdLog()
	case "help", "?":
		return garbageCmdHelp()
	}
	return ErrSubcommandNotFound
}

func garbageCmdList() error {
	ids := garbage.GetTweetIDs()
	for _, i := range ids {
		fmt.Println(i)
	}
	return nil
}

func garbageCmdAdd(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	for _, arg := range args {
		id, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}
		garbage.AddTweetID(id)
	}
	return nil
}

func garbageCmdRemove(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	for _, arg := range args {
		id, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}
		garbage.RmTweetID(id)
	}
	return nil
}

func garbageCmdLog() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(": press ENTER to stop reading log")
	go logStdout(ctx, garbage.LogWriter)

	key := ""
	fmt.Scanln(&key)

	return nil
}

func garbageCmdHelp() error {
	text := `
` + garbageCmd.Regexp.String() + `
Available subcommands:

list|ls         Lists all tweet IDs.
add ID          Adds a tweet ID to the garbage list.
remove|rm ID    Removes a tweet ID from the garbage list.
log             Prints log information from garbage package.
help|?          Shows this help text.
`
	fmt.Println(strings.TrimSpace(text))
	return nil
}
