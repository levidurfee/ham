package main

import (
	"net/http"

	"github.com/levidurfee/ham/handlers"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"

	firebase "firebase.google.com/go"
)

var (
	firebaseConfig = &firebase.Config{
		DatabaseURL:   "https://hamradio-96e65.firebaseio.com",
		ProjectID:     "hamradio",
		StorageBucket: "hamradio.appspot.com",
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/login/", handlers.LoginHandler)
	r.HandleFunc("/tos/", handlers.TOSHandler)
	r.HandleFunc("/record-entry/", handlers.RecordEntryHandler)
	http.Handle("/", r)

	appengine.Main()
}
