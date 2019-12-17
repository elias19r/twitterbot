package crawler

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/elias19r/twitterbot/internal/behavior"
	"github.com/elias19r/twitterbot/internal/config"
	"github.com/elias19r/twitterbot/internal/twitter"
)

// New returns a new *Behavior with Crawler action.
func New(name string) *behavior.Behavior {
	min := config.CrawlerMinInterval
	max := config.CrawlerMaxInterval
	delay := config.CrawlerBootupDelay

	b := behavior.New(name, min, max, action)
	b.SetBootupDelay(delay)

	return b
}

// action listens to a channel of followers (twitter.User) from another user
// picked at random from a list. For every follower of that user, it decides
// whether the bot should follow it as well, until a total of max followers.
func action(ctx context.Context, logger *log.Logger) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.CrawlerRunningTimeout)*time.Second)
	defer cancel()
	defer func() { count = 0 }()

	if name == "" {
		for i := 0; i < len(usernames); i++ {
			// Pick a user to crawl its followers.
			name = usernames[rand.Intn(len(usernames))]

			// Check if user still exists.
			_, err := twitter.GetLookupOne(twitter.User{Username: name})
			if err != nil {
				name = ""
				continue
			}
			break
		}
		if name == "" {
			logger.Println(ErrNoUsername)
			return
		}
	}
	logger.Println("username", name)

	followers := twitter.GetFollowersC(ctx, name)
	users := []twitter.User{}
	for u := range followers {
		users = append(users, u)
		if len(users) >= 100 { // Lookup chunks of 100 users.
			err := lookup(ctx, logger, users)
			users = users[:0]
			if err != nil {
				logger.Println(err)
				break
			}
		}
	}
	lookup(ctx, logger, users)
}

// SetName sets package level variable name.
func SetName(s string) {
	name = s
}

func lookup(ctx context.Context, logger *log.Logger, users []twitter.User) error {
	// Check if Following behavior is running, because Crawler MUST wait
	// for Following behavior to finish.
	b, err := behavior.Get("following")
	if err == nil && b.Started() && !b.Idle() {
		// NOTE: in the worst case, Crawler will wait five times the
		// FollowingRunningTimeout, but no problem.
		// It is important that Following and Crawler
		// behaviors DO NOT compete.
		time.Sleep(time.Duration(4*config.FollowingRunningTimeout) * time.Second)
	}

	lus, _ := twitter.GetLookup(users...)
	for _, lu := range lus {
		err := twitter.WillFollow(lu)
		select {
		case <-ctx.Done():
			return ErrCrawlerMax
		default:
			if err == nil {
				twitter.WillMute(lu)
				count++ // count is package level variable.
				if count >= max {
					return ErrCrawlerMax
				}
			}
			if err == twitter.ErrDailyLimitFollowing {
				return err
			}
		}
	}
	name = ""
	time.Sleep(time.Duration(config.CrawlerDelayBetweenRequests) * time.Second)
	return nil
}

var (
	name      = ""
	max       = config.CrawlerMax
	count     = 0
	usernames = []string{
		`1M__000`,
		`1N2`,
		`2i2i__`,
		`3boox_i`,
		`5vil_`,
		`__R5m`,
		`_mylad`,
		`Abrazos4u`,
		`am102358`,
		`arlenesg`,
		`bata451`,
		`BDRASh3r`,
		`BisbaleraTorres`,
		`CandySlimeAsmr`,
		`cxj_9`,
		`drnoursalem`,
		`EmpireBuild88`,
		`Erteqa___`,
		`followhelpxs`,
		`FollowM_Back_1M`,
		`FollowtrickBR_1`,
		`FollowTrickFab`,
		`freeftrick`,
		`GainWithSD`,
		`Help__F`,
		`HelpGain__`,
		`iaa077`,
		`Jef_Sdv`,
		`jn_shine`,
		`junges_lilia`,
		`MAZ___2`,
		`MichaelNike8`,
		`mlathalrooh25`,
		`mlk___m`,
		`MonicaR926`,
		`N5__e`,
		`nNuhaaa`,
		`ProjetoDMeFT`,
		`ProjetoFTColor`,
		`ProjTrickHelp`,
		`Queeen_Mai`,
		`realTonkaTattie`,
		`rebarbosa28`,
		`scofieldbx`,
		`shj___1`,
		`Tober007`,
		`TOPLUXURYGOALS`,
		`TrocaFollowBR`,
		`tucc4304444`,
		`vip_mon_cute`,
		`VisaAmericanaEU`,
		`xw_1m`,
		`Y__f4`,
		`ZH__7`,
		`ZR__Q22`,
		`qvuui`,
		`angel_sabina3`,
		`YEMEN2_2`,
		`shwnmevdes`,
	}
)
