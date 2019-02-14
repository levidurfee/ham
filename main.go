package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

// HAMPage is a page, but HAMPage sounds funnier
type HAMPage struct {
	Name     string
	Template string
	InNav    bool
}

var templateFolder = "templates"
var baseTemplate = "base.html"

func main() {
	home := HAMPage{
		Name:     "Home",
		Template: "home.html",
		InNav:    true,
	}

	indexHandler := buildHandler(home)

	http.HandleFunc("/", indexHandler)

	appengine.Main()
}

func buildHandler(page HAMPage) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles(templateFolder+"/"+page.Template, templateFolder+"/"+baseTemplate)
		if err != nil {
			panic("could not load template")
		}

		tmpl.ExecuteTemplate(w, "base", page)
	}

	return fn
}
