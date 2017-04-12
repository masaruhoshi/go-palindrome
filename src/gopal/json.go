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

import (
	"log"
	"net/http"
	"encoding/json"
)

// Convenient struct used to marshal json messages
type errorMsg struct {
    code 	int  	`json:"code"`
    Message string	`json:"message"`
}

/*
Writes error message in JSON format to Response.

Message is assigned to erroMsg type along with error code and
marshalled into JSON instance.
*/
func JSONError(w http.ResponseWriter, message string, code int) {
	resp, err := json.Marshal(errorMsg{code, message})
	if err != nil {
		// If this is called, something really bad happened
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(resp)
}

/*
Writes the response in JSON format.

This method is written conveniently expecting a struct to be 
marshalled into a JSON object.
*/
func JSONResponse(w http.ResponseWriter, v interface{}, code int) {
	resp, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		// Again, if this happens, something really bad happened
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(resp)
}
