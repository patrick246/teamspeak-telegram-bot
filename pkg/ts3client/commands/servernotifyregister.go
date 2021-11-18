package commands

type EventType string

const (
	ServerNotifyServer      EventType = "server"
	ServerNotifyChannel     EventType = "channel"
	ServerNotifyTextServer  EventType = "textserver"
	ServerNotifyTextChannel EventType = "textchannel"
	ServerNotifyTextPrivate EventType = "textprivate"
)

func ServerNotifyRegister(eventType EventType, id *string) Command {
	parameters := map[string][]string{
		"event": {string(eventType)},
	}

	if id != nil {
		parameters["id"] = []string{*id}
	}

	return Command{
		Name:       "servernotifyregister",
		Parameters: parameters,
		Options:    nil,
	}
}
