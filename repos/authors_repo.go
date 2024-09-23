package repos

import (
	"books-api/models"
	"errors"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

var ErrAuthorNotFound = "couldn't find author with that id"

func (r *AuthorRepository) GetAuthors() *[]models.Author {
	authors := new([]models.Author)
	r.DB.Find(authors)
	return authors
}

func (r *AuthorRepository) AddAuthor(author *models.Author) uint64 {
	r.DB.Create(author)
	return author.ID
}

func (r *AuthorRepository) DeleteAuthor(id string) error {
	tx := r.DB.Where("id = ? ", id).Delete(&models.Author{})
	if tx.RowsAffected == 0 {
		return errors.New("no authors found")
	}

	r.DB.Where("author_id = ?", id).Delete(&models.Book{})

	return nil
}
