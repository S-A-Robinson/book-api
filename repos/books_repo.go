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
	Author models.Author `gorm:"embedded"`
}

var ErrBookNotFound = "couldn't find book with that id"

func (r *BookRepository) GetBooks(status string) []FullBookDataStruct {
	books := make([]FullBookDataStruct, 0, 10)
	r.DB.
		Model(&models.Book{}).
		Select("books.*, authors.*").
		Joins("left join author_books on (books.book_id = author_books.book_id)").
		Joins("left join authors on (authors.author_id = author_books.author_id)").
		Find(&books, &models.Book{Status: status})
	return books
}

func (r *BookRepository) AddBook(bookWithAuthor *models.BookWithAuthor) error {
	a := &models.Author{}
	r.DB.Where("author_id = ?", bookWithAuthor.AuthorID).First(&a)

	if a.AuthorID != bookWithAuthor.AuthorID {
		return fmt.Errorf("couldn't find author with id %d", bookWithAuthor.AuthorID)
	}

	b := models.Book{
		BookID:    bookWithAuthor.BookID,
		Title:     bookWithAuthor.Title,
		Pages:     bookWithAuthor.Pages,
		WordCount: bookWithAuthor.WordCount,
		Status:    bookWithAuthor.Status,
	}
	r.DB.Create(&b)

	r.DB.Create(&models.AuthorBook{
		AuthorID: bookWithAuthor.AuthorID,
		BookID:   b.BookID,
	})
	return nil
}

func (r *BookRepository) UpdateReadingStatus(id, newStatus string) error {
	tx := r.DB.
		Model(&models.Book{}).
		Where("book_id = ?", id).
		Update("Status", newStatus)

	if tx.RowsAffected == 0 {
		return fmt.Errorf(ErrBookNotFound)
	}

	return nil
}

func (r *BookRepository) DeleteBook(id string) error {
	tx := r.DB.Where("book_id = ?", id).Delete(&models.Book{})

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
