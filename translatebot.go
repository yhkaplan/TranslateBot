package main

import (
	"context" // Context is available in the main lib from v1.7
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/nlopes/slack"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

var ctx context.Context

func main() {
	// Set up Slack part
	fmt.Println("Starting up slackbot")
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// Set up google translate part
	fmt.Println("Starting up translation service")
	ctx = context.Background()
	apiKey := os.Getenv("TRANSLATE_API_KEY")
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// main logic loop
Loop: //Named loop just like Swift
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event received")
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s>", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					respond(rtm, ev, prefix, client)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid slack credentials")
				break Loop

			default:
				// Do nothing
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string, client *translate.Client) {

	// Remove unneeded @mention string
	targetStr := []string{strings.TrimPrefix(msg.Text, prefix)}

	langList, err := client.DetectLanguage(ctx, targetStr)
	if err != nil {
		fmt.Println("Could not detect language")
		return
	}

	inLang := langList[0][0].Language
	outLang := language.English //TODO: make this settable w/ envar or flag

	if inLang == language.English {
		outLang = language.Japanese //TODO: make this settable w/ envar or flag
	}

	trns, err := client.Translate(
		ctx,
		targetStr,
		outLang,
		&translate.Options{
			Source: inLang,
			Format: translate.Text,
		},
	)

	if err != nil {
		fmt.Println("Could not translate")
	}

	trnsReply := trns[0].Text
	fmt.Println(trnsReply)
	rtm.SendMessage(rtm.NewOutgoingMessage(trnsReply, msg.Channel))
}
