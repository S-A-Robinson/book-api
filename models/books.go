package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string
	Author    string
	Pages     int
	WordCount int
	Status    string
}
