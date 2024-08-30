package tests

import (
	"books-api/models"
	"books-api/server"
	"books-api/tests/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorBooks(t *testing.T) {
	s := server.New()
	cases := []helpers.TestCase{
		{
			TestName: "it adds a new author-book",
			Request: httptest.NewRequest(http.MethodPost, "/author-books", helpers.Encode(&models.AuthorBook{
				AuthorID: 1,
				BookID:   1,
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
