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

Application settings should be defined here. 
 - Database host and credentials
 - Application routes
 - Application Port
*/

package main

import(
	"log"
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2"
	"pkg/api"
)

func main() {
	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err) // Mongodb not responding is a good reason to panic
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    // Enforce index usage
    api.EnsureIndex(session)

	var routes = api.Routes{
		api.Route{
			"GET", "/palindrome", api.PalindromeListHandler(session),
		},
		api.Route{
			"POST", "/palindrome", api.PalindromeAddHandler(session),
		},
		api.Route{
			"GET", "/palindrome/:id", api.PalindromeGetHandler(session),
		},
		api.Route{
			"DELETE", "/palindrome/:id", api.PalindromeDeleteHandler(session),
		},
	}

	router := api.NewRouter(routes)

    log.Fatal(http.ListenAndServe(":8080", router))
}