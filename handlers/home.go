package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/models"
)

// HomeHandler handles loading the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := models.NewPageData(r)
	p.Template = "home.html"

	RenderTemplate(w, p)
}
