package unithelpers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

// ReadResponse converts the data in HTTP response to a string
func (h *TestHelper) ReadResponse(resp http.Response) string {
	h.Helper()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.Fatal("Cannot read from the HTTP response")
	}

	return string(data)
}

// CreateServer will initiate a test server and return the given response for the URL
func (h *TestHelper) CreateServer(path string, response string) (*httptest.Server, func()) {
	h.Helper()
	handler := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, path) {
			h.Fatalf("Specified URL path not available. Expected: %v, Got: %v", path, r.URL.Path)
		}
		w.Write([]byte(response))
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		defer ts.Close()
	}
}

// CreateServerWithHandler will initiate a test server and return the given response for the URL
func (h *TestHelper) CreateServerWithHandler(path string, response string, handler func(http.ResponseWriter, *http.Request)) (*httptest.Server, func()) {
	h.Helper()
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		defer ts.Close()
	}
}

// CreateServerWithError will create a test server and return the HTTP error code for the URL
func (h *TestHelper) CreateServerWithError(path string, response string, errorCode int) (*httptest.Server, func()) {
	h.Helper()
	handler := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, path) {
			h.Fatalf("Specified URL path not available. Expected: %v, Got: %v", path, r.URL.Path)
		}
		w.WriteHeader(errorCode)
		w.Write([]byte(response))
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		defer ts.Close()
	}
}
