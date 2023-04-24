package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(Events <-chan *slacker.CommandEvent) {
	for event := range Events {
		fmt.Println("Events Command")
		fmt.Println(event.Command)
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4975528175348-5149770015606-GlYOqc2TrYT5L0w8DS3XAaYD")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A054D9VR7C6-5149344370278-da49f616cb3a78e244e28f1ecc8cb4c61da2beb6ce373ba06d401d1c473b6e0a")
	bot := slacker.NewClient(os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_BOT_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
