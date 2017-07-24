package splitter

import (
	"strings"
	"unicode/utf8"
)

// Splitter provides a bufio.SplitFunc for use with a bufio.Scanner that will
// detect tokens comprised exclusively of runes in the Splitter's Charset.
type Splitter struct {
	// Charset is a string of all the possible characters that can be used
	// in a target token
	Charset string
}

const hex = "1234567890abcdefABCDEF"
const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

// HexSplitter sniffs for hex-encoded strings
var HexSplitter = Splitter{Charset: hex}

// B64Splitter sniffs for base64-encoding strings
var B64Splitter = Splitter{Charset: b64}

// NewSplitter returns a Splitter that uses the provided charset with duplicate
// characters removed.
func NewSplitter(charset string) Splitter {
	var s Splitter
	uniq := map[rune]struct{}{}
	for _, r := range charset {
		if _, ok := uniq[r]; ok {
			continue
		}
		s.Charset += string(r)
		uniq[r] = struct{}{}
	}
	return s
}

// ScanCharset is a split function for a Splitter that returns tokens comprised
// exclusively of runes in the Splitter's Charset.
func (s Splitter) ScanCharset(data []byte, atEOF bool) (int, []byte, error) {
	// Skip leading runes not in s.Charset.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if strings.ContainsRune(s.Charset, r) {
			break
		}
	}
	// Scan until next rune not in s.Charset, marking end of token.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !strings.ContainsRune(s.Charset, r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated token.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
