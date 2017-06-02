package main

import (
	"crypto/tls"
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
		fmt.Println(c.Subject.CommonName)
	}
}
