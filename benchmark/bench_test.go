package main

import (
	"fmt"
	"testing"
)

const str = "s"

var bslice = []byte(str)

func BenchmarkByteSlice(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		bs = append(bs, bslice...)
	}
}

func BenchmarkStringConcat(b *testing.B) {
	strAppend := ""
	for n := 0; n < b.N; n++ {
		strAppend += str
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	strFmt := ""
	for n := 0; n < b.N; n++ {
		strFmt = fmt.Sprintf("%s%s", strFmt, str)
	}
}
