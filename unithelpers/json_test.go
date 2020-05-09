package unithelpers_test

import (
	"fmt"
	"testing"

	uh "github.com/allentv/go-testhelpers/unithelpers"
)

func ExampleTestHelper_GetJSON() {
	type Test struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	th := uh.New(&testing.T{})

	result := th.GetJSON(Test{Field1: "f1", Field2: 2})
	fmt.Println(result)
	// Output: {"field1":"f1","field2":2}
}
