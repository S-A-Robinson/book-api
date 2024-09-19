package helpers

import (
	"io"
	"testing"
)

type TestCase struct {
	TestName           string
	Request            Request
	RequestContentType string
	RequestBody        interface{}
	RequestReader      io.Reader
	Expected           ExpectedResponse
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
