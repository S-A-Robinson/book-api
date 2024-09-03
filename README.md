# Go Api
This is a simple project to test out creating an api using Go, Echo and Gorm
and sqlite

# Dev Setup
* Copy content of `.env.dist` to a `.env` file and populate fields

# Libraries
* web framework - echo: "github.com/labstack/echo/v4"
* orm - gorm: "gorm.io/gorm"
  * gorm sqlite driver: "gorm.io/driver/sqlite"

# Testing
* You can run the test suite with `go test` from the `/tests` folder

# Run
* You can run the project with `go run main.go`

# Build
* Build the project with `go build`