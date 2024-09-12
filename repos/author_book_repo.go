package repos

import (
	"books-api/models"
	"fmt"
	"gorm.io/gorm"
)

type AuthorBookRepository struct {
	DB *gorm.DB
}

func NewAuthorBookRepository(db *gorm.DB) *AuthorBookRepository {
	return &AuthorBookRepository{db}
}

func (r *BookRepository) DeleteAuthorBookByBookID(id string) error {
	tx := r.DB.Where("book_id = ?", id).Delete(&models.AuthorBook{})

	if tx.RowsAffected == 0 {
		return fmt.Errorf(ErrBookNotFound)
	}

	return nil
}
