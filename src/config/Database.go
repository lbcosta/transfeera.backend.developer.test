package config

import (
	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
	Disconnect(db *gorm.DB)
}
