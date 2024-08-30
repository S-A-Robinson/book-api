package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	TestName           string
	Request            *http.Request
	RequestContentType string
	RequestReader      *httptest.ResponseRecorder
	ExpectedStatusCode int
	ExpectedBody       string
}

func ExecuteTest(t *testing.T, e *echo.Echo, testCase TestCase) {
	testCase.Request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e.ServeHTTP(testCase.RequestReader, testCase.Request)

	assertStatusCode(t, testCase.RequestReader.Code, testCase.ExpectedStatusCode)
	assertResponseBody(t, testCase.RequestReader.Body.String(), testCase.ExpectedBody)
}

func Encode(data interface{}) *bytes.Buffer {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(data)

	return &buffer
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted status code %v but got %v", want, got)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted response body of %q got %q", want, got)
	}
}
