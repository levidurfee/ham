package main

import (
	"context"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/levidurfee/ham/hamlog"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

// GOhamData is a struct for storing basic data
type GOhamData struct {
	LoggedIn   bool
	Login      string
	Logout     string
	Template   string
	User       *user.User
	HasEntries bool
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
	g.User = u
	if u == nil {
		g.LoggedIn = false
	}
	login, _ := user.LoginURL(ctx, "/")
	logout, _ := user.LogoutURL(ctx, "/")

	g.Login = login
	g.Logout = logout

	if g.LoggedIn {
		g.HasEntries = userHasEntries(ctx, u.ID)
	}

	return g
}

func userHasEntries(ctx context.Context, uid string) bool {
	var e hamlog.Entry
	key := datastore.NewKey(ctx, "Entry", uid, 0, nil)
	if err := datastore.Get(ctx, key, &e); err != nil {
		return false
	}

	return true
}

func renderTemplate(w http.ResponseWriter, d GOhamData) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
