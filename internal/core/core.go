package core

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func ReadCert(f string) *x509.Certificate {
	r, _ := os.ReadFile(f)
	block, _ := pem.Decode(r)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return cert
}
