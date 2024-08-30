package controllers

import (
	"books-api/models"
	"gorm.io/gorm"
)

type fullBookDataStruct struct {
	models.Book
	Author models.Author `gorm:"embedded"`
}

func GetBooks(db *gorm.DB, status string) []fullBookDataStruct {
	books := make([]fullBookDataStruct, 0)
	db.
		Model(&models.Book{}).
		Select("books.*, authors.*").
		Joins("left join author_books on (books.book_id = author_books.book_id)").
		Joins("left join authors on (authors.author_id = author_books.author_id)").
		Find(&books, models.Book{Status: status})
	return books
}

func AddBook(db *gorm.DB, book *models.Book) {
	db.Create(book)
}

func UpdateReadingStatus(db *gorm.DB, id, newStatus string) {
	db.
		Model(&models.Book{}).
		Where("book_id = ?", id).
		Update("Status", newStatus)
}

func DeleteBook(db *gorm.DB, id string) {
	db.Where("book_id = ?", id).Delete(&models.Book{})
}

type stats struct {
	Pages     int
	WordCount int
}

func GetStats(db *gorm.DB) *stats {
	stats := &stats{0, 0}

	db.
		Table("books").
		Select("sum(pages)").
		Row().
		Scan(&stats.Pages)
	db.
		Table("books").
		Select("sum(word_count)").
		Row().
		Scan(&stats.WordCount)
	return stats
}
