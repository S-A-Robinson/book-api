package models

type Book struct {
	BookID    uint64 `gorm:"primaryKey;autoIncrement"`
	Title     string
	Pages     uint64
	WordCount uint64
	Status    string
}
