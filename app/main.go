package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// GOhamData is a struct for storing basic data
type GOhamData struct {
	LoggedIn bool
	Login    string
	Logout   string
	Template string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	g := buildData(r)
	g.Template = "home.html"

	renderTemplate(w, g)
}

func buildData(r *http.Request) GOhamData {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	var g GOhamData
	g.LoggedIn = true
	if u == nil {
		g.LoggedIn = false
	}
	login, _ := user.LoginURL(ctx, "/")
	logout, _ := user.LogoutURL(ctx, "/")

	g.Login = login
	g.Logout = logout

	return g
}

func renderTemplate(w http.ResponseWriter, d GOhamData) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
