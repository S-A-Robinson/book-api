package tests

import (
	"books-api/server"
	"books-api/tests/helpers"
	"net/http"
	"testing"
)

func TestStats(t *testing.T) {
	s := server.New()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/stats",
	}

	cases := []helpers.TestCase{
		{
			TestName:           "it returns stats",
			Request:            request,
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       "{\"Pages\":2181,\"WordCount\":540500}\n",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}

}
