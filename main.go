package main

import "books-api/server"

func main() {
	s := server.New()

	s.Echo.Logger.Fatal(s.Echo.Start(":8080"))
}
