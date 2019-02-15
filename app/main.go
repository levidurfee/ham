package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/levidurfee/ham/hamlog"
	"github.com/levidurfee/ham/id"
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
	RequestID  int64
}

// QSO Entity type
var QSO = "QSOEntry"

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
	g.Template = "record-entry.html"

	switch r.Method {
	case http.MethodGet:

		if g.LoggedIn == false {
			g.Template = "please-login.html"
		}

		renderTemplate(w, g)
	case http.MethodPost:
		// TODO create CSRF token and check it

		ctx := appengine.NewContext(r)

		hle := &hamlog.Entry{
			RequestID: id.GetID(ctx),
			UserID:    g.User.ID,
			CallSign:  r.PostFormValue("callsign"),
		}

		storeEntry(ctx, hle)

		http.Redirect(w, r, "/record-entry/", 302)
	}
}

func buildData(r *http.Request) GOhamData {
	ctx := appengine.NewContext(r)
	ctx = id.CtxWithID(ctx)
	u := user.Current(ctx)
	var g GOhamData
	g.LoggedIn = true
	g.User = u
	g.RequestID = id.GetID(ctx)
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

	id.PrintID(ctx)

	return g
}

func userHasEntries(ctx context.Context, uid string) bool {
	var e hamlog.Entry
	key := datastore.NewKey(ctx, QSO, uid, 0, nil)
	if err := datastore.Get(ctx, key, &e); err != nil {
		return false
	}

	return true
}

// Maybe have this build a hamlog entry from a request
func storeEntry(ctx context.Context, entry *hamlog.Entry) {
	key := datastore.NewIncompleteKey(ctx, QSO, nil)
	r, err := datastore.Put(ctx, key, entry)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}

func renderTemplate(w http.ResponseWriter, d GOhamData) {
	w.Header().Set("Ham-Request-ID", strconv.FormatInt(d.RequestID, 10))
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/"+d.Template))
	tmpl.ExecuteTemplate(w, "base", d)
}
