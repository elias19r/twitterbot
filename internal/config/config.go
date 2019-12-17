package config

import (
	"math/rand"
	"strings"
	"time"
)

// config package defines the configuration values.
var (
	// Bot info.
	BotID       = int64(955484548053323776)
	BotUsername = "tiozaodamassa"

	// Behavior settings (intervals and delay in seconds).
	AutoFollowsMinInterval = 24 * 60 * 60
	AutoFollowsMaxInterval = 30 * 60 * 60

	AutoLikesMinInterval = 24 * 60 * 60
	AutoLikesMaxInterval = 30 * 60 * 60

	FollowCandidateMin   = 500
	FollowCandidateRatio = 0.99

	CrawlerMax                  = 250
	CrawlerMinInterval          = 24 * 60 * 60
	CrawlerMaxInterval          = 30 * 60 * 60
	CrawlerBootupDelay          = 3 * 60 * 60
	CrawlerRunningTimeout       = 4 * 60 * 60
	CrawlerDelayBetweenRequests = 5 * 60

	FollowerMinInterval          = 36 * 60 * 60
	FollowerMaxInterval          = 48 * 60 * 60
	FollowerBootupDelay          = 1 * 60 * 60
	FollowerRunningTimeout       = 5 * 60 * 60
	FollowerDelayBetweenRequests = 5 * 60

	FollowingMinInterval          = 7 * 24 * 60 * 60
	FollowingMaxInterval          = 9 * 24 * 60 * 60
	FollowingBootupDelay          = 5 * 24 * 60 * 60
	FollowingRunningTimeout       = 5 * 60 * 60
	FollowingDelayBetweenRequests = 5 * 60

	GCMinInterval = 24 * 60 * 60
	GCMaxInterval = 30 * 60 * 60
	GCBootupDelay = 2 * 60 * 60

	LikeMinInterval = 24 * 60 * 60
	LikeMaxInterval = 30 * 60 * 60
	LikeBootupDelay = 5 * 60
	LikeWOEID       = int64(455827) // 455827 is WOEID for São Paulo, SP
	LikeMinDelay    = 15 * 60
	LikeMaxDelay    = 30 * 60
	LikeMinN        = 12
	LikeMaxN        = 24

	VagueMinInterval = 24 * 60 * 60
	VagueMaxInterval = 30 * 60 * 60
	VagueBootupDelay = 3 * 60 * 60

	CommandMinInterval = 24 * 60 * 60
	CommandMaxInterval = 30 * 60 * 60

	RandomUserWOEID = int64(23424768) // 23424768 is WOEID for Brazil

	// Delay between throttled queries.
	Delay      = 5
	BufferSize = int64(3)

	// Replies.
	DefaultReplies = []string{
		"Oi! digite /ajuda para ver meus comandos",
		"Opa, e aí! digite /ajuda para saber mais sobre mim",
		"Blz? Se precisar de algo, digite /ajuda",
		"Como vai? digitando /ajuda eu mostro o que posso fazer",
		"Digite /ajuda para mostrar comandos ;)",
		"Para saber mais sobre mim, digite /ajuda ;)",
		"Se precisar de algo, digite /ajuda :)",
		"Digitando /ajuda eu mostro o que sei fazer :)",
		"Eu aceito alguns comandos, digite /ajuda para ver ;)",
		"Veja os comandos que eu sei, digite /ajuda :)",
	}
	DefaultThanksReplies = []string{
		"Obrigado!",
		"Valeu :)",
		"Opa, obrigadão",
		"👍",
	}
	EmojiOK = []string{
		"😉", "🙂", "😎", "🤓",
		"👍", "👏", "✌", "🤙",
		"💛",
		"🆗", "🆒", "✔", "☑",
	}
	TweetFailed         = "Não consegui tuitar o que você pediu :|"
	TweetReplyFailed    = "Não consegui responder seu tuite :|"
	RTLinkFailed        = "Não parece um link de um tuite..."
	GetRandomUserFailed = "Não consegui encontrar um usuário no momento..."

	// TODO: use golang templates?
	followingStartPT = `
🇧🇷 🔄 Começando meu SDV do dia!

Automaticamente seguindo de volta aqueles que me seguiram ontem :)

... e claro, dando unf em quem me deu unf ontem!

#SDV #{{weekdayPT}}DetremuraSDV #TiozãoDaMassaSDV

{{date}}
`

	followingStartEN = `
🇺🇸 🔄 Starting my F4F of the day!

I'm automatically following back those that followed me yesterday :)

... and of course, unfollowing those that unfollowed me yesterday!

#IFB #F4F #followback #TiozãoDaMassaF4F

{{date}}
`
	dailyLimitFollowing = `
🇧🇷 🔄 Oops! Atingi meu limite de seguidas (follow) por dia!
Mas amanhã voltarei a seguir vocês ;)
	
🇺🇸 🔄 Oops! I've reached my following limit per day!
But I will be back tomorrow to follow you ;)

{{date}}
`

	dailyLimitDM = `
🇧🇷 📨 Oops! Atingi meu limite de mensagens (DM) por dia!
Mas amanhã voltarei a responder DM :)

🇺🇸 📨 Oops! I've reached my messages (DM) limit per day!
But tomorrow I will reply DM again :)

{{date}}
`
	dailyLimitTweet = `
🇧🇷 🐦 Oops! Atingi meu limite de tuites por dia!
Mas amanhã voltarei a tuitar, me aguardem!

🇺🇸 🐦 Oops! I've reached my tweets limit per day!
But tomorrow I will tweet again, wait for me!

{{date}}
`
)

