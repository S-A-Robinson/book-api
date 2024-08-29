package router

import (
	"books-api/controllers"
	"books-api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

const (
	BookStatusReading  string = "Reading"
	BookStatusRead     string = "Read"
	BookStatusToBeRead string = "To Be Read"
)

var BookStatuses = map[string]bool{
	BookStatusReading:  true,
	BookStatusRead:     true,
	BookStatusToBeRead: true,
}

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	// Get all books
	e.GET("/books", func(c echo.Context) error {
		status := c.QueryParam("status")
		return c.JSON(http.StatusOK, controllers.GetBooks(db, status))
	})
	// Get book stats
	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, controllers.GetStats(db))
	})
	// Add new book
	e.POST("/books", func(c echo.Context) error {
		b := new(models.Book)
		err := c.Bind(&b)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		controllers.AddBook(db, b)
		return c.NoContent(http.StatusOK)
	})
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
	// Delete book from list
	e.DELETE("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		controllers.DeleteBook(db, id)
		return c.NoContent(http.StatusOK)
	})

	return e
}
