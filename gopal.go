package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const SERVICE_PORT string = "8080"

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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	log.Println("gopal is running under port "+SERVICE_PORT)
	return http.ListenAndServe(":"+SERVICE_PORT, c.Handler(g.Router))
}
