package main

import (
	"cert-watcher/internal/core"
	"cert-watcher/internal/storage/sqlite"
	"time"
)

const (
	// PATH_CERT = "/etc/letsencrypt/live/example.com/fullchain.pem"
	PATH_CERT = "localhost.pem"
)

func main() {
	cert := core.ReadCert(PATH_CERT)

	db := sqlite.NewClient()

	// fmt.Println(cert)

	db.Insert(cert.NotBefore.Format(time.DateOnly), cert.NotAfter.Format(time.DateOnly),
		cert.Subject.Organization[0], cert.Subject.OrganizationalUnit[0], cert.Subject.CommonName)

	db.Get()

	db.Update(8)

	db.Db.Close()

}
