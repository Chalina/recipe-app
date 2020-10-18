package recipe

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleSearch(t *testing.T) {
	recorder := httptest.NewRecorder()

	bodyString := `{"ingredients": ["chocolate"]}`
	reqBody := strings.NewReader(bodyString)

	req := httptest.NewRequest("POST", "/search", reqBody)

	controller := Controller{
		GetRecipesByIngredient: func(ingredient string) ([]Recipe, error) {
			return []Recipe{}, nil
		},
	}
	handler := http.HandlerFunc(controller.HandleSearch)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "{\"Recipes\":[]}", recorder.Body.String())
}

func TestHandleSearchEmptyBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/search", nil)

	controller := Controller{
		GetRecipesByIngredient: func(ingredient string) ([]Recipe, error) {
			return []Recipe{}, nil
		},
	}
	handler := http.HandlerFunc(controller.HandleSearch)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandleSearchInvalidBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	bodyString := `{"poodles": ["chocolate"]}`
	reqBody := strings.NewReader(bodyString)

	req := httptest.NewRequest("POST", "/search", reqBody)

	controller := Controller{
		GetRecipesByIngredient: func(ingredient string) ([]Recipe, error) {
			return []Recipe{}, nil
		},
	}
	handler := http.HandlerFunc(controller.HandleSearch)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
