package ts3client

import (
	"fmt"
	"strings"
)

type Event struct {
	Type string
	Data []map[string]string
}

func readEvent(line string) (Event, error) {
	firstSpace := strings.Index(line, " ")
	if firstSpace == -1 {
		return Event{}, fmt.Errorf("expected space in line %q, got none", line)
	}

	eventType := line[:firstSpace]

	dataPart := line[firstSpace+1:]
	data, err := readDataLine(dataPart)
	if err != nil {
		return Event{}, err
	}

	return Event{
		Type: eventType,
		Data: data,
	}, nil
}
