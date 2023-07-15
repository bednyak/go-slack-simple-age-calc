package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5527719707623-5560254338663-g7YwF68bouTwfcerFlgYycO1")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05GZ3DA0AG-5574682559538-9dda27051c24a984730a8546b6a81780268be5d04ba608f9a00baacc1e7f2de7")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	var arrayOfExamples []string = []string{"My year of birth is 1994"}

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    arrayOfExamples,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println(err)
			}

			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
