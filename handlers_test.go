package main

import (
	"testing"
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"bytes"

	"github.com/julienschmidt/httprouter"
)

func TestPalindromeListHandlerToReturnJsonListOfPalindromes(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("GET", "/palindrome", nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			listHandler := PalindromeListHandler(session)
			listHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		var palindromes []Palindrome
		decoder := json.NewDecoder(rr.Body)
		decoder.Decode(&palindromes)

		// Status should be OK
		Expect(t, rr.Code, http.StatusOK)
		// Entries created for test should be here
		Expect(t, len(palindromes), 10)
	})
}

func TestPalindromeAddHandlerToReturnCreated(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		var jsonStr = []byte(`{"phrase":"These are not the palindromes you're looking for"}`)

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("POST", "/palindrome", bytes.NewBuffer(jsonStr))
		if err != nil {
        	t.Fatal(err)
    	}
		r.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addHandler := PalindromeAddHandler(session)
			addHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusCreated)
	})
}

func TestPalindromeAddHandlerToReturnBadRequestOnInvalidRequest(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		var jsonStr = []byte(`not a valid json`)

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("POST", "/palindrome", bytes.NewBuffer(jsonStr))
		if err != nil {
        	t.Fatal(err)
    	}
		r.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addHandler := PalindromeAddHandler(session)
			addHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusBadRequest)
	})
}

func TestPalindromeAddHandlerToReturnBadRequestOnInvalidPhrase(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		var jsonStr = []byte(`{"phrase": ""}`)

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("POST", "/palindrome", bytes.NewBuffer(jsonStr))
		if err != nil {
        	t.Fatal(err)
    	}
		r.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addHandler := PalindromeAddHandler(session)
			addHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusBadRequest)
	})
}

func TestPalindromeAddHandlerToReturnBadRequestOnEmptyRequest(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		var jsonStr = []byte(``)

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("POST", "/palindrome", bytes.NewBuffer(jsonStr))
		if err != nil {
        	t.Fatal(err)
    	}
		r.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addHandler := PalindromeAddHandler(session)
			addHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusBadRequest)
	})
}

func TestPalindromeAddHandlerToReturnAlreadyReportedOnExistingPhrase(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		// It has to be an easier way to get the first/last element from a map
		var phrase string
		for _, v := range ht.Entries {
			phrase = fmt.Sprintf(`{"phrase": "%s"}`, v.Phrase)
			break;
		}

		var jsonStr = []byte(phrase)

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("POST", "/palindrome", bytes.NewBuffer(jsonStr))
		if err != nil {
        	t.Fatal(err)
    	}
		r.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addHandler := PalindromeAddHandler(session)
			addHandler(w, r, nil)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusAlreadyReported)
	})
}

func TestPalindromeGetHandlerToReturnValidObject(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		// Get random id from list of entries
		randomId := getRandomIdEntry(ht)

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: randomId,
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("GET", fmt.Sprintf("/palindrome/%s", randomId), nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeGetHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusOK)

	})	
}

func TestPalindromeGetHandlerToReturn404WithNonExistingId(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: "58ee7e93f1119f5c69292cb4",
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("GET", "/palindrome/58ee7e93f1119f5c69292cb4", nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeGetHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusNotFound)

	})	
}

func TestPalindromeGetHandlerToReturnPreconditionFailedWithInvalidId(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: "invalid id",
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("GET", "/palindrome/invalid", nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeGetHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusPreconditionFailed)

	})	
}

func TestPalindromeDeleteHandlerToReturnPreconditionFailedWithNonExistingId(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: "58ee7e93f1119f5c69292cb4",
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("DELETE", "/palindrome/58ee7e93f1119f5c69292cb4", nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeDeleteHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusNotFound)

	})	
}

func TestPalindromeDeleteHandlerToReturnPreconditionFailedWithInvalidId(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: "invalid id",
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("DELETE", "/palindrome/invalid", nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeDeleteHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusPreconditionFailed)

	})	
}

func TestPalindromeDeleteHandlerToReturnAccepted(t *testing.T) {
	ht := new(HandlerTest)
	ht.SetupTest( func() {
		session := ht.Session

		// Get random id from list of entries
		randomId := getRandomIdEntry(ht)

		params := httprouter.Params{
			httprouter.Param{
				Key: "id",
				Value: randomId,
			},
		}

		// Create a request to pass to handler. Parameters are not required
		r, err := http.NewRequest("DELETE", fmt.Sprintf("/palindrome/%s", randomId), nil)
		if err != nil {
        	t.Fatal(err)
    	}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getHandler := PalindromeDeleteHandler(session)
			getHandler(w, r, params)
		})

		// `handler` satisfies http.Handler, so ServeHTTP method is called directly 
		handler.ServeHTTP(rr, r)

		// Status should be OK
		Expect(t, rr.Code, http.StatusAccepted)

	})	
}

// It has to be an easier way to get the first/last element from a map
func getRandomIdEntry(ht *HandlerTest) string {
	// It has to be an easier way to get the first/last element from a map
	var randomId string
	for k, _ := range ht.Entries {
		randomId = k
		break;
	}

	return randomId
}