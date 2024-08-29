package controllers

import (
	"books-api/models"
	"gorm.io/gorm"
)

func AddAuthor(db *gorm.DB, author *models.Author) {
	db.Create(author)
}
