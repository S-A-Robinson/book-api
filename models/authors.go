package models

type Author struct {
	AuthorID  uint64 `gorm:"primary_key;auto_increment"`
	FirstName string
	LastName  string
}
