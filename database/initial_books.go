package database

import "books-api/models"

var InitialBooks = []models.Book{
	{
		ID:        1,
		Title:     "Hyperion",
		Pages:     500,
		WordCount: 120250,
		Status:    "Read",
		AuthorID:  1,
	},
	{
		ID:        2,
		Title:     "The Fall of Hyperion",
		Pages:     528,
		WordCount: 132000,
		Status:    "Read",
		AuthorID:  1,
	},
	{
		ID:        3,
		Title:     "The Stand",
		Pages:     1153,
		WordCount: 288250,
		Status:    "Reading",
		AuthorID:  2,
	},
}
