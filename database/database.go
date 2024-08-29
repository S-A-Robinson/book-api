package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Println("failed to connect database: ", err)
	}

	Seed(db)

	return db
}
