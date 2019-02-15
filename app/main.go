package main

import (
	"net/http"

	"github.com/levidurfee/ham/handlers"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/record-entry/", handlers.RecordEntryHandler)
	http.Handle("/", r)

	appengine.Main()
}
