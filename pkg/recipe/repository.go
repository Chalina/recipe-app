package recipe

type Repository struct{}

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

var collection map[string]Recipe = map[string]Recipe{
	"chocolate": Recipe{Name: "choc cake"},
	"butter":    Recipe{Name: "cookies"},
}

func (r Repository) GetRecipesByIngredient(ingredient string) ([]Recipe, error) {
	allRecipes := []Recipe{}
	recipe := collection[ingredient]
	allRecipes = append(allRecipes, recipe)
	return allRecipes, nil
}
