package unithelpers_test

import (
	"fmt"
	"os"
	"testing"

	uh "github.com/allentv/go-testhelpers/unithelpers"
)

func ExampleTestHelper_SetOSEnvVar() {
	th := uh.New(&testing.T{})

	// Defer the closer function after invocation to clean up at the end of the test
	defer th.SetOSEnvVar("SOME_KEY", "TEST_VALUE")()

	// Test statements here...
	fmt.Println(os.Getenv("SOME_KEY"))
	// Test statements here...

	// Output:
	// TEST_VALUE
}
