package controllers

import (
	"books-api/models"
	"gorm.io/gorm"
)

func AddAuthor(db *gorm.DB, author *models.Author) {
	db.Create(author)
}

func GetAuthors(db *gorm.DB) *[]models.Author {
	authors := new([]models.Author)
	db.Find(authors)
	return authors
}
