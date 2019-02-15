package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/hamlog"
)

// HomeHandler handles loading the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	g := hamlog.BuildData(r)
	g.Template = "home.html"

	RenderTemplate(w, g)
}
