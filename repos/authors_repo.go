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

func (r *AuthorRepository) AddAuthor(author *models.Author) {
	r.DB.Create(author)
}

func (r *AuthorRepository) GetAuthors() *[]models.Author {
	authors := new([]models.Author)
	r.DB.Find(authors)
	return authors
}
