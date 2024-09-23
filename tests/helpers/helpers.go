package helpers

import (
	"io"
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
	StatusCode       int
	BodyPart         string
	BodyParts        []string
	BodyPartMissing  string
	BodyPartsMissing []string
}