// GetDefaultReply picks a random string from DefaultReplies and return.
func GetDefaultReply() string {
	if len(DefaultReplies) <= 0 {
		return "Type /help"
	}
	return DefaultReplies[rand.Intn(len(DefaultReplies))]
}

// GetDefaultThanksReply picks a random string from DefaultThanksReplies and return.
func GetDefaultThanksReply() string {
	if len(DefaultThanksReplies) <= 0 {
		return "👍"
	}
	return DefaultThanksReplies[rand.Intn(len(DefaultThanksReplies))]
}

// GetEmojiOK returns an emoji that "says" OK.
func GetEmojiOK() string {
	if len(EmojiOK) <= 0 {
		return "👍"
	}
	return EmojiOK[rand.Intn(len(EmojiOK))]
}

// FollowingStartEN returns a followingStartEN string with placeholders replaced.
func FollowingStartEN() string {
	return replace(followingStartEN)
}

// FollowingStartPT returns a followingStartPT string with placeholders replaced.
func FollowingStartPT() string {
	return replace(followingStartPT)
}

// DailyLimitFollowing returns dailyFollowingLimit string with placeholders replaced.
func DailyLimitFollowing() string {
	return replace(dailyLimitFollowing)
}

// DailyLimitDM returns dailyDMLimit string with placeholders replaced.
func DailyLimitDM() string {
	return replace(dailyLimitDM)
}

// DailyLimitTweet returns dailyTweetsLimit string with placeholders replaced.
func DailyLimitTweet() string {
	return replace(dailyLimitTweet)
}

func replace(str string) string {
	str = strings.Replace(str, "{{weekdayPT}}", weekdayPT(), -1)
	str = strings.Replace(str, "{{weekday}}", weekday(), -1)
	str = strings.Replace(str, "{{date}}", date(), -1)
	return str
}

func weekdayPT() string {
	t := time.Now()
	days := [...]string{
		"Domingo",
		"Segunda",
		"Terça",
		"Quarta",
		"Quinta",
		"Sexta",
		"Sábado",
	}
	return days[t.Weekday()]
}

func weekday() string {
	t := time.Now()
	return t.Weekday().String()
}

func date() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
