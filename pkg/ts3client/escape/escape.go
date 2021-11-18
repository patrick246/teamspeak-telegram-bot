package escape

import "fmt"

var escapeReplacements = map[byte][]byte{
	7:    {'\\', 'a'},
	8:    {'\\', 'b'},
	9:    {'\\', 't'},
	10:   {'\\', 'n'},
	11:   {'\\', 'v'},
	12:   {'\\', 'f'},
	13:   {'\\', 'r'},
	' ':  {'\\', 's'},
	'/':  {'\\', '/'},
	'\\': {'\\', '\\'},
	'|':  {'\\', 'p'},
}

var unescapeReplacements = map[byte]byte{
	'a':  7,
	'b':  8,
	't':  9,
	'n':  10,
	'v':  11,
	'f':  12,
	'r':  13,
	's':  ' ',
	'/':  '/',
	'\\': '\\',
	'p':  '|',
}

func Escape(in []byte) []byte {
	out := make([]byte, 0, len(in))
	for _, b := range in {
		if replacement, ok := escapeReplacements[b]; ok {
			out = append(out, replacement...)
		} else {
			out = append(out, b)
		}
	}
	return out
}

func Unescape(in []byte) ([]byte, error) {
	out := make([]byte, 0, len(in))

	isEscapeSequence := false
	for _, b := range in {
		if isEscapeSequence {
			isEscapeSequence = false
			replacement, ok := unescapeReplacements[b]
			if ok {
				out = append(out, replacement)
				continue
			}
			return nil, fmt.Errorf("expected Escape sequence, got %q", b)
		}

		if b == 92 {
			isEscapeSequence = true
			continue
		}
		out = append(out, b)
	}
	return out, nil
}
