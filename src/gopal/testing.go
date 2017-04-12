package main

import (
	"fmt"
	"reflect"
	"testing"
	"encoding/json"
	"bytes"

	"gopkg.in/mgo.v2/bson"
)

/* Test Helper */
func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected ||%#v|| (type %v) - Got ||%#v|| (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func ExpectNil(t *testing.T, a interface{}) {
	if a != nil {
		t.Errorf("Expected ||nil|| - Got ||%#v|| (type %v)", a, reflect.TypeOf(a))
	}
}

func ExpectNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Errorf("Expected ||not nil|| - Got ||nil|| (type %v)", reflect.TypeOf(a))
	}
}

type HandlerTest struct {
	Session *Dao
	Entries map[string]Palindrome
}

func (h *HandlerTest) SetupTest(f func()) {
	settings := GetSettings()
	settings.DbName = "test"

	dao := NewDao(settings)
    dao.EnsureIndex()

    defer dao.Close()
	c := dao.Database().C("palindromes")

	h.Entries = make(map[string]Palindrome)
	h.Session = dao

	var palindrome Palindrome
	for i := 0; i < 10; i++ {
		palindrome = Palindrome{
			ID: bson.NewObjectId(),
			Phrase: fmt.Sprintf("test phrase %d", i),
		}
		c.Insert(&palindrome)
		h.Entries[palindrome.ID.Hex()] = palindrome
	}

	f()

	h.tearDown()
}

func (h *HandlerTest) tearDown() {
	c := h.Session.Database().C("palindromes")
	c.RemoveAll(bson.M{})
}

func (h *HandlerTest) AppendEntry(b *bytes.Buffer) {
	var palindrome Palindrome
	decoder := json.NewDecoder(b)
	decoder.Decode(&palindrome)

	h.Entries[palindrome.ID.Hex()] = palindrome
}
