package commands

func WhoAmI() Command {
	return Command{
		Name:       "whoami",
		Parameters: nil,
		Options:    nil,
	}
}
