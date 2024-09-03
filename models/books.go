package models

type Book struct {
	BookID    uint64 `gorm:"primaryKey;autoIncrement"`
	Title     string
	Pages     uint64
	WordCount uint64
	Status    string
}

type BookWithAuthor struct {
	BookID    uint64
	Title     string
	Pages     uint64
	WordCount uint64
	Status    string
	AuthorID  uint64
}
type BookWithAuthorDetails struct {
	BookID    uint64
	Title     string
	Pages     uint64
	WordCount uint64
	Status    string
	Author    Author
}
