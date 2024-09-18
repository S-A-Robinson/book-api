package repos

import (
	"books-api/models"
	"fmt"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

type FullBookDataStruct struct {
	models.Book
	Author models.Author `gorm:"embedded" json:"author"`
}

var ErrBookNotFound = "couldn't find book with that id"

func (r *BookRepository) GetBooks(status string) []*models.Book {
	var books []*models.Book
	r.DB.
		Preload("Author").
		Find(&books, &models.Book{Status: status})

	return books
}

func (r *BookRepository) AddBook(book *models.Book) error {
	a := &models.Author{}
	r.DB.Where("id = ?", book.AuthorID).First(&a)

	if a.ID != book.AuthorID {
		return fmt.Errorf("couldn't find author with id %d", book.AuthorID)
	}

	r.DB.Create(&book)

	return nil
}

func (r *BookRepository) UpdateReadingStatus(id, newStatus string) error {
	tx := r.DB.
		Model(&models.Book{}).
		Where("id = ?", id).
		Update("Status", newStatus)

	if tx.RowsAffected == 0 {
		return fmt.Errorf(ErrBookNotFound)
	}

	return nil
}

func (r *BookRepository) DeleteBook(id string) error {
	tx := r.DB.Where("id = ?", id).Delete(&models.Book{})

	if tx.RowsAffected == 0 {
		return fmt.Errorf(ErrBookNotFound)
	}

	return nil
}

type stats struct {
	Pages     int
	WordCount int
}

func (r *BookRepository) GetStats() *stats {
	stats := &stats{0, 0}

	r.DB.
		Table("books").
		Select("sum(pages)").
		Row().
		Scan(&stats.Pages)
	r.DB.
		Table("books").
		Select("sum(word_count)").
		Row().
		Scan(&stats.WordCount)
	return stats
}
