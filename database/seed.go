package database

import (
	"books-api/controllers"
	"books-api/models"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"os"
)

func Seed(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{})

	if err != nil {
		log.Println("failed to seed the database: ", err)
	}

	SeedBooks(db)
}

func SeedBooks(db *gorm.DB) {
	data, err := os.ReadFile("database/initial_books.json")

	if err != nil {
		log.Println("error reading initial books", err)
	}
	books := &[]models.Book{}
	json.Unmarshal(data, books)
	for _, book := range *books {
		controllers.AddBook(db, &book)
	}
	//TODO now that we have read the file, unmarshal the json and get the books from it to populate the table
}
