package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/hamlog"
	"github.com/levidurfee/ham/id"
	"github.com/levidurfee/ham/repos"
	"google.golang.org/appengine"
)

// RecordEntryHandler records a new log entry
func RecordEntryHandler(w http.ResponseWriter, r *http.Request) {
	g := hamlog.BuildData(r)
	g.Template = "record-entry.html"

	switch r.Method {
	case http.MethodGet:

		if g.LoggedIn == false {
			g.Template = "please-login.html"
		}

		RenderTemplate(w, g)
	case http.MethodPost:
		// TODO create CSRF token and check it

		ctx := appengine.NewContext(r)

		hle := &hamlog.Entry{
			RequestID: id.GetID(ctx),
			UserID:    g.User.ID,
			CallSign:  r.PostFormValue("callsign"),
		}

		repos.StoreEntry(ctx, hle)

		http.Redirect(w, r, "/record-entry/", 302)
	}
}
