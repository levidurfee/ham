package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/page"
)

// HomeHandler handles loading the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	td := page.TemplateData{
		Title:    "Home",
		Template: "home.html",
	}

	renderTemplate(w, td)
}
