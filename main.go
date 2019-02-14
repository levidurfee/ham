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

type HAMSite struct {
	User  string
	Pages []HAMPage
}

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

var ham = HAMSite{
	User:  "Levi",
	Pages: pages,
}

func main() {

	for _, v := range ham.Pages {
		http.HandleFunc(v.Route, buildHandler(v, ham))
	}

	http.HandleFunc("/register/", registerHandler)

	appengine.Main()
}

func buildHandler(page HAMPage, site HAMSite) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles(templateFolder+"/"+page.Template, templateFolder+"/"+baseTemplate)
		if err != nil {
			panic("could not load template")
		}

		tmpl.ExecuteTemplate(w, "base", site)
	}

	return fn
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom"))
}
