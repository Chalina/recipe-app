package main

import (
	"log"
	"net/http"
	"recipe-app/pkg/api"
	"recipe-app/pkg/recipe"
)

func main() {
	dataFilePath := "pkg/recipe/sample_data.json"
	recipeRepo, err := recipe.CreateNewRepository(dataFilePath)
	if err != nil {
		log.Fatalf("Error creating repo : %v", err)
	}
	controller := recipe.Controller{
		GetRecipesByIngredient: recipeRepo.GetRecipesByIngredient,
	}
	server := api.CreateServer(controller)

	port := ":8080"
	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(port, server)
}
