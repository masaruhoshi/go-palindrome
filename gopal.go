package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const SERVICE_PORT string = "80"

type GoPal struct {
	Db *Dao
	Router *httprouter.Router
}

func New() *GoPal {
	settings := GetSettings()

	instance := new(GoPal)
	instance.Db = NewDao(settings)

	// Enforce index usage
	instance.Db.EnsureIndex()

	var routes = Routes{
		Route{
			"GET", "/palindrome", PalindromeListHandler(instance.Db),
		},
		Route{
			"POST", "/palindrome", PalindromeAddHandler(instance.Db),
		},
		Route{
			"GET", "/palindrome/:id", PalindromeGetHandler(instance.Db),
		},
		Route{
			"DELETE", "/palindrome/:id", PalindromeDeleteHandler(instance.Db),
		},
	}

	instance.Router = NewRouter(routes)

	return instance

}

// Run starts the server
func (g *GoPal) Run() error {
	defer g.Db.Close()

	log.Println("gopal is running under port "+SERVICE_PORT)
	return http.ListenAndServe(":"+SERVICE_PORT, g.Router)
}
