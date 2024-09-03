package server

import (
	"books-api/database"
	"books-api/repos"
	"books-api/router"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
)

type Server struct {
	Echo     *echo.Echo
	Database *gorm.DB
	Repos    *repos.Repos
}

var devEnvFileLocation = ".env"
var testEnvFileLocation = ".env.test"

func New() *Server {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "testing" {
		fmt.Println("Running in testing mode")
		godotenv.Load(testEnvFileLocation)
	} else {
		fmt.Println("Running in dev mode")
		godotenv.Load(devEnvFileLocation)
	}

	db := database.New(os.Getenv("DB_LOCATION"))
	r := repos.NewRepos(db)
	e := router.New(r)
	return &Server{e, db, r}
}
