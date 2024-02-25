package main

import (
	"cert-watcher/internal/core"
	"cert-watcher/internal/storage/model"
	"cert-watcher/internal/storage/sqlite"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

const (
	ROOT_PATH_CERT = "/Users/kirill/Documents/Git Repository/cert-watcher/cert/"
)

func main() {
	var certificates []model.Certificate

	db := sqlite.NewClient()

	for {
		certList := core.FindCertificate(ROOT_PATH_CERT)

		// Проходим по списку путей с сертификатами
		for _, path := range certList {

			// Проверка на истечение срока действия
			exp, cn := core.ExpirationCert(db, path)
			if exp {
				log.Printf("Сертификат %s истек срок действия\n", cn)
				continue
			}

			// Чтение сертификата
			cert := core.ReadCert(core.GetFullPathToFile(path))

			// Проверка сертификата
			cert = core.Check(cert)

			// Поиск записи в БД по CN на наличие, если записи нет, то запись создается в БД
			if err := db.Gorm.Where("cn = ?", cert.Subject.CommonName).First(&certificates).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Println("Запись не найдена:", cert.Subject.CommonName)
					create(db, cert)
				} else {
					fmt.Println("Произошла ошибка при поиске записи:", err)
				}
			}
		}
		time.Sleep(30 * time.Second)
		fmt.Println("Sleeping")
	}

}

func create(db *sqlite.Client, cert *x509.Certificate) {
	db.Gorm.Create(&model.Certificate{
		Before:       cert.NotBefore,
		After:        cert.NotAfter,
		Organization: cert.Subject.Organization[0],
		OU:           cert.Subject.OrganizationalUnit[0],
		CN:           cert.Subject.CommonName})
}

// db.Insert(cert.NotBefore.Format(time.DateOnly), cert.NotAfter.Format(time.DateOnly),
// 	cert.Subject.Organization[0], cert.Subject.OrganizationalUnit[0], cert.Subject.CommonName)
// db.Get()
// db.Update(8)
// db.Db.Close()

// var certificate model.Certificate
// db.Gorm.First(&certificate, "ou = ?", "kirill@MacBook-Air-Kirill.local (Кирилл)")
// db.Gorm.Model(&certificate).Update("ou", "test")
// db.Gorm.First(&model.Certificate{}, 2)
