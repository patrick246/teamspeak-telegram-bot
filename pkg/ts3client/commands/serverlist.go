package commands

func Serverlist(uid, short, all, onlyOffline bool) Command {
	var options []string
	if uid {
		options = append(options, "uid")
	}
	if short {
		options = append(options, "short")
	}
	if all {
		options = append(options, "all")
	}
	if onlyOffline {
		options = append(options, "onlyoffline")
	}

	return Command{
		Name:       "serverlist",
		Parameters: nil,
		Options:    options,
	}
}
