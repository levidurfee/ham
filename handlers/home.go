package handlers

import (
	"net/http"
	"text/template"
)

// HomeHandler loads the homepage HTML.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/base.html", "../templates/home.html"))
	tmpl.ExecuteTemplate(w, "base", nil)
}
