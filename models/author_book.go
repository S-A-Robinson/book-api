package models

type AuthorBook struct {
	AuthorID uint64 `gorm:"primary_key"`
	BookID   uint64 `gorm:"primary_key"`
}
