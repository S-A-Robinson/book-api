package tests

import (
	"books-api/tests/helpers"
	"net/http"
	"testing"
)

func TestStats(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/stats",
	}

	cases := []helpers.TestCase{
		{
			TestName: "it returns stats",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					"2181",
					"540500",
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
