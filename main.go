package main

import (
	"context" // Context is available in the main lib from v1.7
	"fmt"
	"os"

	"cloud.google.com/go/translate"
	"github.com/nlopes/slack"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Starting up slackbot")
	slack_token := os.Getenv("Slack_TOKEN")
	rtm := api.NewRTM()

	ctx := context.Background()
	client, err := translate.NewClient(ctx, opts)
	if err != nil {
		fmt.Println("Fatal error: %s", err)
		Panic
	}
	select {
	case msg := <-rtm.IncomingEvents:
		ev := msg.Data.(type)

		switch ev {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			fmt.Sprintf("%s", info.User.ID)
		}
	}
}
