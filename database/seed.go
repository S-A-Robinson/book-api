package database

import (
	"books-api/controllers"
	"books-api/models"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"os"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{}, &models.Author{}, &models.AuthorBook{})

	if err != nil {
		log.Println("failed to migrate the database: ", err)
	}
}

func Seed(db *gorm.DB) {
	SeedBooks(db)
	SeedAuthors(db)
	SeedAuthorBooks(db)
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
}

func SeedAuthors(db *gorm.DB) {
	data, err := os.ReadFile("database/initial_authors.json")

	if err != nil {
		log.Println("error reading initial authors", err)
	}
	authors := &[]models.Author{}
	json.Unmarshal(data, authors)
	for _, author := range *authors {
		controllers.AddAuthor(db, &author)
	}
}
func SeedAuthorBooks(db *gorm.DB) {
	data, err := os.ReadFile("database/initial_author_books.json")

	if err != nil {
		log.Println("error reading initial author_books", err)
	}
	authorBooks := &[]models.AuthorBook{}
	json.Unmarshal(data, authorBooks)
	for _, authorBook := range *authorBooks {
		controllers.AddAuthorBook(db, &authorBook)
	}
}
