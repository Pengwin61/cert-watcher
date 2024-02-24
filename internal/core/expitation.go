package core

import (
	"cert-watcher/internal/storage/model"
	"cert-watcher/internal/storage/sqlite"
	"fmt"
	"time"
)

func ExpirationCert(db *sqlite.Client, certlist []string) {

	var certificate model.Certificate
	db.Gorm.First(&certificate, "cn = ?", "uds.bosch-ru.ru")

	exp(certificate)

}

func exp(cert model.Certificate) bool {
	// 3 days
	deathline, _ := time.ParseDuration("-72h")

	if time.Since(cert.After) > deathline {
		fmt.Println("cert expired")
		return true
	}
	return false
}
