package garbage

import (
	"bytes"
	"log"
	"sync"

	"github.com/elias19r/twitterbot/pkg/int64slice"
)

// List of tweet IDs to be removed.
var (
	mu       sync.Mutex // mu guards tweetIDs
	tweetIDs int64slice.Int64Slice
)

// LogWriter can be used to read log from this package.
var (
	LogWriter = new(bytes.Buffer)
	logger    = log.New(LogWriter, "garbage: ", 0)
)

// AddTweetID adds one or more tweet IDs to the list.
func AddTweetID(ids ...int64) int {
	count := 0
	mu.Lock()
	for _, i := range ids {
		if i == 0 {
			continue
		}
		tweetIDs.Insert(i)
		logger.Println("added tweet ID", i)
		count++
	}
	mu.Unlock()
	return count
}

// RmTweetID removes one or more tweet IDs from the list.
func RmTweetID(ids ...int64) int {
	count := 0
	mu.Lock()
	for _, i := range ids {
		if tweetIDs.Remove(i) {
			logger.Println("removed tweet ID", i)
			count++
		}
	}
	mu.Unlock()
	return count
}

// GetTweetIDs returns a copy of tweetIDs.
func GetTweetIDs() int64slice.Int64Slice {
	mu.Lock()
	defer mu.Unlock()
	return tweetIDs.Copy()
}
