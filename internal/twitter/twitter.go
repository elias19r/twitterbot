package twitter

import (
	"bytes"
	"log"

	"github.com/elias19r/twitterbot/internal/config"

	"github.com/ChimeraCoder/anaconda"
)

var api *anaconda.TwitterApi

// Bot represents a User with Bot username from config.
var Bot = User{
	ID:       config.BotID,
	Username: config.BotUsername,
}

// LogWriter can be used to read log from this package.
var (
	LogWriter      = new(bytes.Buffer)
	logger         = log.New(LogWriter, "twitter: ", 0)
	AnacondaLogger = &anacondaLogger{log: log.New(LogWriter, log.Prefix(), log.LstdFlags)}
)

type anacondaLogger struct {
	log *log.Logger
}

func (l anacondaLogger) Fatal(items ...interface{})               { l.log.Fatal(items...) }
func (l anacondaLogger) Fatalf(s string, items ...interface{})    { l.log.Fatalf(s, items...) }
func (l anacondaLogger) Panic(items ...interface{})               { l.log.Panic(items...) }
func (l anacondaLogger) Panicf(s string, items ...interface{})    { l.log.Panicf(s, items...) }
func (l anacondaLogger) Critical(items ...interface{})            { l.log.Print(items...) }
func (l anacondaLogger) Criticalf(s string, items ...interface{}) { l.log.Printf(s, items...) }
func (l anacondaLogger) Error(items ...interface{})               { l.log.Print(items...) }
func (l anacondaLogger) Errorf(s string, items ...interface{})    { l.log.Printf(s, items...) }
func (l anacondaLogger) Warning(items ...interface{})             { l.log.Print(items...) }
func (l anacondaLogger) Warningf(s string, items ...interface{})  { l.log.Printf(s, items...) }
func (l anacondaLogger) Notice(items ...interface{})              { l.log.Print(items...) }
func (l anacondaLogger) Noticef(s string, items ...interface{})   { l.log.Printf(s, items...) }
func (l anacondaLogger) Info(items ...interface{})                { l.log.Print(items...) }
func (l anacondaLogger) Infof(s string, items ...interface{})     { l.log.Printf(s, items...) }
func (l anacondaLogger) Debug(items ...interface{})               { l.log.Print(items...) }
func (l anacondaLogger) Debugf(s string, items ...interface{})    { l.log.Printf(s, items...) }
