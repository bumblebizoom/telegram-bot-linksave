package main

import (
	"flag"
	"log"
	"telegram-bot-linksave/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	//consumer.Start(fetcher, processor)

	//fetcher = fetcher.New

	//processor = processor.New
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not spec")
	}
	return *token
}
