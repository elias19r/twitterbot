package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/elias19r/twitterbot/internal/behavior"
	autofollows "github.com/elias19r/twitterbot/internal/behavior.autofollows"
	autolikes "github.com/elias19r/twitterbot/internal/behavior.autolikes"
	cmddm "github.com/elias19r/twitterbot/internal/behavior.cmddm"
	cmdtweet "github.com/elias19r/twitterbot/internal/behavior.cmdtweet"
	crawler "github.com/elias19r/twitterbot/internal/behavior.crawler"
	follower "github.com/elias19r/twitterbot/internal/behavior.follower"
	following "github.com/elias19r/twitterbot/internal/behavior.following"
	gc "github.com/elias19r/twitterbot/internal/behavior.gc"
	like "github.com/elias19r/twitterbot/internal/behavior.like"
	vague "github.com/elias19r/twitterbot/internal/behavior.vague"
	"github.com/elias19r/twitterbot/internal/console"
	"github.com/elias19r/twitterbot/internal/twitter"
)

func main() {
	twitter.StartStream()

	behavior.Add(autofollows.New("autofollows"))
	behavior.Add(autolikes.New("autolikes"))
	behavior.Add(cmddm.New("cmddm"))
	behavior.Add(cmdtweet.New("cmdtweet"))
	behavior.Add(crawler.New("crawler"))
	behavior.Add(follower.New("follower"))
	behavior.Add(following.New("following"))
	behavior.Add(gc.New("gc"))
	behavior.Add(like.New("like"))
	behavior.Add(vague.New("vague"))

	go console.Run()

	// Wait for SIGINT and SIGTERM (hit Ctrl-C) then stop stream.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	twitter.StopStream()
}
