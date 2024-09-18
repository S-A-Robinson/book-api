package repos

import "gorm.io/gorm"

type Repos struct {
	Author *AuthorRepository
	Book   *BookRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Author: NewAuthorRepository(db),
		Book:   NewBookRepository(db),
	}
}
