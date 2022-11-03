package main

import (
	"flag"
	"log"
	"telegram-bot-linksave/clients/telegram"
	event_consumer "telegram-bot-linksave/consumer/event-consumer"
	telegram2 "telegram-bot-linksave/events/telegram"
	"telegram-bot-linksave/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "file_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram2.New(
		telegram.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not spec")
	}
	return *token
}
