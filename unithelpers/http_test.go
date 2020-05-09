package unithelpers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uh "github.com/allentv/go-testhelpers/unithelpers"
)

func ExampleTestHelper_ReadResponse() {
	th := uh.New(&testing.T{})
	w := httptest.NewRecorder()
	w.Write([]byte("Something"))
	w.Flush()

	result := th.ReadResponse(w)
	fmt.Println(result)

	// Output: Something
}

func ExampleTestHelper_ReadHTTPResponse() {
	th := uh.New(&testing.T{})

	w := httptest.NewRecorder()
	w.Write([]byte("Something"))
	w.Flush()

	result := th.ReadHTTPResponse(w.Result())
	fmt.Println(result)

	// Output: Something
}

func ExampleTestHelper_CreateServer() {
	th := uh.New(&testing.T{})
	ts, closer := th.CreateServer("/something", "Hello something!")
	defer closer()

	resp, _ := ts.Client().Get(ts.URL + "/something")

	result := th.ReadHTTPResponse(resp)
	fmt.Println(result)

	// Output: Hello something!
}

func ExampleTestHelper_CreateServerWithHandler() {
	th := uh.New(&testing.T{})
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing"))
	}
	ts, closer := th.CreateServerWithHandler("/hello", handler)
	defer closer()

	resp, _ := ts.Client().Get(ts.URL + "/hello")

	result := th.ReadHTTPResponse(resp)
	fmt.Println(result)

	// Output: testing
}

func ExampleTestHelper_CreateServerWithError() {
	th := uh.New(&testing.T{})
	ts, closer := th.CreateServerWithError("/ping", "Cannot process request", http.StatusInternalServerError)
	defer closer()

	resp, _ := ts.Client().Get(ts.URL + "/ping")
	result := th.ReadHTTPResponse(resp)

	fmt.Println(result)
	fmt.Println(resp.StatusCode)

	// Output:
	// Cannot process request
	// 500
}
