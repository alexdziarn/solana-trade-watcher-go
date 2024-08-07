package main

import (
	bot "discordbot/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot.BotToken = os.Getenv("DISCORD_TOKEN")
	bot.Channel = os.Getenv("DISCORD_CHANNEL")
	bot.Run()
}
