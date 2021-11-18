package ts3client

import (
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadErrorLine(t *testing.T) {
	errorLine := "error id=1033 msg=server\\sis\\snot\\srunning"

	errorData, err := readErrorLine(errorLine)
	require.NoError(t, err)

	require.Nil(t, deep.Equal(errorData, map[string]string{
		"id":  "1033",
		"msg": "server is not running",
	}))
}

func TestReadDataLine(t *testing.T) {
	dataLine := "cid=1 pid=0 channel_order=0 channel_name=Forum total_clients=1 channel_needed_subscribe_power=0|cid=10 pid=0 channel_order=1 channel_name=CNC-Werkstatt total_clients=0 channel_needed_subscribe_power=0|cid=2 pid=0 channel_order=10 channel_name=Digitallabor total_clients=0 channel_needed_subscribe_power=0|cid=3 pid=0 channel_order=2 channel_name=Elektronikwerkstatt total_clients=0 channel_needed_subscribe_power=0|cid=4 pid=0 channel_order=3 channel_name=Holzwerkstatt total_clients=0 channel_needed_subscribe_power=0|cid=5 pid=0 channel_order=4 channel_name=Open\\sSpace total_clients=0 channel_needed_subscribe_power=0|cid=9 pid=0 channel_order=5 channel_name=Kuschelbox\\s(AFK) total_clients=1 channel_needed_subscribe_power=0|cid=6 pid=0 channel_order=9 channel_name=Nähwerkstatt total_clients=0 channel_needed_subscribe_power=0|cid=7 pid=0 channel_order=6 channel_name=Medienwerkstatt total_clients=0 channel_needed_subscribe_power=0|cid=11 pid=7 channel_order=0 channel_name=Tonkabine total_clients=0 channel_needed_subscribe_power=0|cid=13 pid=0 channel_order=7 channel_name=Factorio total_clients=0 channel_needed_subscribe_power=0|cid=16 pid=0 channel_order=13 channel_name=Wortwitze total_clients=0 channel_needed_subscribe_power=0|cid=18 pid=0 channel_order=16 channel_name=Muss\\slernen total_clients=0 channel_needed_subscribe_power=0"

	data, err := readDataLine(dataLine)
	require.NoError(t, err)

	require.Nil(t, deep.Equal(data, []map[string]string{
		{"channel_name": "Forum", "channel_needed_subscribe_power": "0", "channel_order": "0", "cid": "1", "pid": "0", "total_clients": "1"},
		{"channel_name": "CNC-Werkstatt", "channel_needed_subscribe_power": "0", "channel_order": "1", "cid": "10", "pid": "0", "total_clients": "0"},
		{"channel_name": "Digitallabor", "channel_needed_subscribe_power": "0", "channel_order": "10", "cid": "2", "pid": "0", "total_clients": "0"},
		{"channel_name": "Elektronikwerkstatt", "channel_needed_subscribe_power": "0", "channel_order": "2", "cid": "3", "pid": "0", "total_clients": "0"},
		{"channel_name": "Holzwerkstatt", "channel_needed_subscribe_power": "0", "channel_order": "3", "cid": "4", "pid": "0", "total_clients": "0"},
		{"channel_name": "Open Space", "channel_needed_subscribe_power": "0", "channel_order": "4", "cid": "5", "pid": "0", "total_clients": "0"},
		{"channel_name": "Kuschelbox (AFK)", "channel_needed_subscribe_power": "0", "channel_order": "5", "cid": "9", "pid": "0", "total_clients": "1"},
		{"channel_name": "Nähwerkstatt", "channel_needed_subscribe_power": "0", "channel_order": "9", "cid": "6", "pid": "0", "total_clients": "0"},
		{"channel_name": "Medienwerkstatt", "channel_needed_subscribe_power": "0", "channel_order": "6", "cid": "7", "pid": "0", "total_clients": "0"},
		{"channel_name": "Tonkabine", "channel_needed_subscribe_power": "0", "channel_order": "0", "cid": "11", "pid": "7", "total_clients": "0"},
		{"channel_name": "Factorio", "channel_needed_subscribe_power": "0", "channel_order": "7", "cid": "13", "pid": "0", "total_clients": "0"},
		{"channel_name": "Wortwitze", "channel_needed_subscribe_power": "0", "channel_order": "13", "cid": "16", "pid": "0", "total_clients": "0"},
		{"channel_name": "Muss lernen", "channel_needed_subscribe_power": "0", "channel_order": "16", "cid": "18", "pid": "0", "total_clients": "0"},
	}))
}
