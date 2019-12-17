package command

var (
	help          Command
	inspirational Command
	pickone       Command
	rt            Command
	tweet         Command
)

func init() {
	help = Command{
		Name:   "help",
		Regexp: helpRegexp,
		Info:   helpInfo,
		Run:    helpRun,
	}
	inspirational = Command{
		Name:   "inspirational",
		Regexp: inspirationalRegexp,
		Info:   inspirationalInfo,
		Run:    inspirationalRun,
	}
	pickone = Command{
		Name:   "pickone",
		Regexp: pickoneRegexp,
		Info:   pickoneInfo,
		Run:    pickoneRun,
	}
	rt = Command{
		Name:   "rt",
		Regexp: rtRegexp,
		Info:   rtInfo,
		Run:    rtRun,
	}
	tweet = Command{
		Name:   "tweet",
		Regexp: tweetRegexp,
		Info:   tweetInfo,
		Run:    tweetRun,
	}

	// Add commands to the list.
	cmds = append(cmds,
		help,
		inspirational,
		pickone,
		// tweet,
		// rt,
	)

	for _, cmd := range cmds {
		helpDM += cmd.Info
	}
}
