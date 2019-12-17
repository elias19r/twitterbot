package console

import "regexp"

var (
	logCmd      Command
	behaviorCmd Command
	commandCmd  Command
	garbageCmd  Command
	twitterCmd  Command
	helpCmd     Command
)

func init() {
	logCmd = Command{
		Regexp: regexp.MustCompile(`log`),
		Main:   logCmdMain,
		Help:   "Prints all log information.",
	}
	behaviorCmd = Command{
		Regexp: regexp.MustCompile(`behavior|bhvr`),
		Main:   behaviorCmdMain,
		Help:   "Interacts with behaviors.",
	}
	commandCmd = Command{
		Regexp: regexp.MustCompile(`command|cmd`),
		Main:   commandCmdMain,
		Help:   "Interacts with command package.",
	}
	garbageCmd = Command{
		Regexp: regexp.MustCompile(`garbage|gc`),
		Main:   garbageCmdMain,
		Help:   "Interacts with garbage package.",
	}
	twitterCmd = Command{
		Regexp: regexp.MustCompile(`twitter|tt`),
		Main:   twitterCmdMain,
		Help:   "Interacts with twitter package.",
	}
	helpCmd = Command{
		Regexp: regexp.MustCompile(`help|\?`),
		Main:   helpCmdMain,
		Help:   "Show this help info.",
	}

	// Add console commands to the list.
	cmds = append(cmds,
		logCmd,
		behaviorCmd,
		commandCmd,
		garbageCmd,
		twitterCmd,
		helpCmd,
	)
}
