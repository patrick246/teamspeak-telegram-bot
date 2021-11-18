package ts3client

import (
	"fmt"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client/escape"
	"strings"
)

type Result struct {
	Id        string
	Message   string
	Data      []map[string]string
	ErrorData map[string]string
}

func (r Result) Success() bool {
	return r.Id == "0"
}

func (r Result) Error() string {
	return r.Message
}

func readErrorLine(errorLine string) (map[string]string, error) {
	parts := strings.Split(errorLine, " ")
	parts = parts[1:]

	errorData := make(map[string]string)
	for _, p := range parts {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("expected key-value pair, got %q", p)
		}
		unescapedValue, err := escape.Unescape([]byte(kv[1]))
		if err != nil {
			return nil, err
		}
		errorData[kv[0]] = string(unescapedValue)
	}
	return errorData, nil
}

func readDataLine(dataLine string) ([]map[string]string, error) {
	var data []map[string]string
	elements := strings.Split(dataLine, "|")
	for _, element := range elements {
		dataMap := make(map[string]string)
		parts := strings.Split(element, " ")
		for _, p := range parts {
			kv := strings.SplitN(p, "=", 2)
			if len(kv) == 1 {
				dataMap[kv[0]] = ""
				continue
			}

			unescapedValue, err := escape.Unescape([]byte(kv[1]))
			if err != nil {
				return nil, err
			}
			dataMap[kv[0]] = string(unescapedValue)
		}

		data = append(data, dataMap)
	}
	return data, nil
}
