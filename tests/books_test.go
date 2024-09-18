package tests

import (
	"books-api/database"
	"books-api/models"
	"books-api/repos"
	"books-api/server"
	"books-api/server/handlers"
	"books-api/tests/helpers"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"
)

var expectedBooks = []models.Book{
	{
		database.InitialBooks[0].ID,
		database.InitialBooks[0].AuthorID,
		database.InitialBooks[0].Title,
		database.InitialBooks[0].Pages,
		database.InitialBooks[0].WordCount,
		database.InitialBooks[0].Status,
		database.InitialAuthors[0],
	},
	{
		database.InitialBooks[1].ID,
		database.InitialBooks[1].AuthorID,
		database.InitialBooks[1].Title,
		database.InitialBooks[1].Pages,
		database.InitialBooks[1].WordCount,
		database.InitialBooks[1].Status,
		database.InitialAuthors[0],
	},
	{
		database.InitialBooks[2].ID,
		database.InitialBooks[2].AuthorID,
		database.InitialBooks[2].Title,
		database.InitialBooks[2].Pages,
		database.InitialBooks[2].WordCount,
		database.InitialBooks[2].Status,
		database.InitialAuthors[1],
	},
}

var expectedFilteredBooks = []models.Book{
	{
		database.InitialBooks[2].ID,
		database.InitialBooks[2].AuthorID,
		database.InitialBooks[2].Title,
		database.InitialBooks[2].Pages,
		database.InitialBooks[2].WordCount,
		database.InitialBooks[2].Status,
		database.InitialAuthors[1],
	},
}

var marshalledBooks, _ = json.Marshal(&expectedBooks)
var marshalledFilteredBooks, _ = json.Marshal(&expectedFilteredBooks)

var invalidNewBook = `{"title": 12, "status": "Reading" }`

func TestGetBooks(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/books",
	}

	cases := []helpers.TestCase{
		{
			TestName:           "it successfully gets all books in db",
			Request:            request,
			RequestContentType: echo.MIMEApplicationJSON,
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       string(marshalledBooks) + "\n",
		},
		{
			TestName: "it successfully filters books by given status",
			Request: helpers.Request{
				Method: http.MethodGet,
				Url:    "/books?status=Reading",
			},
			RequestContentType: echo.MIMEApplicationJSON,
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       string(marshalledFilteredBooks) + "\n",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}
}

func TestPostBooks(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/books",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it adds a new book",
			Request:  request,
			RequestBody: &models.Book{
				ID:        4,
				Title:     "Test New Book",
				Pages:     100,
				WordCount: 2123,
				Status:    "Reading",
				AuthorID:  2,
			},

			ExpectedStatusCode: http.StatusCreated,
		},
		{
			TestName: "it returns a bad request if the status is incorrect",
			Request:  request,
			RequestBody: &models.Book{
				ID:        5,
				Title:     "Invalid Status Book",
				Pages:     101,
				WordCount: 1234,
				Status:    "Invalid Status",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       handlers.ErrBadBookStatus,
		},
		{
			TestName:           "it returns a bad request if the body isn't valid",
			Request:            request,
			RequestBody:        invalidNewBook,
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       handlers.ErrBadBook,
		},
		{
			TestName: "it returns an error if the author doesn't exist",
			Request:  request,
			RequestBody: &models.Book{
				Title:     "Test New Book With Invalid Author",
				Pages:     100,
				WordCount: 2123,
				Status:    "Reading",
				AuthorID:  4000,
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       "couldn't find author with id 4000",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}
}

func TestPutBooks(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/books/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it updates book with new status",
			Request:  request,
			RequestBody: &models.Book{
				Status: "Read",
			},
			ExpectedStatusCode: http.StatusAccepted,
		},
		{
			TestName: "it returns a bad request if the status is invalid",
			Request:  request,
			RequestBody: &models.Book{
				Status: "Invalid Status",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       handlers.ErrBadBookStatus,
		},
		{
			TestName: "it returns a 404 if a book with that id does not exist",
			Request: helpers.Request{
				Method: http.MethodPut,
				Url:    "/books/4000",
			},
			RequestBody: &models.Book{
				Status: "Read",
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedBody:       repos.ErrBookNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}
}
func TestDeleteBooks(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/books/1",
	}

	cases := []helpers.TestCase{
		{
			TestName:           "it successfully deletes a book",
			Request:            request,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			TestName: "it returns a 404 if a book with that id does not exist",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/books/4000",
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedBody:       repos.ErrBookNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}
}
