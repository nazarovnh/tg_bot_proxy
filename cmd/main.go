package main

import (
	"context"
	"os"
)

var (
	BotToken = os.Getenv("BotToken")

	WebhookURL = "https://525f2cb5.ngrok.io"
)

func startTaskBot(ctx context.Context) error {
	// сюда пишите ваш код
	return nil
}

func main() {
	configPath := flag.String("c", "./cmd/go-telegram-bot-example/config.yaml", "path to go-telegram-bot-example config")
	flag.Parse()

	err := startTaskBot(context.Background())
	if err != nil {
		panic(err)
	}
}
