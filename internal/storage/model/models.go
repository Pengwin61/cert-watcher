package model

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Before       time.Time
	After        time.Time
	Organization string
	OU           string
	CN           string `gorm:"primaryKey"`
}
