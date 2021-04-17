package main

import (
	"log"
	"os"

	"github.com/kanata2/slack-go-tester-bot/issue903"
	"github.com/slack-go/slack"
)

var (
	c *slack.Client
)

func init() {
	c = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"), slack.OptionDebug(true))
}

func main() {
	if err := issue903.Main(c); err != nil {
		log.Fatal(err)
	}
}
