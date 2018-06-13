package main

import (
	"os"

	"github.com/TinyKitten/sugoibot/bot"
	"github.com/TinyKitten/sugoibot/env"
)

func main() {
	env.LoadEnv()

	slackToken := os.Getenv("SLACK_USER_TOKEN")
	slackbot := bot.NewBot(slackToken)
	slackbot.StartListenMessage()
}
