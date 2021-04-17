package issue903

import (
	"fmt"

	"github.com/slack-go/slack"
)

func Main(c *slack.Client) error {
	params := slack.FileUploadParameters{
		File: "example.txt",
	}
	file, err := c.UploadFile(params)

	if err != nil {
		return err
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	return nil
}
