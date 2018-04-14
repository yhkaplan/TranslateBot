package main

import (
	"context" // Context is available in the main lib from v1.7
	"fmt"
	"log" // For panic func, might delete later
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
	ctx = context.Background() //This might not be the right way to initialize
	// TODO: any way to avoid absolut paths??
	//var err error //This feels like a hack, but it's to allow the statement below to work
	// with the globally defined "client" rather than creating a new one
	client, err := translate.NewClient(ctx, option.WithServiceAccountFile("/Users/joshk/.keys/keyfile.json"))
	if err != nil {
		log.Panicln("Fatal error: %s", err) //TODO: is this best way to handle?
	}

	// main logic loop
Loop:
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
	targetStr := strings.TrimPrefix(msg.Text, prefix)

	trns, err := client.Translate(
		ctx,
		[]string{targetStr},
		language.Japanese,
		&translate.Options{
			Source: language.English,
			Format: translate.Text,
		})

	if err != nil {
		fmt.Println("Could not translate")
	}

	trnsReply := trns[0].Text
	fmt.Println(trnsReply)
	rtm.SendMessage(rtm.NewOutgoingMessage(trnsReply, msg.Channel))
}
