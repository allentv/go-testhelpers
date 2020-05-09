package unithelpers

import (
	"testing"
)

// TestDataFolder is the folder where test data is stored for unit tests
var TestDataFolder = "testdata"

// TestHelper wraps the testing object
type TestHelper struct {
	t              *testing.T
	TestDataFolder string
}

// Helperer contains the functions for testing.T
type Helperer interface {
	helper()
	fatal(args ...interface{})
	fatalf(format string, args ...interface{})
}

// Helper is a convenience wrapper
func (h *TestHelper) helper() {
	h.t.Helper()
}

// Fatal is a convenience wrapper
func (h *TestHelper) fatal(args ...interface{}) {
	h.t.Fatal(args...)
}

// Fatalf is a convenience wrapper
func (h *TestHelper) fatalf(format string, args ...interface{}) {
	h.t.Fatalf(format, args...)
}

// New is the factory function to generate the struct
func New(t *testing.T) *TestHelper {
	return &TestHelper{
		t:              t,
		TestDataFolder: TestDataFolder,
	}
}

// NewWithFolder is the factory function to generate the struct with a specified test data folder name
func NewWithFolder(t *testing.T, testDataFolder string) *TestHelper {
	return &TestHelper{
		t:              t,
		TestDataFolder: testDataFolder,
	}
}

// GetTestDataFolder retuns the configured test data folder
func GetTestDataFolder() string {
	return TestDataFolder
}
