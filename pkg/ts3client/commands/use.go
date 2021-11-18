package commands

func UseBySid(sid string, virtual bool) Command {
	var options []string
	if virtual {
		options = append(options, "virtual")
	}

	return Command{
		Name: "use",
		Parameters: map[string][]string{
			"sid": {sid},
		},
		Options: options,
	}
}

func UseByPort(port string, virtual bool) Command {
	var options []string
	if virtual {
		options = append(options, "virtual")
	}

	return Command{
		Name: "use",
		Parameters: map[string][]string{
			"port": {port},
		},
		Options: options,
	}
}
