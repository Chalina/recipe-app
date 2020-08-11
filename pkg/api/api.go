package api

import (
	"log"
	"net/http"
)

func CreateServer() *http.ServeMux {
	r := http.ServeMux{}
	r.HandleFunc("/", handleRoutes)
	return &r
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET" && r.URL.Path == "/recipes":
		handleGetRecipes(w)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleGetRecipes(w http.ResponseWriter) {
	log.Print("in get recipes")
	w.Write([]byte(`{"hello": "json"}`))
}
