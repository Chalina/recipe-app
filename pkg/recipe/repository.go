package recipe

import (
	"encoding/json"
	"io/ioutil"
)

type Repository struct {
	collection map[string][]Recipe
}

var collection map[string]Recipe = map[string]Recipe{
	"chocolate": Recipe{Name: "choc cake"},
	"butter":    Recipe{Name: "cookies"},
}

// CreateNewRepository returns a new Repository containing recipe data from the sample file
func CreateNewRepository(path string) (Repository, error) {
	// recipeMap, err := parseDataFile("pkg/recipe/sample_data.json")
	recipeMap, err := parseDataFile(path)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		collection: recipeMap,
	}, nil
}

func parseDataFile(filePath string) (map[string][]Recipe, error) {
	recipeMap := map[string][]Recipe{}
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
			// if ingredient has been added before, append to list
			if _, ok := recipeMap[ingredient]; ok {
				recipeMap[ingredient] = append(recipeMap[ingredient], recipe)
			} else {
				// otherwise add ingredient as new key in map and push recipe
				recipeMap[ingredient] = append(recipeMap[ingredient], recipe)
			}
		}
	}

	return recipeMap, nil
}

// GetRecipesByIngredient returns all recipes containing that ingredient
func (r Repository) GetRecipesByIngredient(ingredient string) ([]Recipe, error) {
	return r.collection[ingredient], nil
}
