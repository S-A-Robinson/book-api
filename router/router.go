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

	//// Add a new author book
	//e.POST("/author-books", func(c echo.Context) error {
	//	ab := new(models.AuthorBook)
	//	err := c.Bind(&ab)
	//
	//	if err != nil {
	//		return c.String(http.StatusBadRequest, "bad request")
	//	}
	//
	//	repos.AddAuthorBook(db, ab)
	//	return c.NoContent(http.StatusCreated)
	//})

	return e
}
