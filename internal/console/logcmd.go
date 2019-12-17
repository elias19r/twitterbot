package console

import (
	"context"
	"fmt"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/command"
	"github.com/elias19r/twitterbot/internal/garbage"
	"github.com/elias19r/twitterbot/internal/twitter"
)

func logCmdMain(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(": press ENTER to stop reading log")

	bhvrNames := behavior.List()
	for _, b := range bhvrNames {
		bhvr, err := behavior.Get(b)
		if err != nil {
			return err
		}
		go logStdout(ctx, bhvr.LogWriter)
	}
	go logStdout(ctx, command.LogWriter)
	go logStdout(ctx, garbage.LogWriter)
	go logStdout(ctx, twitter.LogWriter)

	key := ""
	fmt.Scanln(&key)

	return nil
}
