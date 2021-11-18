package commands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var commandTestResults = map[string]Command{
	"serverlist": {
		Name:       "serverlist",
		Parameters: nil,
		Options:    nil,
	},
	"clientlist -uid -away -groups": {
		Name:       "clientlist",
		Parameters: nil,
		Options:    []string{"uid", "away", "groups"},
	},
	"clientdbfind pattern=ScP": {
		Name:       "clientdbfind",
		Parameters: map[string][]string{"pattern": {"ScP"}},
		Options:    nil,
	},
	"clientdbfind pattern=FPMPSC6MXqXq751dX7BKV0JniSo= -uid": {
		Name:       "clientdbfind",
		Parameters: map[string][]string{"pattern": {"FPMPSC6MXqXq751dX7BKV0JniSo="}},
		Options:    []string{"uid"},
	},
	"clientkick clid=1|clid=2|clid=3 reasonid=5 reasonmsg=Go\\saway!": {
		Name: "clientkick",
		Parameters: map[string][]string{
			"reasonid":  {"5"},
			"reasonmsg": {"Go away!"},
			"clid":      {"1", "2", "3"},
		},
		Options: nil,
	},
	"channelmove cid=16 cpid=1 order=0": {
		Name:       "channelmove",
		Parameters: map[string][]string{"cid": {"16"}, "cpid": {"1"}, "order": {"0"}},
		Options:    nil,
	},
	"sendtextmessage msg=Hello\\sWorld! target=12 targetmode=2": {
		Name: "sendtextmessage",
		Parameters: map[string][]string{
			"targetmode": {"2"},
			"target":     {"12"},
			"msg":        {"Hello World!"},
		},
		Options: nil,
	},
}

func TestCommand_Serialize(t *testing.T) {
	for result, command := range commandTestResults {
		t.Run(result, func(t *testing.T) {
			serialized := command.Serialize()
			require.Equal(t, result, serialized)
		})
	}
}
