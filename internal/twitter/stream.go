package twitter

import (
	"container/list"
	"context"
	"net/url"
	"sync"

	"github.com/ChimeraCoder/anaconda"
)

// stream holds a *anaconda.Stream and a list of listeners to which messages
// will be sent.
var stream = struct {
	stream           *anaconda.Stream
	mu               sync.Mutex // mu guards (tweet|dm|follows)Listeners.
	tweetListeners   *list.List
	dmListeners      *list.List
	followsListeners *list.List
}{
	tweetListeners:   list.New(),
	dmListeners:      list.New(),
	followsListeners: list.New(),
}

// listener have a context and a channel.
type listener struct {
	ctx context.Context
	c   interface{} // holds any channel
}

// StartStream creates a UserStream and starts reading from it.
func StartStream() {
	stream.stream = api.UserStream(url.Values{})
	logger.Println("reading stream")
	go readStream()
}

// StopStream stops the anaconda.Stream.
func StopStream() {
	if stream.stream != nil {
		logger.Println("stopped reading stream")
		stream.stream.Stop()
	}
}

// readStream reads from anaconda.Stream.C channel, and handles a message
// according to its type.
func readStream() {
	for msg := range stream.stream.C {
		switch msg := msg.(type) {
		case anaconda.EventFollow:
			go handleFollows(msg)
		case anaconda.Tweet:
			go handleTweet(msg)
		case anaconda.DirectMessage:
			go handleDM(msg)
		}
	}
}
