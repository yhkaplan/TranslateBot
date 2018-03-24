package main

import (
	"context" // Context is available in the main lib from v1.7
	"fmt"
	"log" // For panic func, might delete later
	"os"

	"cloud.google.com/go/translate"
	"github.com/nlopes/slack"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Starting up slackbot")
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	ctx := context.Background()
	// TODO: any way to avoid absolut paths??
	client, err := translate.NewClient(ctx, option.WithServiceAccountFile("/Users/joshk/.keys/keyfile.json"))
	if err != nil {
		log.Panicln("Fatal error: %s", err) //TODO: is this best way to handle?
	}

	// What's the diff between ctx and context??
	trns, err := client.Translate(ctx,
		[]string{"この翻訳はどうだろう？うまくいけるのかな？"},
		language.English,
		&translate.Options{
			Source: language.Japanese,
			Format: translate.Text,
		})

	if err != nil {
		fmt.Println("Could not translate")
	}

	fmt.Println(trns[0].Text)

	rtm.SendMessage(rtm.NewOutgoingMessage(trns[0].Text, "#general"))
}
