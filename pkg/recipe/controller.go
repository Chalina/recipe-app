package recipe

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct{
	GetRecipeByIngredient func(ingredient string) (Recipe, error)
}

type Query struct {
	Ingredients []string
}

func (c Controller) HandleSearch(w http.ResponseWriter, r *http.Request) {
	q := Query{}

	// r.Body is a reader
	// can use ioutil with helper functions to access the data, or json pkg which has a decoder
	// we can choose to read all the data at once, or bit by bit
	// in this case, the data is just a json obj as opposed to a big file so it's safe to read all at once
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("error reading request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		log.Print("invalid json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("req body: %v", q)
	recipes, err := c.GetRecipeByIngredient(q.Ingredients[0])
	if err != nil {
		log.Print("error getting ingredients")
		w.WriteHeader(http.StatusInternalServerError)
	}
	// TODO: check for empty recipes

	// recipes := []recipe.Recipe{
	// 	{Name: "Chocolate cake"},
	// }

	recipeResp, _ := json.Marshal(recipes)
	w.Write([]byte(recipeResp))
}
