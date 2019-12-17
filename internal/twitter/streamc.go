package twitter

import (
	"container/list"
	"context"
)

func addStreamListener(ctx context.Context, c interface{}) {
	var l *list.List
	var e *list.Element

	stream.mu.Lock()
	switch c.(type) {
	case chan Tweet:
		logger.Println("added tweet listener")
		l = stream.tweetListeners
	case chan DM:
		logger.Println("added DM listener")
		l = stream.dmListeners
	case chan Follows:
		logger.Println("added follows listener")
		l = stream.followsListeners
	}
	e = l.PushBack(listener{ctx: ctx, c: c})
	stream.mu.Unlock()

	go func(ctx context.Context, l *list.List, e *list.Element) {
		select {
		case <-ctx.Done():
			stream.mu.Lock()
			l.Remove(e)
			logger.Println("removed listener")
			stream.mu.Unlock()
		}
	}(ctx, l, e)
}

// GetStreamTweetC returns a buffered channel of tweets messages from stream.
func GetStreamTweetC(ctx context.Context) <-chan Tweet {
	c := make(chan Tweet, 200)
	addStreamListener(ctx, c)
	return c
}

// GetStreamDMC returns a buffered channel of DM messages from stream.
func GetStreamDMC(ctx context.Context) <-chan DM {
	c := make(chan DM, 200)
	addStreamListener(ctx, c)
	return c
}

// GetStreamFollowsC returns a buffered channel of Follows messages from stream.
func GetStreamFollowsC(ctx context.Context) <-chan Follows {
	c := make(chan Follows, 200)
	addStreamListener(ctx, c)
	return c
}
