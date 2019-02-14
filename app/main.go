package main

import (
	"context"
	"log"
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
	r.HandleFunc("/record-entry/", recordEntryHandler)
	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	g := buildData(r)
	g.Template = "home.html"

	renderTemplate(w, g)
}

func recordEntryHandler(w http.ResponseWriter, r *http.Request) {
	g := buildData(r)

	switch r.Method {
	case http.MethodGet:
		g.Template = "record-entry.html"

		if g.LoggedIn == false {
			g.Template = "please-login.html"
		}

		renderTemplate(w, g)
	case http.MethodPost:
		// Save data
	}
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
		e := &hamlog.Entry{
			UserID: u.ID,
		}
		storeEntry(ctx, e)
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

func storeEntry(ctx context.Context, entry *hamlog.Entry) {
	key := datastore.NewIncompleteKey(ctx, "Entry", nil)
	_, err := datastore.Put(ctx, key, entry)
	if err != nil {
		log.Println(err)
	}
}

func renderTemplate(w http.ResponseWriter, d GOhamData) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
