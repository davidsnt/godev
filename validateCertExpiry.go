package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	certFile := "path/to/cert"
	fd, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Println("Error in opening cert file")
	}
	certblock, _ := pem.Decode(fd)
	cert, err := x509.ParseCertificate(certblock.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	certValidUntil := cert.NotAfter
	t := time.Now()
	currDate := t.UTC()
	duration := int(certValidUntil.Sub(currDate).Hours() / float64(24))
	if duration <= 0 {
		fmt.Println("Expired Cert")
	} else {
		fmt.Println("Cert Valid Until: ", certValidUntil.Format("2006-01-02"))
	}
}
