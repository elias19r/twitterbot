package twitter

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/elias19r/twitterbot/internal/config"
)

func init() {
	// Get keys from environment variables.
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("consumer key/secret and access token/secret required")
	}

	// Initialize client API.
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(AnacondaLogger)

	// Turn on anaconda throttling of queries.
	api.EnableThrottling(time.Duration(config.Delay)*time.Second, config.BufferSize)

	// Seed rand package.
	rand.Seed(time.Now().UnixNano() + 171036)
}
