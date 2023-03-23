package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresDatabase struct {
	DSN string
}

func NewPostgresDatabase() PostgresDatabase {
	return PostgresDatabase{
		DSN: fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s timezone=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			"America/Sao_Paulo"),
	}
}

func (p PostgresDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(p.DSN), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	return db, nil
}

func (p PostgresDatabase) Disconnect(db *gorm.DB) {
	database, _ := db.DB()
	database.Close()
}
