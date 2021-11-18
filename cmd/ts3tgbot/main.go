package main

import (
	"context"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/receivers"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/telegram"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client/commands"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3listener"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Print("connecting to server")
	client, err := ts3client.Connect(os.Getenv("TS3_SERVERQUERY_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("logging in")
	_, err = client.Command(context.Background(), commands.Login(os.Getenv("TS3_SERVERQUERY_USER"), os.Getenv("TS3_SERVERQUERY_PASSWORD")))
	if err != nil {
		log.Fatalf("login error: %v", err)
	}

	log.Print("selecting virtual server")
	_, err = client.Command(context.Background(), commands.UseByPort(os.Getenv("TS3_SERVERQUERY_VPORT"), false))
	if err != nil {
		log.Fatalf("server selection error: %v", err)
	}

	log.Print("configuring telegram client")
	tgReceiverChannel, err := strconv.ParseInt(os.Getenv("TELEGRAM_TARGET_CHANNEL"), 10, 64)
	if err != nil {
		log.Fatalf("expected telegram channel to be an integer, got error while parsing %v", err)
	}
	tgReceiver, err := telegram.NewClient(os.Getenv("TELEGRAM_TOKEN"), tgReceiverChannel)

	log.Print("starting ts3 listener")
	listener := ts3listener.NewListener(client, []receivers.OnlineReceiver{tgReceiver})
	err = listener.Run()
	if err != nil {
		log.Fatalf("error while running ts3 listener: %v", err)
	}
}
