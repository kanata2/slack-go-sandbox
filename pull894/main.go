package pull894

import (
	"log"

	"github.com/slack-go/slack"
)

func Main(c *slack.Client) error {
	b, err := c.GetBotInfo("")
	if err != nil {
		return err
	}
	log.Printf("%v\n", b)
	return nil
}
