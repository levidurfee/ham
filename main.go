package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

// HAMPage is a page, but HAMPage sounds funnier
type HAMPage struct {
	Name     string
	Route    string
	Template string
}

// This is where the HTML templates live
var templateFolder = "templates"

// This is the base template that others extend
var baseTemplate = "base.html"

// pages is all the static pages on the site, mapped to their routes
var pages = []HAMPage{
	HAMPage{
		Route:    "/",
		Name:     "Home",
		Template: "home.html",
	},
	HAMPage{
		Route:    "/user",
		Name:     "User",
		Template: "home.html",
	},
}

func main() {

	for _, v := range pages {
		http.HandleFunc(v.Route, buildHandler(v))
	}

	http.HandleFunc("/register/", registerHandler)

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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom"))
}
