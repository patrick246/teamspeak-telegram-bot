package escape

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEscape(t *testing.T) {
	testIn := []byte("TeamSpeak ]|[ Server")
	testOut := []byte("TeamSpeak\\s]\\p[\\sServer")

	realOut := Escape(testIn)

	require.Equal(t, testOut, realOut)
}
