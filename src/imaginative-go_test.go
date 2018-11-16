package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
    "strings"
)

func TestDefaultHome(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := newRequestRecorder(req, "GET", "/", DefaultHome)
    if rr.Code != 200 {
		t.Error("Expected response code to be 200")
    }

    expected := `Hi, I’m Imaginative Go`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Error("Response body does not match")
    }
}

func TestSampleHelloWorld(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello-world", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := newRequestRecorder(req, "GET", "/hello-world", SampleHelloWorld)
    if rr.Code != 200 {
		t.Error("Expected response code to be 200")
    }

    expected := `hello, worlds`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Error("Response body does not match")
    }
}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}
