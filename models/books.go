package models

type Book struct {
	BookID    uint64 `gorm:"primaryKey;autoIncrement" json:"book_id"`
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`
}

type BookWithAuthor struct {
	BookID    uint64 `json:"book_id"`
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`
	AuthorID  uint64 `json:"author_id"`
}
type BookWithAuthorDetails struct {
	BookID    uint64 `json:"book_id"`
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`
	Author    Author `json:"author"`
}
