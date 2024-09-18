package tests

import (
	"books-api/database"
	"books-api/models"
	"books-api/server"
	"books-api/tests/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var marshalledAuthors, _ = json.Marshal(&database.InitialAuthors)

func TestGetAuthors(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/authors",
	}

	buffer := bytes.Buffer{}
	json.NewEncoder(&buffer).Encode("{}")

	cases := []helpers.TestCase{

		{
			TestName:           "it successfully gets all authors in db",
			Request:            request,
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

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/authors",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it adds a new author",
			Request:  request,
			RequestBody: &models.Author{
				ID:        4,
				FirstName: "Test",
				LastName:  "Author",
			},
			ExpectedStatusCode: http.StatusCreated,
			ExpectedBody:       "4",
		},
		{
			TestName:           "it returns a bad request if an invalid author is sent",
			Request:            request,
			RequestBody:        "{}",
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
