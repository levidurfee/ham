package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/base.html", "../templates/home.html"))
	tmpl.ExecuteTemplate(w, "base", nil)
}
