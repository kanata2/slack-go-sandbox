package issue881

import (
	"errors"
	"os"

	"github.com/slack-go/slack"
)

func Main(c *slack.Client) error {
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		return errors.New("CHANNEL_ID is required")
	}
	textBlockObject := slack.NewTextBlockObject("mrkdwn", "hello world!", true, false)
	sectionBlock := slack.NewSectionBlock(textBlockObject, nil, nil)

	blocks := []slack.Block{sectionBlock}
	blockOptions := slack.MsgOptionBlocks(blocks...)
	_, _, err := c.PostMessage(channelID, blockOptions)
	return err
}
