package api

import (
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestPalindromeListHandlerToReturnJSONListOfPalindromes(t *testing.T) {
	PalindromeListHandler(s *mgo.Session)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, "test message", http.StatusInternalServerError)
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	h.ServeHTTP(res, req)	
}