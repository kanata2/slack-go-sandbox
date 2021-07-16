package main

import (
	"log"
	"os"

	"github.com/kanata2/slack-go-sandbox/pull888"
	"github.com/slack-go/slack"
)

var (
	c *slack.Client
)

func init() {
	c = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"), slack.OptionDebug(true))
}

func main() {
	if err := pull888.Main(c); err != nil {
		log.Fatal(err)
	}
}
