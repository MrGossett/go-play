package main

import (
	"encoding/json"
	"io/ioutil"

	"testing"
)

type definedType struct {
	Str string `json:"string"`
}

func BenchmarkDefinedType(b *testing.B) {
	var v definedType
	for n := 0; n < b.N; n++ {
		v = definedType{Str: "foo"}
	}
	if v.Str != "" {
		// no-op
	}
}

func BenchmarkInlineType(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		type inlineType struct {
			Str string `json:"string"`
		}
		v = inlineType{Str: "foo"}
	}
	if v != nil {
		// no-op
	}
}

func BenchmarkAnonymousType(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		v = struct {
			Str string `json:"string"`
		}{
			Str: "foo",
		}
	}
	if v != nil {
		// no-op
	}
}

var enc = json.NewEncoder(ioutil.Discard)

func BenchmarkDefinedTypeEncoded(b *testing.B) {
	var v definedType
	for n := 0; n < b.N; n++ {
		v = definedType{Str: "foo"}
		enc.Encode(v)
	}
}

func BenchmarkInlineTypeEncoded(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		type inlineType struct {
			Str string `json:"string"`
		}
		v = inlineType{Str: "foo"}
		enc.Encode(v)
	}
}

func BenchmarkAnonymousTypeEncoded(b *testing.B) {
	var v interface{}
	for n := 0; n < b.N; n++ {
		v = struct {
			Str string `json:"string"`
		}{
			Str: "foo",
		}
		enc.Encode(v)
	}
}
