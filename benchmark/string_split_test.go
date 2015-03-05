package main

import (
	"strings"
	"testing"
)

const strToSplit = "one,two"

func BenchmarkStringsSplitAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parts := strings.SplitN(strToSplit, ",", -1)
		_ = parts[0]
		_ = parts[1]
	}
}

func BenchmarkStringsSplitLimited(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parts := strings.SplitN(strToSplit, ",", 2)
		_ = parts[0]
		_ = parts[1]
	}
}

func BenchmarkStringsIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := strings.Index(strToSplit, ",")
		_ = strToSplit[:i]
		_ = strToSplit[i+1:]
	}
}

func BenchmarkStringsSplitAllValidated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parts := strings.SplitN(strToSplit, ",", -1)
		if len(parts) == 2 {
			_ = parts[0]
			_ = parts[1]
		}
	}
}

func BenchmarkStringsSplitLimitedValidated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parts := strings.SplitN(strToSplit, ",", 2)
		if len(parts) == 2 {
			_ = parts[0]
			_ = parts[1]
		}
	}
}

func BenchmarkStringsIndexValidated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := strings.Index(strToSplit, ",")
		if i >= 0 {
			_ = strToSplit[:i]
			_ = strToSplit[i+1:]
		}
	}
}
