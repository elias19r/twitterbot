package console

import (
	"fmt"
	"strings"
)

func helpCmdMain(args []string) error {
	return helpCmdHelp()
}

func helpCmdHelp() error {
	text := `
` + strings.Replace(helpCmd.Regexp.String(), `\`, "", -1) + `
Available commands:

`
	for _, cmd := range cmds {
		text += fmt.Sprintf("%-17s%-s\n", strings.Replace(cmd.Regexp.String(), `\`, "", -1), cmd.Help)
	}
	fmt.Println(strings.TrimSpace(text))
	return nil
}
