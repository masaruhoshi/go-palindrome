package utils

import (
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func TestJSONErrorToReturnJSONMessage(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JSONError(w, "test message", http.StatusInternalServerError)
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	h.ServeHTTP(res, req)

	Expect(t, res.Code, http.StatusInternalServerError)
	Expect(t, res.Header().Get("Content-Type"), "application/json; charset=utf-8")
	Expect(t, res.Body.String(), `{"message":"test message"}`)
}

func TestJSONResponseToReturnJSONMessage(t *testing.T) {
	type testData struct {
		Name	string 	`json:"name"`
		Age		int 	`json:"age"`
	}

	data := testData{Name: "Gandalf the Grey", Age: 2019}
	expected, _ := json.MarshalIndent(data, "", "  ")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, data, http.StatusOK)
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	h.ServeHTTP(res, req)

	Expect(t, res.Code, http.StatusOK)
	Expect(t, res.Header().Get("Content-Type"), "application/json; charset=utf-8")
	Expect(t, res.Body.String(), string(expected))
}
