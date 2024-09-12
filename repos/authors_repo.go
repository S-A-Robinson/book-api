package repos

import (
	"books-api/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

func (r *AuthorRepository) AddAuthor(author *models.Author) uint64 {
	r.DB.Create(author)
	return author.AuthorID
}

func (r *AuthorRepository) GetAuthors() *[]models.Author {
	authors := new([]models.Author)
	r.DB.Find(authors)
	return authors
}
