package receivers

type OnlineReceiver interface {
	ReceiveOnline(user string)
}

type MessageReceiver interface {
	ReceiveMessage(source, user, content string)
}
