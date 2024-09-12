package models

type Author struct {
	AuthorID  uint64 `gorm:"primary_key;auto_increment" json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
