package database

import "books-api/models"

var InitialAuthors = []models.Author{
	{
		ID:        1,
		FirstName: "Dan",
		LastName:  "Simmons",
	},
	{
		ID:        2,
		FirstName: "Stephen",
		LastName:  "King",
	},
}
