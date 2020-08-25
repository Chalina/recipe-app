package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockController struct {
	MockHandleSearch func(w http.ResponseWriter, r *http.Request)
}

// Add method so that it implements controller interface
func (m mockController) HandleSearch(w http.ResponseWriter, r *http.Request) {
	m.MockHandleSearch(w, r)
}

func TestHandleSearch(t *testing.T) {
	controller := mockController{
		MockHandleSearch: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`test resp`))
			w.WriteHeader(http.StatusOK)
		},
	}
	server := CreateServer(controller)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/search", nil)

	server.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, `test resp`, recorder.Body.String())
}

func TestNotFound(t *testing.T) {
	controller := mockController{}

	server := CreateServer(controller)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/random", nil)

	server.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
