package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Setenv("ENVIRONMENT", "testing")

	if err != nil {
		fmt.Printf("Error setting up testing environment: %v", err)
	}

	code := m.Run()

	os.Exit(code)
}
