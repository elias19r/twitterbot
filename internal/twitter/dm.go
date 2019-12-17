package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/elias19r/twitterbot/internal/config"
)

// SendDM sends a DM to a specified user.
func SendDM(user User, content string) error {
	if user.Username == config.BotUsername {
		return ErrCannotDMYourself
	}
	if !dailyLimitDM.dec() {
		return ErrDailyLimitDM
	}
	_, err := api.PostDMToScreenName(content, user.Username)
	if err != nil {
		if err, ok := err.(*anaconda.ApiError); ok {
			if twitterError, ok := err.Decoded.First().(anaconda.TwitterError); ok {
				switch twitterError.Code {
				case 64, 89, 205, 226, 326:
					logger.Println(err)
					dailyLimitDM.drain()
					return ErrDailyLimitDM
				}
			}
		}
		return err
	}
	logger.Printf("sent DM to %s: %q\n", user, content)
	return nil
}

// SendDMThanks sends a default thanks DM message.
func SendDMThanks(user User) error {
	content := config.GetDefaultThanksReply()
	return SendDM(user, content)
}
