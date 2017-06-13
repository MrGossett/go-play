package main

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kr/pretty"
)

func main() {
	buf := new(bytes.Buffer)
	io.Copy(buf, os.Stdin)

	list, err := x509.ParseDERCRL(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%# v\n", pretty.Formatter(list))
}
