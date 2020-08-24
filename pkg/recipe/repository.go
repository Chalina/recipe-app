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

func (r Repository) GetRecipeByIngredient(ingredient string) (Recipe, error) {
	recipe := collection[ingredient]

	return recipe, nil
}
