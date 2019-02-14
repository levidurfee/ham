package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

type HAMPage struct {
	Name string
}

var folder = "templates"
var base = "base.html"

func main() {
	home := HAMPage{
		Name: "Home",
	}

	indexHandler := buildHandler(home)

	http.HandleFunc("/", indexHandler)

	appengine.Main()
}

func buildHandler(page HAMPage) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("templates/home.html", folder+"/"+base)
		if err != nil {
			panic("could not load template")
		}

		tmpl.ExecuteTemplate(w, "base", page)
	}

	return fn
}
