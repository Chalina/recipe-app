package recipes

type Recipe struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Method string `json:"method"`
}

type Ingredient struct {
	ID       int    `json:"id"`
	RecipeID int    `json:"recipeId"`
	Name     string `json:"name"`
	// Amount   int
	// Unit     unit
}
