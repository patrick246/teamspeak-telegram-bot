package commands

import (
	"fmt"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client/escape"
	"sort"
	"strings"
)

type Command struct {
	Name       string
	Parameters map[string][]string
	Options    []string
}

func (c Command) Serialize() string {
	serializedParameters := c.serializeParameters()
	serializedOptions := c.serializeOptions()

	serializedCommand := c.Name
	if serializedParameters != "" {
		serializedCommand += " " + serializedParameters
	}
	if serializedOptions != "" {
		serializedCommand += " " + serializedOptions
	}

	return serializedCommand
}

func (c Command) serializeParameters() string {
	var params []string
	for k, v := range c.Parameters {
		params = append(params, c.serializeParameterValues(k, v))
	}

	sort.Slice(params, func(i, j int) bool {
		return params[i] < params[j]
	})

	return strings.Join(params, " ")
}

func (c Command) serializeParameterValues(key string, values []string) string {
	var params []string
	for _, v := range values {
		params = append(params, fmt.Sprintf("%s=%s", key, escape.Escape([]byte(v))))
	}
	return strings.Join(params, "|")
}

func (c Command) serializeOptions() string {
	var options []string
	for _, o := range c.Options {
		options = append(options, fmt.Sprintf("-%s", o))
	}
	return strings.Join(options, " ")
}
