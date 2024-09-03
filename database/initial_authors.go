package database

import "books-api/models"

var InitialAuthors = []models.Author{
	{
		AuthorID:  1,
		FirstName: "Dan",
		LastName:  "Simmons",
	},
	{
		AuthorID:  2,
		FirstName: "Stephen",
		LastName:  "King",
	},
}
