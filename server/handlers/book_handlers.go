package handlers

import (
	"books-api/models"
	"books-api/repos"
	"github.com/labstack/echo/v4"
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
const ErrBadBook = "bad request: can't create book with invalid details"

type BookHandler struct {
	Repo *repos.BookRepository
}

func NewBookHandler(repo *repos.BookRepository) *BookHandler {
	return &BookHandler{repo}
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	status := c.QueryParam("status")
	return c.JSON(http.StatusOK, h.Repo.GetBooks(status))
}

func (h *BookHandler) AddBook(c echo.Context) error {
	b := new(models.BookWithAuthor)

	if err := c.Bind(&b); err != nil {
		return c.String(http.StatusBadRequest, ErrBadBook)
	}

	if !BookStatuses[b.Status] {
		return c.String(http.StatusBadRequest, ErrBadBookStatus)
	}

	err := h.Repo.AddBook(b)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Book)
	c.Bind(&b)

	if !BookStatuses[b.Status] {
		return c.String(http.StatusBadRequest, ErrBadBookStatus)
	}
	h.Repo.UpdateReadingStatus(id, b.Status)
	return c.NoContent(http.StatusAccepted)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	h.Repo.DeleteBook(id)
	return c.NoContent(http.StatusOK)
}

func (h *BookHandler) GetReadingStats(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Repo.GetStats())
}
