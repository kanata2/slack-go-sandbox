package pull888

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/slack-go/slack"
)

func Main(c *slack.Client) error {
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	threadTs := os.Getenv("SLACK_THREAD_TS")
	_, _, err := c.PostMessage(channelID,
		slack.MsgOptionTS(threadTs),
		slack.MsgOptionBlocks(slack.NewActionBlock(
			"",
			slack.NewButtonBlockElement(
				"",
				"approve",
				slack.NewTextBlockObject(slack.PlainTextType, "Approve", false, false),
			),
			slack.NewButtonBlockElement(
				"",
				"deny",
				slack.NewTextBlockObject(slack.PlainTextType, "Deny", false, false),
			),
		)),
		slack.MsgOptionText("test", false),
	)
	if err != nil {
		return err
	}
	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(500)
			return
		}
		payload, err := url.QueryUnescape(string(buf)[8:])
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Println(payload)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	})
	return http.ListenAndServe(":3000", nil)
}
