package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecipesByIngredient(t *testing.T) {
	repo, err := CreateNewRepository("sample_data.json")
	assert.Nil(t, err)

	tests := []struct {
		ingredientToSearch      string
		expectedNumberOfRecipes int
	}{
		{
			ingredientToSearch:      "1 tablespoon canola oil",
			expectedNumberOfRecipes: 12,
		},
		{
			ingredientToSearch:      "2 cloves garlic chopped",
			expectedNumberOfRecipes: 17,
		},
		{
			ingredientToSearch:      "chocolate",
			expectedNumberOfRecipes: 0,
		},
	}

	for _, tc := range tests {
		recipes, err := repo.GetRecipesByIngredient(tc.ingredientToSearch)
		assert.Nil(t, err)
		assert.Equal(t, tc.expectedNumberOfRecipes, len(recipes))
	}
}
