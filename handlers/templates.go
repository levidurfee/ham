package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/levidurfee/ham/hamlog"
)

// RenderTemplate will write the template data to the browser
func RenderTemplate(w http.ResponseWriter, d hamlog.GOhamData) {
	w.Header().Set("Ham-Request-ID", strconv.FormatInt(d.RequestID, 10))
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
