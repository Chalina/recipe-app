package recipe

import (
	"encoding/json"
	"io/ioutil"
)

type Repository struct {
	collection map[string][]Recipe
}

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"title"`
	Ingredients []string `json:"ingredients"`
}

var collection map[string]Recipe = map[string]Recipe{
	"chocolate": Recipe{Name: "choc cake"},
	"butter":    Recipe{Name: "cookies"},
}

func CreateNewRepository() (Repository, error) {
	recipeMap, err := parseDataFile("pkg/recipe/sample_data.json")
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		collection: recipeMap,
	}, nil
}

func parseDataFile(filePath string) (map[string]Recipe, error) {
	recipeMap := map[string]Recipe{}
	recipes := []Recipe{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return recipeMap, err
	}

	// unmarshall into array of recipes
	if err := json.Unmarshal(file, &recipes); err != nil {
		return recipeMap, err
	}

	// index by ingredient
	for _, recipe := range recipes {
		// loop through ingredient list
		for _, ingredient := range recipe.Ingredients {
			recipeMap[ingredient] = recipe
		}
	}

	return recipeMap, nil
}

func (r Repository) GetRecipesByIngredient(ingredient string) ([]Recipe, error) {
	allRecipes := []Recipe{}
	recipe := r.collection[ingredient]
	allRecipes = append(allRecipes, recipe)
	return allRecipes, nil
}

// TODO: pass path to file from main
// TODO: update map with list of recipes
// TODO: add tests
