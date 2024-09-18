package database

import "books-api/models"

var InitialAuthors = []models.Author{
	{
		ID:        1,
		FirstName: "Dan",
		LastName:  "Simmons",
		ImageURL:  "https://static.wikia.nocookie.net/absolutehorror/images/9/97/Dan_Simmons.jpg/",
	},
	{
		ID:        2,
		FirstName: "Stephen",
		LastName:  "King",
		ImageURL:  "https://static01.nyt.com/images/2015/10/31/arts/31KING/31KING-superJumbo.jpg?quality=75&auto=webp",
	},
}
