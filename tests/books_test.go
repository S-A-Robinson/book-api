package tests

import (
	"books-api/models"
	"books-api/router"
	"books-api/server"
	"books-api/tests/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBooks(t *testing.T) {
	s := server.New()
	cases := []helpers.TestCase{
		{
			TestName:           "it successfully gets all books in db",
			Request:            httptest.NewRequest(http.MethodGet, "/books", nil),
			RequestContentType: echo.MIMEApplicationJSON,
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       "[]\n",
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
			ExpectedBody:       router.ErrBadBookStatus,
		},
		{
			TestName: "it adds a new book",
			Request: httptest.NewRequest(http.MethodPost, "/books", helpers.Encode(&models.Book{
				BookID:    1,
				Title:     "Valid Test Book",
				Pages:     100,
				WordCount: 2123,
				Status:    "Plan To Read",
			})),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusCreated,
		},
		{
			TestName:           "it successfully deletes a book",
			Request:            httptest.NewRequest(http.MethodDelete, "/books/1", nil),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
		},
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
