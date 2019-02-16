package handlers

import (
	"net/http"
)

// HomeHandler handles loading the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var template = "home.html"

	RenderTemplate(w, p)
}
