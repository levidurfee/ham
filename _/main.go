package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var (
	homeTmpl = ParseTemplate("home.html")
	tosTmpl  = ParseTemplate("tos.html")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/oauth2callback", oauthCallbackHandler)

	r.HandleFunc("/tos/", tosHandler)

	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pd := buildData(r)
	// ctx := appengine.NewContext(r)
	// log.Debugf(ctx, "Logged in: %v", pd.LoggedIn)
	homeTmpl.Execute(w, r, pd)
}

func tosHandler(w http.ResponseWriter, r *http.Request) {
	pd := buildData(r)
	tosTmpl.Execute(w, r, pd)
}

func buildData(r *http.Request) *PageData {
	return &PageData{
		LoggedIn: isLoggedIn(r),
		Data: struct {
			Msg        string
			IsLoggedIn bool
			Profile    *Profile
		}{
			Msg:        "Hi",
			IsLoggedIn: isLoggedIn(r),
			Profile:    profileFromSession(r),
		},
	}
}
