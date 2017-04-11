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

The Routing management is written based on HTTP Router package, a
lightweight and high performant http router. If you don't believe 
in the advantages of using this guy, check the performance benchmark
here:

	https://github.com/julienschmidt/go-http-routing-benchmark

I trust in the author's benchmark and also his oppinion about Pat
not being a very goo routing option.

This package is only used to keep routing attributions out of command
line application.
*/

package api

import (
	// Third party packages
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func NewRouter(routes Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		router.Handle( route.Method, route.Pattern, route.HandlerFunc )
	}

	return router
}

