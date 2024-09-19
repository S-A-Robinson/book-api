package tests

import (
	"books-api/database"
	"books-api/models"
	"books-api/tests/helpers"
	"net/http"
	"testing"
)

func TestGetAuthors(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/authors",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it successfully gets all authors in db",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					database.InitialAuthors[0].FirstName,
					database.InitialAuthors[0].LastName,
					database.InitialAuthors[1].FirstName,
					database.InitialAuthors[1].LastName,
				},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			RunTestCase(t, testCase)
		})
	}

}

func TestPostAuthors(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/authors",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it adds a new author",
			Request:  request,
			RequestBody: &models.Author{
				FirstName: "Test",
				LastName:  "Author",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyPart:   "3",
			},
		},
		{
			TestName:    "it returns a bad request if an invalid author is sent",
			Request:     request,
			RequestBody: "{}",
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "bad request",
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			RunTestCase(t, testCase)
		})
	}

}
