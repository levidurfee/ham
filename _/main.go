package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	appengine.Main()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}
