package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/levidurfee/ham/handlers"
	"google.golang.org/appengine"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)

	http.Handle("/", r)

	appengine.Main()
}
