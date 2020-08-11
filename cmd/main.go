package main

import (
	"log"
	"net/http"
	"recipe-app/pkg/api"
)

func main() {
	server := api.CreateServer()

	port := ":8080"
	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(port, server)
}
