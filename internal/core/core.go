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

func Check(cert *x509.Certificate) *x509.Certificate {

	if cert.Subject.Organization == nil {
		cert.Subject.Organization = make([]string, 1)
		cert.Subject.Organization[0] = cert.Issuer.CommonName

	}

	if cert.Subject.OrganizationalUnit == nil {
		cert.Subject.OrganizationalUnit = make([]string, 1)
		cert.Subject.OrganizationalUnit[0] = cert.Issuer.Organization[0]
	}
	return cert
}
