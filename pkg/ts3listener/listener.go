package ts3listener

import (
	"context"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/receivers"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client/commands"
)

type Listener struct {
	client          *ts3client.Connection
	onlineReceivers []receivers.OnlineReceiver
}

func NewListener(client *ts3client.Connection, onlineReceivers []receivers.OnlineReceiver) *Listener {
	return &Listener{
		client:          client,
		onlineReceivers: onlineReceivers,
	}
}

func (l *Listener) Run() error {
	_, err := l.client.Command(context.Background(), commands.ServerNotifyRegister(commands.ServerNotifyServer, nil))
	if err != nil {
		return err
	}

	token, events := l.client.RegisterListener()
	for event := range events {
		switch event.Type {
		case "notifycliententerview":
			for _, receiver := range l.onlineReceivers {
				receiver.ReceiveOnline(event.Data[0]["client_nickname"])
			}
		}
	}
	l.client.UnregisterListener(token)
	return nil
}
