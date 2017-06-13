package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
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

	gocode(host)
	// openssl(host)
	// certs := conn.ConnectionState().VerifiedChains

}

func gocode(host string) {
	conn, err := tls.Dial("tcp", host+":443", nil)
	lolzNope(err)
	defer conn.Close()

	for _, cert := range conn.ConnectionState().PeerCertificates {
		fmt.Printf("%#v\n\n", cert.Subject)
	}
}

func openssl(host string) {
	cmd := exec.Command("openssl", "s_client", "-showcerts", "-connect", host+":443")
	p, err := cmd.CombinedOutput()
	if err != nil {
		lolzNope(err)
	}

	begin := bytes.LastIndex(p, beginCertBytes)
	end := bytes.LastIndex(p, endCertBytes)

	if begin == -1 || end == -1 {
		lolzNope(errors.New("Could not find pull certificate"))
	}
	certBytes := p[begin : end+len(endCertBytes)]

	f, err := ioutil.TempFile("", "iss_cert")
	if err != nil {
		lolzNope(err)
	}
	defer f.Close()

	if _, err := f.Write(certBytes); err != nil {
		lolzNope(err)
	}

	cmd = exec.Command("openssl", "x509", "-in", f.Name(), "-fingerprint", "-noout")
	p, err = cmd.CombinedOutput()
	if err != nil {
		lolzNope(err)
	}

	p = bytes.Replace(p, []byte(":"), []byte{}, -1)
	p = bytes.TrimPrefix(p, []byte("SHA1 Fingerprint="))
	p = bytes.TrimSuffix(p, []byte("\n"))

	fmt.Println("OPENSSL (correct)", string(p))
}

// variables used to extract the cert from the cognito issuer
var (
	beginCertBytes = []byte("-----BEGIN CERTIFICATE-----")
	endCertBytes   = []byte("-----END CERTIFICATE-----")
)
