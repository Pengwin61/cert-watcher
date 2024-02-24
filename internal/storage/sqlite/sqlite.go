package sqlite

import (
	"cert-watcher/internal/storage/model"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	Gorm *gorm.DB
}

func NewClient() *Client {
	db, err := gorm.Open(sqlite.Open("cert.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Certificate{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &Client{Gorm: db}
}
