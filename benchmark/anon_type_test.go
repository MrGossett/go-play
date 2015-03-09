package main

import (
	"encoding/json"
	"io/ioutil"

	"testing"
)

type definedType struct {
	str string
}

func BenchmarkDefinedType(b *testing.B) {
	var v definedType
	for n := 0; n < b.N; n++ {
		v = definedType{str: "foo"}
	}
	if v.str != "" {
		// no-op
	}
}

func BenchmarkInlineType(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		type inlineType struct {
			str string
		}
		v = inlineType{str: "foo"}
	}
	if v != nil {
		// no-op
	}
}

func BenchmarkAnonymousType(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		v = struct{ str string }{str: "foo"}
	}
	if v != nil {
		// no-op
	}
}

var enc = json.NewEncoder(ioutil.Discard)

func BenchmarkDefinedTypeEncoded(b *testing.B) {
	var v definedType
	for n := 0; n < b.N; n++ {
		v = definedType{str: "foo"}
		enc.Encode(v)
	}
}

func BenchmarkInlineTypeEncoded(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		type inlineType struct {
			str string
		}
		v = inlineType{str: "foo"}
		enc.Encode(v)
	}
}

func BenchmarkAnonymousTypeEncoded(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		v = struct{ str string }{str: "foo"}
		enc.Encode(v)
	}
}
