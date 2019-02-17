package handlers

import (
	"html/template"
	"net/http"

	"github.com/levidurfee/ham/app"
)

func renderTemplate(w http.ResponseWriter, d app.App) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
