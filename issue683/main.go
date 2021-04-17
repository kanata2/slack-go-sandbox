package issue683

import (
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

func Main(c *slack.Client) error {
	http.HandleFunc("/slack/sc", func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err := c.OpenView(s.TriggerID, slack.ModalViewRequest{
			Type: "modal",
			Title: &slack.TextBlockObject{
				Type: "plain_text",
				Text: "Hi there",
			},
			Blocks: slack.Blocks{},
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%v", resp)
	})
	http.HandleFunc("/slack/ic", func(w http.ResponseWriter, r *http.Request) {
	})
	return http.ListenAndServe(":3000", nil)
}
