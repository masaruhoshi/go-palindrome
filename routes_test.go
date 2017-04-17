package main

import (
	"testing"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func TestNewRouterToReturnRouterObject(t *testing.T) {
	settings := GetSettings()
	dao := NewDao(settings)

	routed := false
	testHandler := func(dao *Dao) func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		routed = true
		return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}
	}

	var routes = Routes{
		Route{
			"GET", "/testrequest", testHandler(dao),
		},
	}

	router := NewRouter(routes)
	// Mock response writter
	rr := new(mockResponseWriter)
	// Create a request to pass to handler.
	r, _ := http.NewRequest("GET", "/testrequest", nil)

	// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
	router.ServeHTTP(rr, r)

	Expect(t, routed, true)
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}
