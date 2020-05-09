package unithelpers

import (
	"testing"
)

// TestHelper wraps the testing object
type TestHelper struct {
	*testing.T
}

// New is the factory function to generate the struct
func New(t *testing.T) *TestHelper {
	return &TestHelper{t}
}
