package commands

type ChannelListOptions struct {
	Topic        bool
	Flags        bool
	Voice        bool
	Limits       bool
	Icon         bool
	SecondsEmpty bool
}

func ChannelList(options ChannelListOptions) Command {
	var opts []string
	if options.Topic {
		opts = append(opts, "topic")
	}
	if options.Flags {
		opts = append(opts, "flags")
	}
	if options.Voice {
		opts = append(opts, "voice")
	}
	if options.Limits {
		opts = append(opts, "limits")
	}
	if options.Icon {
		opts = append(opts, "icon")
	}
	if options.SecondsEmpty {
		opts = append(opts, "secondsempty")
	}

	return Command{
		Name:       "channellist",
		Parameters: nil,
		Options:    opts,
	}
}
