package handlers

import (
	"html/template"
	"net/http"

	"github.com/levidurfee/ham/page"
)

func renderTemplate(w http.ResponseWriter, d page.TemplateData) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
