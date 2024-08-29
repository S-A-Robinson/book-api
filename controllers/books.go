package controllers

import (
	"books-api/models"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB, status string) []models.Book {
	books := make([]models.Book, 0)
	db.Find(&books, models.Book{Status: status})
	return books
}

func AddBook(db *gorm.DB, book *models.Book) {
	db.Create(book)
}

func UpdateReadingStatus(db *gorm.DB, id, newStatus string) {
	db.Model(&models.Book{}).Where("id = ?", id).Update("Status", newStatus)
}

func DeleteBook(db *gorm.DB, id string) {
	db.Where("id = ?", id).Delete(&models.Book{})
}

type stats struct {
	Pages     int
	WordCount int
}

func GetStats(db *gorm.DB) *stats {
	stats := &stats{0, 0}

	db.Debug().Table("books").Select("sum(pages)").Row().Scan(&stats.Pages)
	db.Table("books").Select("sum(word_count)").Row().Scan(&stats.WordCount)
	return stats
}
