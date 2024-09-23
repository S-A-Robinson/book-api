package handlers

import (
	"books-api/models"
	"books-api/repos"
	"books-api/server/requests"
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
	title := c.QueryParam("title")
	return c.JSON(http.StatusOK, h.Repo.GetBooks(status, title))
}

func (h *BookHandler) AddBook(c echo.Context) error {
	newBookRequest := new(requests.BookRequest)

	err := c.Bind(newBookRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, ErrBadBook)
	}

	if !BookStatuses[newBookRequest.Status] {
		return c.String(http.StatusBadRequest, ErrBadBookStatus)
	}

	book := &models.Book{
		Title:     newBookRequest.Title,
		Pages:     newBookRequest.Pages,
		WordCount: newBookRequest.WordCount,
		Status:    newBookRequest.Status,
		AuthorID:  newBookRequest.AuthorID,
	}

	addBookErr := h.Repo.AddBook(book)

	if addBookErr != nil {
		return c.String(http.StatusBadRequest, addBookErr.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	newBookStatusRequest := new(requests.BookStatusRequest)

	if err := c.Bind(newBookStatusRequest); err != nil {
		return err
	}

	if !BookStatuses[newBookStatusRequest.Status] {
		return c.String(http.StatusBadRequest, ErrBadBookStatus)
	}

	err := h.Repo.UpdateReadingStatus(id, newBookStatusRequest.Status)

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.Repo.DeleteBook(id)

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *BookHandler) GetReadingStats(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Repo.GetStats())
}
