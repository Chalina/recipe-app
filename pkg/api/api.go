package api

import (
	"net/http"
)

type Controller interface {
	HandleSearch(w http.ResponseWriter, r *http.Request)
}

func CreateServer(c Controller) *http.ServeMux {
	r := http.ServeMux{}
	r.HandleFunc("/", handleRoutes(c))
	return &r
}

func handleRoutes(c Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST" && r.URL.Path == "/search":
			c.HandleSearch(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
