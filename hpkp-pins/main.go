package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

func lolzNope(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var host string
	flag.StringVar(&host, "host", "bare hostname or IP of the remote server", "")
	flag.Parse()

	conn, err := tls.Dial("tcp", host+":443", nil)
	lolzNope(err)
	defer conn.Close()

	for _, c := range conn.ConnectionState().PeerCertificates {
		digest := sha256.Sum256(c.RawSubjectPublicKeyInfo)
		pin := base64.StdEncoding.EncodeToString(digest[:])
		fmt.Printf("%s : %s\n", pin, c.Subject.CommonName)
	}
}
