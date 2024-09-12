package models

type AuthorBook struct {
	AuthorID uint64 `gorm:"primary_key" json:"author_id"`
	BookID   uint64 `gorm:"primary_key" json:"book_id"`
}
