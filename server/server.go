package server

import (
	"books-api/database"
	"books-api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
)

type Server struct {
	Echo     *echo.Echo
	Database *gorm.DB
}

func New() *Server {
	godotenv.Load("../.env.test")
	db := database.New(os.Getenv("DB_LOCATION"))
	e := router.New(db)
	return &Server{e, db}
}
