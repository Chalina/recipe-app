package main

import (
	"log"
	"net/http"
	"recipe-app/pkg/api"
	"recipe-app/pkg/recipe"
)

func main() {
	recipeRepo := recipe.Repository{}
	controller := recipe.Controller{
		GetRecipeByIngredient: recipeRepo.GetRecipeByIngredient,
	}
	server := api.CreateServer(controller)

	port := ":8080"
	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(port, server)
}
