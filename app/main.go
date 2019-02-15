package main

import (
	"context"
	"net/http"

	"github.com/levidurfee/ham/handlers"
	"github.com/levidurfee/ham/repos"

	"github.com/gorilla/mux"
	"github.com/levidurfee/ham/hamlog"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/record-entry/", handlers.RecordEntryHandler)
	http.Handle("/", r)

	appengine.Main()
}

func userHasEntries(ctx context.Context, uid string) bool {
	var e hamlog.Entry
	key := datastore.NewKey(ctx, repos.QSO, uid, 0, nil)
	if err := datastore.Get(ctx, key, &e); err != nil {
		return false
	}

	return true
}
