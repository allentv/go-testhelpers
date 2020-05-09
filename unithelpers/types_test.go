package unithelpers_test

import (
	"fmt"
	"testing"

	uh "github.com/allentv/go-testhelpers/unithelpers"
)

func ExampleNew() {
	th := uh.New(&testing.T{})

	fmt.Printf("%T\n", th)
	fmt.Println(uh.TestDataFolder)

	// Output:
	// *unithelpers.TestHelper
	// testdata
}

func ExampleNewWithFolder() {
	th := uh.NewWithFolder(&testing.T{}, "something")
	defer func() {
		uh.TestDataFolder = "testdata"
	}()

	fmt.Printf("%T\n", th)
	fmt.Println(th.TestDataFolder)

	// Output:
	// *unithelpers.TestHelper
	// something
}

func ExampleGetTestDataFolder() {
	fmt.Println(uh.GetTestDataFolder())

	// Output: testdata
}

func ExampleGetTestDataFolder_ManualUpdate() {
	uh.TestDataFolder = "something-else"
	defer func() {
		uh.TestDataFolder = "testdata"
	}()

	fmt.Println(uh.GetTestDataFolder())

	// Output: something-else

}
