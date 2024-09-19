package database

import (
	"books-api/models"
	"books-api/repos"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{}, &models.Author{})

	if err != nil {
		log.Println("failed to migrate the database: ", err)
	}
}

func Seed(db *gorm.DB) {
	SeedAuthors(db)
	SeedBooks(db)
}

func SeedAuthors(db *gorm.DB) {
	authors := &InitialAuthors
	authorRepo := repos.NewAuthorRepository(db)
	for _, author := range *authors {
		authorRepo.DB.Where("id = ?", author.ID).FirstOrCreate(&author)
	}
}

func SeedBooks(db *gorm.DB) {
	books := &InitialBooks
	bookRepo := repos.NewBookRepository(db)
	for _, book := range *books {
		bookRepo.DB.Where("id = ?", book.ID).FirstOrCreate(&book)
	}
}
