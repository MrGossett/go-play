package main

import (
	"fmt"
	"testing"
)

const str = "s"

var bslice = []byte(str)

func BenchmarkByteSliceGrow(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		bs = append(bs, bslice...)
	}
}

func BenchmarkStringConcatGrow(b *testing.B) {
	strAppend := ""
	for n := 0; n < b.N; n++ {
		strAppend += str
	}
}

func BenchmarkFmtSprintfGrow(b *testing.B) {
	strFmt := ""
	for n := 0; n < b.N; n++ {
		strFmt = fmt.Sprintf("%s%s", strFmt, str)
	}
}

func BenchmarkByteSliceDiscard(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		_ = append(bs, bslice...)
	}
}

func BenchmarkStringConcatDiscard(b *testing.B) {
	strAppend := ""
	for n := 0; n < b.N; n++ {
		_ = strAppend + str
	}
}

func BenchmarkFmtSprintfDiscard(b *testing.B) {
	strFmt := ""
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("%s%s", strFmt, str)
	}
}
