package main

import (
	"books-api/server"
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("ENVIRONMENT", "development")

	if err != nil {
		fmt.Printf("Error setting up testing environment: %v", err)
	}

	s := server.New()

	s.Echo.Logger.Fatal(s.Echo.Start(":8080"))
}
