package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kr/pretty"
)

func main() {
	buf := new(bytes.Buffer)
	io.Copy(buf, os.Stdin)

	block, _ := pem.Decode(buf.Bytes())

	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%# v\n", pretty.Formatter(csr))

	fmt.Printf("signature error: %s\n", csr.CheckSignature())
}
