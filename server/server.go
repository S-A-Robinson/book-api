package server

import (
	"books-api/database"
	"books-api/repos"
	"books-api/router"
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

func New() *Server {
	godotenv.Load(".env")
	db := database.New(os.Getenv("DB_LOCATION"))
	r := repos.NewRepos(db)
	e := router.New(r)
	return &Server{e, db, r}
}
