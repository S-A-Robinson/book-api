package tests

import (
	"books-api/server"
	"books-api/tests/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStats(t *testing.T) {
	s := server.New()
	cases := []helpers.TestCase{
		{
			TestName:           "it returns stats",
			Request:            httptest.NewRequest(http.MethodGet, "/stats", nil),
			RequestReader:      httptest.NewRecorder(),
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       "{\"Pages\":0,\"WordCount\":0}\n",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			helpers.ExecuteTest(t, s.Echo, testCase)
		})
	}

}
