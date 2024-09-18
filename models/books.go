package models

type Book struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorID  uint64 `json:"author_id"`
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`

	Author Author `json:"author"`
}

type BookWithAuthorDetails struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`
	Author    Author `json:"author"`
}
