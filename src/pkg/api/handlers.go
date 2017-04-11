/*
Copyright 2017 Masaru Hoshi.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.

You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import(
	"encoding/json"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	// application modules
	"pkg/models"
	"pkg/utils"
)

func EnsureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("gopal").C("palindromes")

	index := mgo.Index{
		Key:        []string{"phrase"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func PalindromeListHandler(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := s.Copy()
        defer session.Close()
		c := session.DB("gopal").C("palindromes")

		var palindromes []models.Palindrome
		err := c.Find(bson.M{}).All(&palindromes)
		if err != nil {
			utils.JSONError(w, "Database error", http.StatusInternalServerError)
			log.Println("[palindromes] List fail: ", err)
			return
		}

		utils.JSONResponse(w, palindromes, http.StatusOK)
	}
}

func PalindromeAddHandler(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var palindrome models.Palindrome
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&palindrome)
		if err != nil {
			utils.JSONError(w, "Invalid request", http.StatusBadRequest)
			log.Println("[palindromes] Invalid request: ", err)
			return
		}

		session := s.Copy()
        defer session.Close()
		c := session.DB("gopal").C("palindromes")

		err = palindrome.Validate()
		if err != nil {
			utils.JSONError(w, "Invalid palindrome", http.StatusBadRequest)
			log.Println("[palindromes] Validation: ", err)
			return
		}

		err = c.Insert(palindrome)
		if err != nil {
			if mgo.IsDup(err) {
				utils.JSONError(w, "Palindrome already exists", http.StatusBadRequest)
				log.Println("[palindromes] Duplicate: ", err)
				return
			}

			utils.JSONError(w, "Database error", http.StatusInternalServerError)
			log.Println("[palindromes] Failed insert: ", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func PalindromeGetHandler(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        id := p.ByName("id")
        if !bson.IsObjectIdHex(id) {
        	utils.JSONError(w, "Invalid id", http.StatusNotFound)
        	log.Println("[palindrome] Invalid id: ", id)
        	return
        }

		session := s.Copy()
        defer session.Close()
		c := session.DB("gopal").C("palindromes")

		var palindrome models.Palindrome
		err := c.FindId(bson.ObjectIdHex(id)).One(&palindrome)
		if err != nil {
			switch err {
			default:
				utils.JSONError(w, "Database error", http.StatusInternalServerError)
				log.Println("[palindrome] Failed get: ", err)
				return
			case mgo.ErrNotFound:
				utils.JSONError(w, "Palindrome not found", http.StatusNotFound)
				log.Println("[palindrome] Not found: ", err)
				return
			}
		}

		utils.JSONResponse(w, palindrome, http.StatusOK)
	}
}

func PalindromeDeleteHandler(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session := s.Copy()
		defer session.Close()
		c := session.DB("gopal").C("palindromes")

        id := p.ByName("id")
        if !bson.IsObjectIdHex(id) {
        	utils.JSONError(w, "Invalid id", http.StatusNotFound)
        	log.Println("[palindrome] Invalid id: ", id)
        	return
        }

		err := c.RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			switch err {
			default:
				utils.JSONError(w, "Database error", http.StatusInternalServerError)
				log.Println("[palindrome] Failed delete: ", err)
				return
			case mgo.ErrNotFound:
				utils.JSONError(w, "Palindrome not found", http.StatusNotFound)
				log.Println("[palindrome] Not found: ", err)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

