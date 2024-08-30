package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func New(dbLocation string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})

	if err != nil {
		log.Println("failed to connect database: ", err)
	}

	Migrate(db)
	Seed(db)

	return db
}
