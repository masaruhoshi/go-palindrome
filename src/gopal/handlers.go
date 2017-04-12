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

package main

import(
	"encoding/json"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func PalindromeListHandler(dao *Dao) func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		instance := dao.GetInstance()
        defer instance.Close()
		c := instance.Database().C("palindromes")

		var palindromes []Palindrome
		err := c.Find(bson.M{}).All(&palindromes)
		if err != nil {
			JSONError(w, "Database error", http.StatusInternalServerError)
			log.Println("[palindromes] List fail: ", err)
			return
		}

		JSONResponse(w, palindromes, http.StatusOK)
	}
}

func PalindromeAddHandler(dao *Dao) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var palindrome Palindrome
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&palindrome)
		if err != nil {
			JSONError(w, "Invalid request", http.StatusBadRequest)
			log.Println("[palindromes] Invalid request: ", err)
			return
		}

		instance := dao.GetInstance()
        defer instance.Close()
		c := instance.Database().C("palindromes")

		err = palindrome.Validate()
		if err != nil {
			JSONError(w, "Invalid palindrome", http.StatusBadRequest)
			log.Println("[palindromes] Validation: ", err)
			return
		}

		err = c.Insert(palindrome)
		if err != nil {
			if mgo.IsDup(err) {
				JSONError(w, "Palindrome already exists", http.StatusAlreadyReported)
				log.Println("[palindromes] Duplicate: ", err)
				return
			}

			JSONError(w, "Database error", http.StatusInternalServerError)
			log.Println("[palindromes] Failed insert: ", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func PalindromeGetHandler(dao *Dao) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        id := p.ByName("id")
        if !bson.IsObjectIdHex(id) {
        	JSONError(w, "Invalid id", http.StatusPreconditionFailed)
        	log.Println("[palindrome] Invalid id: ", id)
        	return
        }

		instance := dao.GetInstance()
        defer instance.Close()
		c := instance.Database().C("palindromes")

		var palindrome Palindrome
		err := c.FindId(bson.ObjectIdHex(id)).One(&palindrome)
		if err != nil {
			switch err {
			default:
				JSONError(w, "Database error", http.StatusInternalServerError)
				log.Println("[palindrome] Failed get: ", err)
				return
			case mgo.ErrNotFound:
				JSONError(w, "Palindrome not found", http.StatusNotFound)
				log.Println("[palindrome] Not found: ", err)
				return
			}
		}

		JSONResponse(w, palindrome, http.StatusOK)
	}
}

func PalindromeDeleteHandler(dao *Dao) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		instance := dao.GetInstance()
		defer instance.Close()
		c := instance.Database().C("palindromes")

        id := p.ByName("id")
        if !bson.IsObjectIdHex(id) {
        	JSONError(w, "Invalid id", http.StatusPreconditionFailed)
        	log.Println("[palindrome] Invalid id: ", id)
        	return
        }

		err := c.RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			switch err {
			default:
				JSONError(w, "Database error", http.StatusInternalServerError)
				log.Println("[palindrome] Failed delete: ", err)
				return
			case mgo.ErrNotFound:
				JSONError(w, "Palindrome not found", http.StatusNotFound)
				log.Println("[palindrome] Not found: ", err)
				return
			}
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

