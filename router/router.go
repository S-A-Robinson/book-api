package router

import (
	"books-api/controllers"
	"books-api/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

const (
	BookStatusReading    string = "Reading"
	BookStatusRead       string = "Read"
	BookStatusPlanToRead string = "Plan To Read"
)

var BookStatuses = map[string]bool{
	BookStatusReading:    true,
	BookStatusRead:       true,
	BookStatusPlanToRead: true,
}

const ErrBadBookStatus = "bad request: a book status must be either: 'Reading', 'Read' or 'Plan To Read'"

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	// GET

	// Get all books
	e.GET("/books", func(c echo.Context) error {
		status := c.QueryParam("status")
		return c.JSON(http.StatusOK, controllers.GetBooks(db, status))
	})

	// Get book stats
	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, controllers.GetStats(db))
	})

	// Get all authors
	e.GET("/authors", func(c echo.Context) error {
		return c.JSON(http.StatusOK, controllers.GetAuthors(db))
	})

	// POST

	// Add new book
	e.POST("/books", func(c echo.Context) error {
		b := new(models.Book)

		if err := c.Bind(&b); err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("bad request: %s", err))
		}

		if !BookStatuses[b.Status] {
			return c.String(http.StatusBadRequest, ErrBadBookStatus)
		}

		controllers.AddBook(db, b)
		return c.NoContent(http.StatusCreated)
	})

	// Add a new author
	e.POST("/authors", func(c echo.Context) error {
		a := new(models.Author)
		err := c.Bind(&a)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		controllers.AddAuthor(db, a)
		return c.NoContent(http.StatusCreated)
	})

	// Add a new author book
	e.POST("/author-books", func(c echo.Context) error {
		ab := new(models.AuthorBook)
		err := c.Bind(&ab)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		controllers.AddAuthorBook(db, ab)
		return c.NoContent(http.StatusCreated)
	})

	// PUT

	// Update reading status of book
	e.PUT("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		b := new(models.Book)
		c.Bind(&b)

		if !BookStatuses[b.Status] {
			return c.String(http.StatusBadRequest, "bad status")
		}
		controllers.UpdateReadingStatus(db, id, b.Status)
		return c.NoContent(http.StatusAccepted)
	})

	//DELETE

	// Delete book from list
	e.DELETE("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		controllers.DeleteBook(db, id)
		return c.NoContent(http.StatusOK)
	})

	return e
}
