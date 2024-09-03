package tests

import (
	"books-api/database"
	"books-api/models"
	"books-api/server"
	"books-api/server/handlers"
	"books-api/tests/helpers"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectedBooks = []models.BookWithAuthorDetails{
	{
		1,
		database.InitialBooks[0].Title,
		database.InitialBooks[0].Pages,
		database.InitialBooks[0].WordCount,
		database.InitialBooks[0].Status,
		database.InitialAuthors[0],
	},
	{
		2,
		database.InitialBooks[1].Title,
		database.InitialBooks[1].Pages,
		database.InitialBooks[1].WordCount,
		database.InitialBooks[1].Status,
		database.InitialAuthors[0],
	},
	{
		3,
		database.InitialBooks[2].Title,
		database.InitialBooks[2].Pages,
		database.InitialBooks[2].WordCount,
		database.InitialBooks[2].Status,
		database.InitialAuthors[1],
	},
}

var marshalledBooks, _ = json.Marshal(&expectedBooks)

func TestGetBooks(t *testing.T) {
	s := server.New()
	cases := []helpers.TestCase{
		{
			TestName:           "it successfully gets all books in db",
			Request:            httptest.NewRequest(http.MethodGet, "/books", nil),
			RequestContentType: echo.MIMEApplicationJSON,
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       string(marshalledBooks) + "\n",
		},
		{
			TestName: "it returns a bad request if the status is incorrect",
			Request: httptest.NewRequest(http.MethodPost, "/books", helpers.Encode(&models.Book{
				BookID:    2,
				Title:     "Invalid Status Book",
				Pages:     101,
				WordCount: 1234,
				Status:    "Invalid Status",
			})),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       handlers.ErrBadBookStatus,
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
	cases := []helpers.TestCase{
		{
			TestName: "it adds a new book",
			Request: httptest.NewRequest(http.MethodPost, "/books", helpers.Encode(&models.BookWithAuthor{
				Title:     "Test New Book",
				Pages:     100,
				WordCount: 2123,
				Status:    "Reading",
				AuthorID:  2,
			})),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusCreated,
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
	cases := []helpers.TestCase{
		{
			TestName: "it updates book with new status",
			Request: httptest.NewRequest(http.MethodPut, "/books/1", helpers.Encode(&models.Book{
				Status: "Read",
			})),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusAccepted,
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
	cases := []helpers.TestCase{
		{
			TestName:           "it successfully deletes a book",
			Request:            httptest.NewRequest(http.MethodDelete, "/books/1", nil),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}
}
