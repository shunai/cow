package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestSendBodyChunked(t *testing.T) {
	testData := []struct {
		raw string
	}{
		{"1a; ignore-stuff-here\r\nabcdefghijklmnopqrstuvwxyz\r\n10\r\n1234567890abcdef\r\n0\r\n\r\n"},
		{"0\r\n\r\n"},
	}

	buf := make([]byte, bufSize)
	for _, td := range testData {
		r := bufio.NewReader(bytes.NewBufferString(td.raw))
		var w bytes.Buffer

		sendBodyChunked(buf, r, &w)
		if w.String() != td.raw {
			t.Errorf("sendBodyChunked wrong, raw data is:\n%sgot:%s\n", td.raw, w.String())
		}
	}
}
