package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type homeHandler struct{}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeWithIDRe = regexp.MustCompile(`/recipes/[0-9]+`)
)

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my home page"))
}

//Recipes Handler

type recipesHandler struct{}

func (h *recipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// GET	   /recipes
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		w.Write([]byte("GET /recipes reached"))
	// POST	   /recipes
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		w.Write([]byte("POST /recipes reached"))
	// GET	   /recipes/{id}
	case r.Method == http.MethodGet && RecipeWithIDRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		w.Write([]byte(fmt.Sprintf("GET	   /recipes/%v", id)))
	// PATCH   /recipes/{id}
	case r.Method == http.MethodPatch && RecipeWithIDRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		w.Write([]byte(fmt.Sprintf("PATCH   /recipes/%v", id)))
	// DELETE  /recipes/{id}
	case r.Method == http.MethodDelete && RecipeWithIDRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		w.Write([]byte(fmt.Sprintf("DELETE	   /recipes/%v", id)))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}
func main() {

	mux := http.NewServeMux()
	//Note : Handle does not support wildcard [PathValue]
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &recipesHandler{})
	mux.Handle("/recipes/", &recipesHandler{})
	http.ListenAndServe(":8080", mux)
}
