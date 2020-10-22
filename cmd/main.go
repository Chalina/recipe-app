package main

import (
	"log"
	"net/http"
	"recipe-app/pkg/api"
	"recipe-app/pkg/recipe"
)

func main() {
	repo, err := recipe.NewMongoClient()
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}

	err = repo.AddNumbers()
	log.Printf("err: %v", err)

	controller := recipe.Controller{
		GetRecipesByIngredient: repo.GetRecipesByIngredient,
	}
	server := api.CreateServer(controller)

	port := ":8080"
	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(port, server)
}
