package db

import (
	"log"
	"os"

	"github.com/spctr-cc/backend-bugtrack/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Account{},
		&model.Ticket{},
	)
}
