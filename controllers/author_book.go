package controllers

import (
	"books-api/models"
	"gorm.io/gorm"
)

func AddAuthorBook(db *gorm.DB, authorBook *models.AuthorBook) {
	db.Create(authorBook)
}
