package tests

import (
	"books-api/database"
	"books-api/models"
	"books-api/server"
	"books-api/tests/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var marshalledAuthors, _ = json.Marshal(&database.InitialAuthors)

func TestGetAuthors(t *testing.T) {
	s := server.New()

	buffer := bytes.Buffer{}

	json.NewEncoder(&buffer).Encode("{}")

	cases := []helpers.TestCase{

		{
			TestName:           "it successfully gets all authors in db",
			Request:            httptest.NewRequest(http.MethodGet, "/authors", nil),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       string(marshalledAuthors) + "\n",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}

}

func TestPostAuthors(t *testing.T) {
	s := server.New()

	buffer := bytes.Buffer{}

	json.NewEncoder(&buffer).Encode("{}")

	cases := []helpers.TestCase{
		{
			TestName: "it adds a new author",
			Request: httptest.NewRequest(http.MethodPost, "/authors", helpers.Encode(&models.Author{
				AuthorID:  4,
				FirstName: "Test",
				LastName:  "Author",
			})),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusCreated,
		},
		{
			TestName:           "it returns a bad request if an invalid author is sent",
			Request:            httptest.NewRequest(http.MethodPost, "/authors", &buffer),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       "bad request",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}

}
