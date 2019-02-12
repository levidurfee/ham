package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	appengine.Main()
}
