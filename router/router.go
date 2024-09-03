package router

import (
	"books-api/repos"
	"books-api/server/handlers"
	"github.com/labstack/echo/v4"
)

func New(r *repos.Repos) *echo.Echo {
	e := echo.New()

	bookHandler := handlers.NewBookHandler(r.Book)
	e.GET("/books", bookHandler.GetBooks)
	e.GET("/stats", bookHandler.GetReadingStats)
	e.POST("/books", bookHandler.AddBook)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	authorHandler := handlers.NewAuthorHandler(r.Author)
	e.GET("/authors", authorHandler.GetAuthors)
	e.POST("/authors", authorHandler.AddAuthor)

	return e
}
