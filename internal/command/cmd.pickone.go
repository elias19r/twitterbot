package command

import (
	"math/rand"
	"regexp"
	"strings"

	"github.com/elias19r/twitterbot/internal/twitter"
	"github.com/elias19r/twitterbot/pkg/text"
)

var pickoneRegexp = regexp.MustCompile(`(?i)(^|[[:space:]]+)/(pick(one)?|escolh[ae](um)?)([[:space:]]+|$)`)
var pickoneInfo = `
/escolhe OPÇÃO ou OPÇÃO ou OPÇÃO

- Eu escolho uma das opções e respondo via DM, resposta ou tuite
- Por exemplo:
/escolhe É pavê ou pá comê?!kkkk
`

// Regex patterns for the "pick one" sentence.
var (
	option    = `(.+)`
	or        = `(?i:[\t\n\v\f\r .,!?'"-)]+ou[\t\n\v\f\r .,!?'"-(]+)`
	fromThree = regexp.MustCompile(option + or + option + or + option)
	fromTwo   = regexp.MustCompile(option + or + option)
)

// pickoneRun tries to pick one option out of two or three and build a
// response message.
func pickoneRun(msg Message) error {
	var results []string

	// Pick one from three options or pick one from two options.
	if ok := fromThree.MatchString(msg.Args); ok {
		results = fromThree.FindStringSubmatch(msg.Args)[1:]
	} else if ok = fromTwo.MatchString(msg.Args); ok {
		results = fromTwo.FindStringSubmatch(msg.Args)[1:]
	} else {
		return ErrInvalidArgs
	}

	// Treat options.
	for i := range results {
		// Strip pretexts.
		group := pretextRegexp.FindStringSubmatch(results[i])
		if len(group) >= 6 {
			if group[2] != "" {
				// Format: "pretext, option, pretext" or "option, pretext, pretext"
				if i != len(results)-1 {
					results[i] = group[2]
				} else {
					results[i] = group[1]
				}
				// Format: "pretext, option" or "option, pretext"
			} else {
				if i == 0 {
					results[i] = group[5]
				} else {
					results[i] = group[4]
				}
			}
		}
		// Remove garbage words and clear text.
		results[i] = garbageRegexp.ReplaceAllString(results[i], " ")
		results[i] = text.Clear(results[i])
	}

	// Build content and tweet/dm.
	intro := introTexts[rand.Intn(len(introTexts))]
	ending := endingTexts[rand.Intn(len(endingTexts))]
	choice := results[rand.Intn(len(results))]

	content := intro + choice + ending

	switch msg.Type {
	case "tweet":
		_, err := twitter.PostTweetAt(msg.User, content)
		return err
	case "tweetreply":
		_, err := twitter.PostTweetReply(msg.User, content, msg.ReplyID)
		return err
	case "dm":
		return twitter.SendDM(msg.User, content)
	}
	return ErrInvalidMessage
}

// Introduction texts for response.
var introTexts = []string{
	"",
	"Hmm... ",
	"Acho que ",
	"Ah, eu acho que ",
	"Por mim ",
	"Sei não, se pá ",
	"Cara, eu prefiro ",
	"Claro que ",
	"Eu diria que ",
}

// Ending texts for response.
var endingTexts = []string{
	"",
	".",
	"...",
	"!",
	"!!",
	"!!1",
	" né",
	" hahah",
	" :D",
	" -.-",
}

// Texts that are not part of an option.
// e.g:
//
// Diz aí o que acha pior: espirrar mijando,ou comendo, na casa da sogra
// haha ou dirigindo....quero ver sua resposta agora! #DúvidaDoSéculo
//
// => pretexts: [
//   "Diz aí o que acha pior:"
//   ", na casa da sogra haha"
//   "....quero ver sua resposta agora! #DúvidaDoSéculo'"
// ]
var pretextPatterns = []string{
	`^(.+?)[\n:.,!?>-]+(.+?)[\n:.,!?>-]+(.+)$`,
	`|`,
	`^(.+?)[\n:.,!?>-]+(.+)$`,
}
var pretextRegexp = regexp.MustCompile(strings.Join(pretextPatterns, ""))

// Verbs or nouns that are not part of an option.
var garbagePatterns = []string{
	`(?i)`,
	`[[:space:]]*`,
	`(`,
	`ou`,
	`|`,
	`(pick( ?one)?|choose( ?one)?|select( ?one)?|escolh[ea]( ?um)?|selecion[ea]( ?um)?)`,
	`|`,
	`(deseja|prefere|gosta)`,
	`|`,
	`(qual|o ?qu[eê])`,
	`|`,
	`(pa?r[ao] |por )?(voc[eê]|vc|o?c[eê]|n[oó]i?[sz]?|a? ?gente)`,
	`)`,
	`[[:space:]]*`,
}
var garbageRegexp = regexp.MustCompile(strings.Join(garbagePatterns, ""))
