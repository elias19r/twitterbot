package console

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

// Command represents a command recognized by console.
type Command struct {
	Regexp *regexp.Regexp
	Main   func([]string) error
	Help   string
}

// List of available console commands.
var cmds []Command

// Run starts console, scanning stdio for lines.
func Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("console> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			continue
		}
		args := strings.Fields(string(line))
		if len(args) <= 0 {
			continue
		}
		fmt.Println("")
		err = findCommand(args)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
}

func findCommand(args []string) error {
	if len(args) <= 0 {
		return nil
	}
	for _, cmd := range cmds {
		if !cmd.Regexp.MatchString(args[0]) {
			continue
		}
		return cmd.Main(args[1:])
	}
	return ErrCommandNotFound
}

func logStdout(ctx context.Context, in io.Reader) {
	reader := bufio.NewReader(in)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				fmt.Println(string(line))
			}
		}
	}
}
