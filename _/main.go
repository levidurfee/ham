package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/levidurfee/ham/tmpl"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	homeTmpl = tmpl.ParseTemplate("home.html")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/oauth2callback", oauthCallbackHandler)

	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pd := tmpl.PageData{
		LoggedIn: isLoggedIn(r),
		Data: struct {
			Msg        string
			IsLoggedIn bool
		}{
			Msg:        "Hi",
			IsLoggedIn: isLoggedIn(r),
		},
	}
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "Logged in: %v", pd.LoggedIn)
	homeTmpl.Execute(w, r, pd)
}
