package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	TestName           string
	Request            Request
	RequestContentType string
	RequestBody        interface{}
	RequestReader      io.Reader
	ExpectedStatusCode int
	ExpectedBody       string
}

type PathParam struct {
	Name  string
	Value string
}

type Request struct {
	Method    string
	Url       string
	PathParam *PathParam
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
	BodyParts  []string
}

func ExecuteTest(t *testing.T, e *echo.Echo, testCase TestCase) {
	req, _ := GenerateRequest(&testCase)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	assertStatusCode(t, rr.Code, testCase.ExpectedStatusCode)
	assertResponseBody(t, rr.Body.String(), testCase.ExpectedBody)
}

func GenerateRequest(testCase *TestCase) (*http.Request, error) {
	reqJson, err := json.Marshal(testCase.RequestBody)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if testCase.RequestReader != nil {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, testCase.RequestReader)
	} else {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, bytes.NewBuffer(reqJson))
	}

	if err != nil {
		return nil, err
	}

	if testCase.RequestContentType != "" {
		req.Header.Set(echo.HeaderContentType, testCase.RequestContentType)
	} else {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	return req, nil
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
