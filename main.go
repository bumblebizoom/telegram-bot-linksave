package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	//tgClient = telegram.New(token)

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
