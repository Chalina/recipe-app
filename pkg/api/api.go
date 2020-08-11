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
	case r.Method == "POST" && r.URL.Path == "/recipes":
		handlePostRecipes(w)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func handlePostRecipes(w http.ResponseWriter) {
	log.Print("in post recipes")
	w.Write([]byte(`{"hello": "json"}`))
}
