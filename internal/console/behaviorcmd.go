package console

import (
	"context"
	"fmt"
	"strings"

	"github.com/elias19r/twitterbot/internal/behavior"
)

func behaviorCmdMain(args []string) error {
	if len(args) <= 0 {
		return behaviorCmdHelp()
	}
	switch args[0] {
	case "list", "ls":
		return behaviorCmdList()
	case "start":
		return behaviorCmdStart(args[1:])
	case "startnow":
		return behaviorCmdStartNow(args[1:])
	case "startall":
		return behaviorCmdStartAll()
	case "stop":
		return behaviorCmdStop(args[1:])
	case "stopall":
		return behaviorCmdStopAll()
	case "status":
		return behaviorCmdStatus(args[1:])
	case "log":
		return behaviorCmdLog(args[1:])
	case "help", "?":
		return behaviorCmdHelp()
	}
	return ErrSubcommandNotFound
}

func behaviorCmdList() error {
	fmt.Println(strings.Join(behavior.Info(), "\n"))
	return nil
}

func behaviorCmdStart(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	b, err := behavior.Get(args[0])
	if err != nil {
		return err
	}
	return b.Start(false)
}

func behaviorCmdStartNow(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	b, err := behavior.Get(args[0])
	if err != nil {
		return err
	}
	return b.Start(true)
}

func behaviorCmdStartAll() error {
	behavior.StartAll()
	return nil
}

func behaviorCmdStop(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	b, err := behavior.Get(args[0])
	if err != nil {
		return err
	}
	return b.Stop()
}

func behaviorCmdStopAll() error {
	behavior.StopAll()
	return nil
}

func behaviorCmdStatus(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	b, err := behavior.Get(args[0])
	if err != nil {
		return err
	}
	fmt.Println(b.Status())
	return nil
}

func behaviorCmdLog(args []string) error {
	if len(args) <= 0 {
		return ErrInvalidArgs
	}
	b, err := behavior.Get(args[0])
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(": press ENTER to stop reading log")
	go logStdout(ctx, b.LogWriter)

	key := ""
	fmt.Scanln(&key)

	return nil
}

func behaviorCmdHelp() error {
	text := `
` + behaviorCmd.Regexp.String() + `
Available subcommands:

list|ls          Lists all behaviors along with its status info.
start NAME       Starts specified behavior.
startnow NAME    Starts specified behavior skipping bootup delay.
startall         Starts all behaviors.
stop NAME        Stops specified behavior.
stopall          Stops all behaviors.
status NAME      Shows status of specified behavior.
log NAME         Prints log information from specified behavior.
help|?           Shows this help text.
`
	fmt.Println(strings.TrimSpace(text))
	return nil
}
