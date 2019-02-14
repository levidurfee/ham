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
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	var g GOhamData
	if u == nil {
		g.LoggedIn = false
	}
	login, _ := user.LoginURL(ctx, "/login")
	logout, _ := user.LogoutURL(ctx, "/")

	g.Login = login
	g.Logout = logout

	renderTemplate(w, g)
}

func renderTemplate(w http.ResponseWriter, d GOhamData) {
	tmpl := template.Must(template.ParseFiles("../templates/base.html", "../templates/home.html"))
	tmpl.ExecuteTemplate(w, "base", d)
}
