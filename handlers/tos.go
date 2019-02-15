package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/models"
)

// TOSHandler handles loading the login page
func TOSHandler(w http.ResponseWriter, r *http.Request) {
	p := models.NewPageData(r)
	p.Template = "tos.html"

	RenderTemplate(w, p)
}
