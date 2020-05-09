package unithelpers_test

import (
	"fmt"
	"testing"

	uh "github.com/allentv/go-testhelpers/unithelpers"
)

func ExampleTestHelper_LoadData() {
	th := uh.New(&testing.T{})

	result := th.LoadData("sample1.txt")

	fmt.Println(result)
	// Output: Some Test data for Sample1\n
}
