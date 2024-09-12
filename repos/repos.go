package repos

import "gorm.io/gorm"

type Repos struct {
	Author     *AuthorRepository
	Book       *BookRepository
	AuthorBook *AuthorBookRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Author:     NewAuthorRepository(db),
		Book:       NewBookRepository(db),
		AuthorBook: NewAuthorBookRepository(db),
	}
}
