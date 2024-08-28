package main

import (
	"books-api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
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

func createDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	db.AutoMigrate(&models.Book{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func main() {
	db := createDBConnection()

	e := echo.New()
	// Get all books
	e.GET("/books", func(c echo.Context) error {
		status := c.QueryParam("status")
		return c.JSON(http.StatusOK, getBooks(db, status))
	})
	// Get book stats
	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, getStats(db))
	})
	// Add new book
	e.POST("/books", func(c echo.Context) error {
		b := new(models.Book)
		err := c.Bind(&b)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		addBook(db, b)
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
		updateReadingStatus(db, id, b.Status)
		return c.NoContent(http.StatusAccepted)
	})
	// Delete book from list
	e.DELETE("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		deleteBook(db, id)
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getBooks(db *gorm.DB, status string) []models.Book {
	books := make([]models.Book, 0)
	db.Find(&books, models.Book{Status: status})
	return books
}

func addBook(db *gorm.DB, book *models.Book) {
	db.Create(book)
}

func updateReadingStatus(db *gorm.DB, id, newStatus string) {
	db.Model(&models.Book{}).Where("id = ?", id).Update("Status", newStatus)
}

func deleteBook(db *gorm.DB, id string) {
	db.Where("id = ?", id).Delete(&models.Book{})
}

type stats struct {
	Pages     int
	WordCount int
}

func getStats(db *gorm.DB) *stats {
	stats := &stats{0, 0}

	db.Debug().Table("books").Select("sum(pages)").Row().Scan(&stats.Pages)
	db.Table("books").Select("sum(word_count)").Row().Scan(&stats.WordCount)
	return stats
}
