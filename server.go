package main

import (
	"books-api/database"
	"books-api/router"
)

func main() {
	db := database.New()
	e := router.New(db)

	e.Logger.Fatal(e.Start(":8080"))
}
