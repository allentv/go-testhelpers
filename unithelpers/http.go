package unithelpers

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (h *TestHelper) getResponse(body io.ReadCloser) string {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		h.fatal("Cannot read from the HTTP response")
	}

	return string(data)
}

// ReadResponse converts the data in HTTP test response to a string
func (h *TestHelper) ReadResponse(resp *httptest.ResponseRecorder) string {
	h.helper()
	return h.getResponse(resp.Result().Body)
}

// ReadHTTPResponse converts the data in HTTP response to a string
func (h *TestHelper) ReadHTTPResponse(resp *http.Response) string {
	h.helper()
	return h.getResponse(resp.Body)
}

// CreateServer will initiate a test server and return the given response for the URL
func (h *TestHelper) CreateServer(path string, response string) (*httptest.Server, func()) {
	h.helper()
	handler := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, path) {
			h.fatalf("Specified URL path not available. Expected: %v, Got: %v", path, r.URL.Path)
		}
		w.Header().Set("Content-Type", "utf-8")
		w.Write([]byte(response))
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		ts.Close()
	}
}

// CreateServerWithHandler will initiate a test server and return the given response for the URL
func (h *TestHelper) CreateServerWithHandler(path string, handler func(http.ResponseWriter, *http.Request)) (*httptest.Server, func()) {
	h.helper()
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		ts.Close()
	}
}

// CreateServerWithError will create a test server and return the HTTP error code for the URL
func (h *TestHelper) CreateServerWithError(path string, response string, errorCode int) (*httptest.Server, func()) {
	h.helper()
	handler := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, path) {
			h.fatalf("Specified URL path not available. Expected: %v, Got: %v", path, r.URL.Path)
		}
		w.WriteHeader(errorCode)
		w.Write([]byte(response))
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts, func() {
		ts.Close()
	}
}
