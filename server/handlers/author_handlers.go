package handlers

import (
	"books-api/models"
	"books-api/repos"
	"books-api/server/requests"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	Repo *repos.AuthorRepository
}

func NewAuthorHandler(repo *repos.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{repo}
}

func (h *AuthorHandler) GetAuthors(c echo.Context) error {
	authors := h.Repo.GetAuthors()
	return c.JSON(http.StatusOK, authors)
}

func (h *AuthorHandler) AddAuthor(c echo.Context) error {
	newAuthorRequest := new(requests.Author)
	err := c.Bind(&newAuthorRequest)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	author := &models.Author{
		FirstName: newAuthorRequest.FirstName,
		LastName:  newAuthorRequest.LastName,
		ImageURL:  newAuthorRequest.ImageURL,
	}

	id := h.Repo.AddAuthor(author)
	return c.String(http.StatusCreated, strconv.FormatUint(id, 10))
}

func (h *AuthorHandler) DeleteAuthor(c echo.Context) error {
	id := c.Param("id")

	err := h.Repo.DeleteAuthor(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
