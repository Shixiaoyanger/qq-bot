package main

import (
	"qq-bot/bot"
	"qq-bot/view"
)

func main() {
	debug := false

	botAPI := bot.NewBotAPI(debug)
	view.ListenMessage()

	botAPI.MessageHandler()
}
