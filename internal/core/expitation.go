package core

import (
	"cert-watcher/internal/storage/model"
	"cert-watcher/internal/storage/sqlite"
	"fmt"
	"time"
)

func ExpirationCert(db *sqlite.Client, path string) (bool, string) {
	var certificate model.Certificate

	cert := ReadCert(GetFullPathToFile(path))

	db.Gorm.First(&certificate, "cn = ?", cert.Subject.CommonName)

	return exp(certificate)

}

func exp(cert model.Certificate) (bool, string) {
	// daysToExp := days * 24 * 60
	// d := time.Now().Add(time.Duration(daysToExp) * time.Minute)
	// fmt.Println("DAYS TO EXP: ", )

	deathDay, _ := time.ParseDuration("1849h") // 70 day

	last := cert.After.Truncate(time.Hour * 24).Sub(time.Now().Truncate(time.Hour * 24))
	fmt.Printf("сертификату %30s осталось дней: %d\n", cert.CN, last/time.Hour/24)

	if last < deathDay {
		return true, cert.CN
	}
	return false, cert.CN
}
