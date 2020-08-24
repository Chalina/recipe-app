package recipe

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	GetRecipesByIngredient func(ingredient string) ([]Recipe, error)
}

func (c Controller) HandleSearch(w http.ResponseWriter, r *http.Request) {
	type query struct {
		Ingredients []string
	}
	q := query{}

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
	recipes, err := c.GetRecipesByIngredient(q.Ingredients[0])
	if err != nil {
		log.Print("error getting ingredients")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// TODO: check for empty recipes
	type response struct {
		Recipes []Recipe
	}
	resp, _ := json.Marshal(response{
		Recipes: recipes,
	})
	w.Write([]byte(resp))
}
