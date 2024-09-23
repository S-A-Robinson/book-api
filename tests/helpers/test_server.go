package helpers

import (
	"books-api/database"
	"books-api/server"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	S *server.Server
}

func NewTestServer() *TestServer {
	return &TestServer{
		S: server.New(),
	}
}

func (ts *TestServer) ExecuteTestCase(testCase *TestCase) *httptest.ResponseRecorder {
	ts.SetupDefaults()

	req, _ := ts.GenerateRequest(testCase)
	res := ts.ExecuteRequest(req)
	return res
}

func (ts *TestServer) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ts.S.Echo.ServeHTTP(rr, req)
	return rr
}

func (ts *TestServer) GenerateRequest(testCase *TestCase) (*http.Request, error) {
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

func (ts *TestServer) ClearTable(tableName string) {
	err := ts.S.Database.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}
}

func (ts *TestServer) SetupDefaults() {
	ts.ClearTable("books")
	ts.ClearTable("authors")
	database.Seed(ts.S.Database)
}
